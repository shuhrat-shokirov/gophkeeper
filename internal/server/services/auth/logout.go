package auth

import (
	"context"
	"fmt"

	"gophkeeper/internal/server/exceptions"
)

func (s *service) Logout(ctx context.Context, accessToken string) error {
	if accessToken == "" {
		return fmt.Errorf("empty access token, %w", exceptions.ErrTokenNotFound)
	}

	err := s.sessionRepo.Delete(ctx, accessToken)
	if err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}

	return nil
}
