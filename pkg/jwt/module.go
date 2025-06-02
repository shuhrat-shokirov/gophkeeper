package jwt

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"go.uber.org/fx"

	"github.com/golang-jwt/jwt/v5"

	"gophkeeper/pkg/config"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In

	Config config.Config
}

type JWT interface {
	GenerateTokenPair(userID int) (*TokenPair, error)
	GenerateOnlyAccessToken(userID int) (string, error)
}

type jwtI struct {
	privateKey []byte
}

func New(p Params) (JWT, error) {
	decodeString, err := base64.StdEncoding.DecodeString(p.Config.GetString("jwt.private_key"))
	if err != nil {
		return nil, fmt.Errorf("failed to decode private key: %w", err)
	}

	return &jwtI{
		privateKey: decodeString,
	}, nil
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

// GenerateTokenPair generates a pair of tokens: access token and refresh token.
func (j *jwtI) GenerateTokenPair(userID int) (*TokenPair, error) {
	accessTokenString, err := j.GenerateOnlyAccessToken(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	// Refresh token (secure random string)
	refreshToken, err := generateSecureToken(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return &TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshToken,
	}, nil
}

// GenerateOnlyAccessToken generates only the access token without a refresh token.
func (j *jwtI) GenerateOnlyAccessToken(userID int) (string, error) {

	key, err := jwt.ParseRSAPrivateKeyFromPEM(j.privateKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["user_id"] = userID                       // User ID
	claims["exp"] = now.Add(30 * time.Minute).Unix() // Expiration time (30 minutes)
	claims["iat"] = now.Unix()                       // Issued at time
	claims["nbf"] = now.Unix()                       // Not before time
	claims["iss"] = "gophkeeper"                     // Issuer

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}

// Helper: generates a cryptographically secure random string
func generateSecureToken(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

var ErrTokenExpired = jwt.ErrTokenExpired

func Parse(token string, publicKey []byte) (jwt.Claims, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, fmt.Errorf("validate: parse key: %w", err)
	}

	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (any, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("validate: invalid")
	}

	return claims, nil
}
