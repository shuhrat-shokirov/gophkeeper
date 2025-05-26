package auth

import (
	"context"
	"errors"
	"fmt"

	"gophkeeper/internal/server/exceptions"
	pb "gophkeeper/proto"
)

func (h *handler) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	otpId, err := h.authService.Login(ctx, request.GetEmail(), request.GetPassword())
	if err != nil {
		if errors.Is(err, exceptions.ErrNotFound) {
			return &pb.LoginResponse{
				Status:  pb.LoginStatus_INVALID_CREDENTIALS,
				Message: "Invalid email or password",
			}, nil
		}
		return &pb.LoginResponse{
			Status:  pb.LoginStatus_LOGIN_ERROR,
			Message: "Login error: " + err.Error(),
		}, fmt.Errorf("Login error for user %s: %w", request.GetEmail(), err)
	}

	return &pb.LoginResponse{
		Status: pb.LoginStatus_LOGIN_SUCCESS,
		OtpId:  otpId,
	}, nil
}
