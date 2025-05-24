package auth

import (
	"context"
	"fmt"
	"time"

	"gophkeeper/internal/server/exceptions"
)

func (s *service) RefreshToken(ctx context.Context, token string) (string, error) {
	session, err := s.sessionRepo.Get(ctx, token)
	if err != nil {
		return "", fmt.Errorf("refresh token: %w", err)
	}

	if session.ExpiredAt.Before(time.Now()) {
		return "", fmt.Errorf("refresh token expired, %w", exceptions.ErrRefreshTokenExpired)
	}

	accessToken, err := s.jwt.GenerateOnlyAccessToken(session.UserID)
	if err != nil {
		return "", fmt.Errorf("generate token pair: %w", err)
	}

	return accessToken, nil
}
