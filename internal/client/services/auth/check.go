package auth

import (
	"context"
	"fmt"
	"log"
	"time"

	"gophkeeper/internal/client/exceptions"
	"gophkeeper/pkg/jwt"
)

func (s *service) CheckAuth(ctx context.Context) error {
	if s.accessToken == "" || s.refreshToken == "" {
		return exceptions.ErrTokenNotFound
	}

	token, err := jwt.Parse(s.accessToken, s.publicKey)
	if err != nil {
		log.Printf("validate: parse token: %v", err)
		return fmt.Errorf("token validate error: %w", err)
	}

	claimExpiration, err := token.GetExpirationTime()
	if err != nil {
		return fmt.Errorf("get expiration time error: %w", err)
	}

	if !claimExpiration.Before(time.Now()) {
		return nil
	}

	authToken, err := s.serverGateway.RefreshToken(ctx, s.refreshToken)
	if err != nil {
		return fmt.Errorf("failed to refresh token: %w", err)
	}

	s.accessToken = authToken.AccessToken
	return nil
}
