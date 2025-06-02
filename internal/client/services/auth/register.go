package auth

import (
	"context"
	"fmt"
	"time"

	"gophkeeper/internal/client/errorx"
	"gophkeeper/pkg/utils"
)

func (s *service) Register(ctx context.Context, email, password string) error {
	if !utils.ValidateEmail(email) {
		return errorx.ErrEmailInvalidFormat
	}

	if len(password) < 6 {
		return errorx.ErrPasswordTooShort
	}

	otpID, err := s.serverGateway.Register(ctx, email, password)
	if err != nil {
		return fmt.Errorf("register error: %w", err)
	}

	s.cache.Set(otpCodeKey, otpID, cacheDuration)

	return nil
}

const (
	cacheDuration = 5 * time.Minute
	otpCodeKey    = "otp_code"
)
