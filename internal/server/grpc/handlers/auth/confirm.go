package auth

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/protobuf/proto"

	"gophkeeper/internal/server/errorx"
	pb "gophkeeper/proto"
)

func (h *handler) ConfirmOTP(ctx context.Context,
	request *pb.ConfirmOTPRequest) (*pb.ConfirmOTPResponse, error) {
	response, err := h.authService.ConfirmOTP(ctx, request.GetOtpId(), request.GetCode())
	if err != nil {
		if errors.Is(err, errorx.ErrOTPExpired) {
			return &pb.ConfirmOTPResponse{
				Status:  pb.ConfirmOTPStatus_CODE_EXPIRED.Enum(),
				Message: proto.String("OTP code has expired"),
			}, nil
		}
		if errors.Is(err, errorx.ErrInvalidOTP) {
			return &pb.ConfirmOTPResponse{
				Status:  pb.ConfirmOTPStatus_INVALID_CODE.Enum(),
				Message: proto.String("Invalid OTP code"),
			}, nil
		}

		if errors.Is(err, errorx.ErrNotFound) {
			return &pb.ConfirmOTPResponse{
				Status:  pb.ConfirmOTPStatus_USER_NOT_FOUND_CODE.Enum(),
				Message: proto.String("User not found"),
			}, nil
		}

		return &pb.ConfirmOTPResponse{
			Status:  pb.ConfirmOTPStatus_CONFIRM_ERROR.Enum(),
			Message: proto.String("Confirmation error: " + err.Error()),
		}, fmt.Errorf("error confirming registration: %w", err)
	}

	return &pb.ConfirmOTPResponse{
		Status:       pb.ConfirmOTPStatus_CONFIRM_SUCCESS.Enum(),
		UserId:       proto.Int64(int64(response.UserId)),
		Token:        proto.String(response.Token),
		RefreshToken: proto.String(response.RefreshToken),
	}, nil
}
