package auth

import (
	"context"
	"errors"
	"fmt"

	"gophkeeper/internal/server/exceptions"
	"gophkeeper/internal/server/services/auth"
	pb "gophkeeper/proto"
)

func (h *handler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	id, err := h.authService.Registration(ctx, auth.Registration{
		Email:    request.GetEmail(),
		Password: request.GetPassword(),
	})
	if err != nil {
		if errors.Is(err, exceptions.ErrAlreadyExists) {
			return &pb.RegisterResponse{
				Status:  pb.RegisterStatus_USER_ALREADY_EXISTS,
				Message: "User with this email already exists",
			}, nil
		}

		return &pb.RegisterResponse{
			Status:  pb.RegisterStatus_REGISTER_ERROR,
			Message: "Registration error: " + err.Error(),
		}, fmt.Errorf("error registering user: %w", err)
	}

	return &pb.RegisterResponse{
		Status: pb.RegisterStatus_OTP_SENT,
		OtpId:  id,
	}, nil
}

func (h *handler) ConfirmRegistration(ctx context.Context,
	request *pb.ConfirmRegistrationRequest) (*pb.ConfirmRegistrationResponse, error) {
	response, err := h.authService.ConfirmRegistration(ctx, request.GetOtpId(), request.GetCode())
	if err != nil {
		if errors.Is(err, exceptions.ErrOTPExpired) {
			return &pb.ConfirmRegistrationResponse{
				Status:  pb.ConfirmRegistrationStatus_CODE_EXPIRED,
				Message: "OTP code has expired",
			}, nil
		}
		if errors.Is(err, exceptions.ErrInvalidOTP) {
			return &pb.ConfirmRegistrationResponse{
				Status:  pb.ConfirmRegistrationStatus_INVALID_CODE,
				Message: "Invalid OTP code",
			}, nil
		}

		if errors.Is(err, exceptions.ErrNotFound) {
			return &pb.ConfirmRegistrationResponse{
				Status:  pb.ConfirmRegistrationStatus_USER_NOT_FOUND_CODE,
				Message: "User not found",
			}, nil
		}

		return &pb.ConfirmRegistrationResponse{
			Status:  pb.ConfirmRegistrationStatus_CONFIRM_ERROR,
			Message: "Confirmation error: " + err.Error(),
		}, fmt.Errorf("error confirming registration: %w", err)
	}

	return &pb.ConfirmRegistrationResponse{
		Status: pb.ConfirmRegistrationStatus_CONFIRM_SUCCESS,
		UserId: int64(response.UserId),
		Token:  response.Token,
	}, nil
}
