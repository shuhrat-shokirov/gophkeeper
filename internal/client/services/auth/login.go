package auth

import (
	"context"

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

	return nil
}
