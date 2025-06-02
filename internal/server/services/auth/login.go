package auth

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"gophkeeper/internal/server/errorx"
)

func (s *service) Login(ctx context.Context, email, password string) (string, error) {
	userInfo, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", fmt.Errorf("failed to get user by email: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid credentials: %w", errorx.ErrInvalidCredentials)
	}

	otpId, err := s.senOtp(ctx, userInfo.ID, email)
	if err != nil {
		return "", fmt.Errorf("failed to send OTP: %w", err)
	}

	return otpId, nil
}
