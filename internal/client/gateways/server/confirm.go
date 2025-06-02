package server

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	"gophkeeper/internal/client/errorx"
	pb "gophkeeper/proto"
)

type Token struct {
	AccessToken  string
	RefreshToken string
}

func (g *gateway) ConfirmOtp(ctx context.Context, otpId, otpCode string) (*Token, error) {
	resp, err := g.client.ConfirmOTP(ctx, &pb.ConfirmOTPRequest{
		OtpId: proto.String(otpId),
		Code:  proto.String(otpCode),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to confirm OTP: %w", err)
	}

	if resp.GetStatus() == pb.ConfirmOTPStatus_INVALID_CODE {
		return nil, errorx.ErrOtpInvalid
	}

	if resp.GetStatus() == pb.ConfirmOTPStatus_CODE_EXPIRED {
		return nil, errorx.ErrOtpExpired
	}

	if resp.GetStatus() == pb.ConfirmOTPStatus_USER_NOT_FOUND_CODE {
		return nil, errorx.ErrUserNotFound
	}

	if resp.GetStatus() != pb.ConfirmOTPStatus_CONFIRM_SUCCESS {
		return nil, fmt.Errorf("unexpected response status: %s", resp.GetMessage())
	}

	return &Token{
		AccessToken:  resp.GetToken(),
		RefreshToken: resp.GetRefreshToken(),
	}, nil
}
