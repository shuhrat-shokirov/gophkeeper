package auth

import (
	"context"
	"errors"
	"fmt"
	"log"

	"gophkeeper/internal/server/errorx"
	"gophkeeper/internal/server/repositories/session"
)

func (s *service) ConfirmOTP(ctx context.Context, id, code string) (*ConfirmResponse, error) {
	cacheKey := otpKeyPrefix + id
	var otp otpVerification

	err := s.cache.Find(ctx, cacheKey, &otp)
	if err != nil {
		if errors.Is(err, errorx.ErrNotFound) {
			return nil, errorx.ErrOTPExpired
		}

		return nil, fmt.Errorf("cache find error: %w", err)
	}

	if otp.Otp != code {
		return nil, errorx.ErrInvalidOTP
	}

	_, err = s.userRepo.GetUserByID(ctx, otp.UserID)
	if err != nil {
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

	log.Printf("access token: %s, refresh token: %s", pair.AccessToken, pair.RefreshToken)

	_ = s.cache.Delete(ctx, cacheKey)

	return &ConfirmResponse{
		UserId:       otp.UserID,
		Token:        pair.AccessToken,
		RefreshToken: pair.RefreshToken,
	}, nil
}
