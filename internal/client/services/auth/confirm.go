package auth

import (
	"context"
	"fmt"
	"os"

	"gophkeeper/internal/client/exceptions"
)

func (s *service) ConfirmOTP(ctx context.Context, code string) error {
	otpId, ok := s.cache.Get(otpCodeKey)
	if !ok {
		return exceptions.ErrOtpExpired
	}

	id, _ := otpId.(string)

	token, err := s.serverGateway.ConfirmOtp(ctx, id, code)
	if err != nil {
		return fmt.Errorf("failed to confirm otp: %w", err)
	}

	s.accessToken = token.AccessToken
	s.refreshToken = token.RefreshToken
	return nil
}

func writeEnvFile(accessToken, refreshToken string) error {
	file, err := os.Create(".env")
	if err != nil {
		return fmt.Errorf("failed to create .env file: %w", err)
	}
	defer func() {
		_ = file.Close()
	}()

	_, err = fmt.Fprintf(file, "GOPH_KEEPER_ACCESS_TOKEN=%s\n", accessToken)
	if err != nil {
		return fmt.Errorf("failed to write access token: %w", err)
	}

	_, err = fmt.Fprintf(file, "GOPH_KEEPER_REFRESH_TOKEN=%s\n", refreshToken)
	if err != nil {
		return fmt.Errorf("failed to write refresh token: %w", err)
	}

	return nil
}
