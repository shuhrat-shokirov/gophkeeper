package auth

import (
	"context"
	"errors"
	"fmt"

	"gophkeeper/internal/server/exceptions"
	pb "gophkeeper/proto"
)

func (h *handler) ConfirmOTP(ctx context.Context,
	request *pb.ConfirmOTPRequest) (*pb.ConfirmOTPResponse, error) {
	response, err := h.authService.ConfirmOTP(ctx, request.GetOtpId(), request.GetCode())
	if err != nil {
		if errors.Is(err, exceptions.ErrOTPExpired) {
			return &pb.ConfirmOTPResponse{
				Status:  pb.ConfirmOTPStatus_CODE_EXPIRED,
				Message: "OTP code has expired",
			}, nil
		}
		if errors.Is(err, exceptions.ErrInvalidOTP) {
			return &pb.ConfirmOTPResponse{
				Status:  pb.ConfirmOTPStatus_INVALID_CODE,
				Message: "Invalid OTP code",
			}, nil
		}

		if errors.Is(err, exceptions.ErrNotFound) {
			return &pb.ConfirmOTPResponse{
				Status:  pb.ConfirmOTPStatus_USER_NOT_FOUND_CODE,
				Message: "User not found",
			}, nil
		}

		return &pb.ConfirmOTPResponse{
			Status:  pb.ConfirmOTPStatus_CONFIRM_ERROR,
			Message: "Confirmation error: " + err.Error(),
		}, fmt.Errorf("error confirming registration: %w", err)
	}

	return &pb.ConfirmOTPResponse{
		Status:       pb.ConfirmOTPStatus_CONFIRM_SUCCESS,
		UserId:       int64(response.UserId),
		Token:        response.Token,
		RefreshToken: response.RefreshToken,
	}, nil
}
