package auth

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/protobuf/proto"

	"gophkeeper/internal/server/errorx"
	"gophkeeper/internal/server/services/auth"
	pb "gophkeeper/proto"
)

func (h *handler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	id, err := h.authService.Registration(ctx, auth.Registration{
		Email:    request.GetEmail(),
		Password: request.GetPassword(),
	})
	if err != nil {
		if errors.Is(err, errorx.ErrAlreadyExists) {
			return &pb.RegisterResponse{
				Status:  pb.RegisterStatus_USER_ALREADY_EXISTS.Enum(),
				Message: proto.String("User with this email already exists"),
			}, nil
		}

		return &pb.RegisterResponse{
			Status:  pb.RegisterStatus_REGISTER_ERROR.Enum(),
			Message: proto.String("Registration error: " + err.Error()),
		}, fmt.Errorf("error registering user: %w", err)
	}

	return &pb.RegisterResponse{
		Status: pb.RegisterStatus_OTP_SENT.Enum(),
		OtpId:  proto.String(id),
	}, nil
}
