package server

import (
	"context"
	"fmt"

	"gophkeeper/internal/client/exceptions"
	pb "gophkeeper/proto"
)

type Token struct {
	AccessToken  string
	RefreshToken string
}

func (g *gateway) ConfirmOtp(ctx context.Context, otpId, otpCode string) (*Token, error) {
	resp, err := g.client.ConfirmOTP(ctx, &pb.ConfirmOTPRequest{
		OtpId: otpId,
		Code:  otpCode,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to confirm OTP: %w", err)
	}

	if resp.GetStatus() == pb.ConfirmOTPStatus_INVALID_CODE {
		return nil, exceptions.ErrOtpInvalid
	}

	if resp.GetStatus() == pb.ConfirmOTPStatus_CODE_EXPIRED {
		return nil, exceptions.ErrOtpExpired
	}

	if resp.GetStatus() == pb.ConfirmOTPStatus_USER_NOT_FOUND_CODE {
		return nil, exceptions.ErrUserNotFound
	}

	if resp.GetStatus() != pb.ConfirmOTPStatus_CONFIRM_SUCCESS {
		return nil, fmt.Errorf("unexpected response status: %s", resp.GetMessage())
	}

	return &Token{
		AccessToken:  resp.GetToken(),
		RefreshToken: resp.GetRefreshToken(),
	}, nil
}
