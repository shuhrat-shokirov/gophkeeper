package auth

import (
	"context"
	"fmt"

	"gophkeeper/internal/client/exceptions"
	"gophkeeper/pkg/utils"
)

func (s *service) Login(ctx context.Context, email, password string) error {
	if !utils.ValidateEmail(email) {
		return exceptions.ErrEmailInvalidFormat
	}

	if len(password) < 6 {
		return exceptions.ErrPasswordTooShort
	}

	otpId, err := s.serverGateway.Login(ctx, email, password)
	if err != nil {
		return fmt.Errorf("login: %w", err)
	}

	s.cache.Set(otpCodeKey, otpId, cacheDuration)

	return nil
}
