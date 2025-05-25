package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"gophkeeper/internal/server/exceptions"
	"gophkeeper/internal/server/gateways/emailtotp"
	"gophkeeper/internal/server/repositories/user"
	"gophkeeper/pkg/utils"
)

func (s *service) Registration(ctx context.Context, request Registration) (string, error) {
	getUser, err := s.userRepo.GetUserByEmail(ctx, request.Email)
	if err != nil {
		if !errors.Is(err, exceptions.ErrNotFound) {
			return "", fmt.Errorf("error getting user: %w", err)
		}
	}

	if getUser != nil {
		return "", exceptions.ErrAlreadyExists
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}

	userID, err := s.userRepo.CreateUser(ctx, &user.User{
		Email:    request.Email,
		Password: string(password),
	})
	if err != nil {
		return "", fmt.Errorf("error creating user: %w", err)
	}

	otp, err := utils.GenerateOTP()
	if err != nil {
		return "", fmt.Errorf("error generating OTP: %w", err)
	}

	message, err := utils.GenerateOTPMessage(otp)
	if err != nil {
		return "", fmt.Errorf("error generating OTP message: %w", err)
	}

	err = s.emailTotpGateway.SendEmail(ctx, &emailtotp.Request{
		To:      request.Email,
		Subject: "Your One-Time Password",
		Body:    message,
	})
	if err != nil {
		return "", fmt.Errorf("error sending OTP email: %w", err)
	}

	uuid := utils.GenerateShortUUID()
	cacheKey := otpKeyPrefix + uuid

	if err := s.cache.Save(ctx, cacheKey, otpVerification{
		UserID: userID,
		Otp:    otp,
	}, otpExpiration); err != nil {
		return "", fmt.Errorf("error setting OTP in cache: %w", err)
	}

	return uuid, nil
}

type otpVerification struct {
	UserID int    `json:"user_id"`
	Otp    string `json:"otp"`
}

const (
	otpExpiration = 5 * time.Minute // OTP expiration time
	otpKeyPrefix  = "otp:"          // Prefix for OTP keys in cache
)
