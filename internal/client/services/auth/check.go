package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"gophkeeper/internal/client/errorx"
	"gophkeeper/pkg/jwt"
)

func (s *service) CheckAuth(ctx context.Context) error {
	if s.accessToken == "" || s.refreshToken == "" {
		return errorx.ErrTokenNotFound
	}

	token, err := jwt.Parse(s.accessToken, s.publicKey)
	if err != nil {
		if !errors.Is(err, jwt.ErrTokenExpired) {
			log.Printf("validate: parse token: %v", err)
			return fmt.Errorf("token validate error: %w", err)
		}
	} else {
		claimExpiration, err := token.GetExpirationTime()
		if err != nil {
			return fmt.Errorf("get expiration time error: %w", err)
		}

		if time.Now().Add(5 * time.Minute).Before(claimExpiration.Time) {
			return nil
		}
	}

	authToken, err := s.serverGateway.RefreshToken(ctx, s.refreshToken)
	if err != nil {
		return fmt.Errorf("failed to refresh token: %w", err)
	}

	s.accessToken = authToken.AccessToken
	return nil
}

func (s *service) GetUserID(_ context.Context) (int64, error) {
	if s.userID != 0 {
		return int64(s.userID), nil
	}

	parse, err := jwt.Parse(s.accessToken, s.publicKey)
	if err != nil {
		if !errors.Is(err, jwt.ErrTokenExpired) {
			log.Printf("parse token: %v", err)
			return 0, fmt.Errorf("parse token: %w", err)
		}
	}

	userID, err := jwt.GetUserIDFromClaims(parse)
	if err != nil {
		return 0, fmt.Errorf("get user id from claims: %w", err)
	}

	s.userID = userID
	return int64(userID), nil
}
