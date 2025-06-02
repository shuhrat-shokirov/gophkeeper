package auth

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/protobuf/proto"

	"gophkeeper/internal/server/errorx"
	pb "gophkeeper/proto"
)

func (h *handler) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	otpId, err := h.authService.Login(ctx, request.GetEmail(), request.GetPassword())
	if err != nil {
		if errors.Is(err, errorx.ErrNotFound) {
			return &pb.LoginResponse{
				Status:  pb.LoginStatus_INVALID_CREDENTIALS.Enum(),
				Message: proto.String("Invalid email or password"),
			}, nil
		}
		return &pb.LoginResponse{
			Status:  pb.LoginStatus_LOGIN_ERROR.Enum(),
			Message: proto.String("Login error: " + err.Error()),
		}, fmt.Errorf("login error for user %s: %w", request.GetEmail(), err)
	}

	return &pb.LoginResponse{
		Status: pb.LoginStatus_LOGIN_SUCCESS.Enum(),
		OtpId:  proto.String(otpId),
	}, nil
}
