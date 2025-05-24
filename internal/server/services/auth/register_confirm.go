package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"gophkeeper/internal/server/exceptions"
	"gophkeeper/internal/server/repositories/session"
)

func (s *service) ConfirmRegistration(ctx context.Context, id, code string) (*ConfirmResponse, error) {
	cacheKey := otpRegisterKeyPrefix + id
	value, err := s.cache.Find(ctx, cacheKey)
	if err != nil {
		if errors.Is(err, exceptions.ErrNotFound) {
			return nil, exceptions.ErrOTPExpired
		}

		return nil, fmt.Errorf("cache find error: %w", err)
	}

	var otp otpVerification

	err = json.Unmarshal(value, &otp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal otp verification error: %w", err)
	}

	if otp.Otp != code {
		return nil, exceptions.ErrInvalidOTP
	}

	_, err = s.userRepo.GetUserByID(ctx, otp.UserID)
	if err != nil {
		if errors.Is(err, exceptions.ErrNotFound) {
			return nil, exceptions.ErrNotFound
		}

		return nil, fmt.Errorf("error getting user by id: %w", err)
	}

	pair, err := s.jwt.GenerateTokenPair(otp.UserID)
	if err != nil {
		return nil, fmt.Errorf("error generating token pair: %w", err)
	}

	err = s.sessionRepo.Create(ctx, &session.Session{
		UserID:       otp.UserID,
		RefreshToken: pair.RefreshToken,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating session: %w", err)
	}

	_ = s.cache.Delete(ctx, cacheKey)

	return &ConfirmResponse{
		UserId: otp.UserID,
		Token:  pair.AccessToken,
	}, nil
}
