package auth

import (
	"context"
	"fmt"

	"gophkeeper/internal/client/errorx"
	"gophkeeper/pkg/utils"
)

func (s *service) Login(ctx context.Context, email, password string) error {
	if !utils.ValidateEmail(email) {
		return errorx.ErrEmailInvalidFormat
	}

	if len(password) < 6 {
		return errorx.ErrPasswordTooShort
	}

	otpId, err := s.serverGateway.Login(ctx, email, password)
	if err != nil {
		return fmt.Errorf("login: %w", err)
	}

	s.cache.Set(otpCodeKey, otpId, cacheDuration)

	return nil
}
