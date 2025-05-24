package jwt

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strconv"
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
	jwtSecret string
}

func New(p Params) JWT {
	return &jwtI{
		jwtSecret: p.Config.GetString("jwt.secret"),
	}
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
	// Access token (JWT)
	accessTokenClaims := jwt.MapClaims{
		"user_id": strconv.Itoa(userID),
		"exp":     time.Now().Add(30 * time.Minute).Unix(),
		"iat":     time.Now().Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString(j.jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign access token: %w", err)
	}

	return accessTokenString, nil
}

// Helper: generates a cryptographically secure random string
func generateSecureToken(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
