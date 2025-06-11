package auth

import (
	"context"
	"fmt"

	"gophkeeper/internal/server/errorx"
)

func (s *service) Logout(ctx context.Context, refreshToken string) error {
	if refreshToken == "" {
		return fmt.Errorf("empty access token, %w", errorx.ErrTokenNotFound)
	}

	err := s.sessionRepo.Delete(ctx, refreshToken)
	if err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}

	return nil
}
