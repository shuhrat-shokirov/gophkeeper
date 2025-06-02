package auth

import (
	"context"
	"fmt"
)

func (s *service) RefreshToken(ctx context.Context, token string) (string, error) {
	session, err := s.sessionRepo.Get(ctx, token)
	if err != nil {
		return "", fmt.Errorf("refresh token: %w", err)
	}

	accessToken, err := s.jwt.GenerateOnlyAccessToken(session.UserID)
	if err != nil {
		return "", fmt.Errorf("generate token pair: %w", err)
	}

	return accessToken, nil
}
