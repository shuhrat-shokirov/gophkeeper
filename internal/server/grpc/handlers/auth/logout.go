package auth

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	pb "gophkeeper/proto"
)

func (h *handler) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	err := h.authService.Logout(ctx, request.GetToken())
	if err != nil {
		return &pb.LogoutResponse{
			Status:  pb.LogoutStatus_LOGOUT_ERROR.Enum(),
			Message: proto.String(fmt.Sprintf("failed to logout: %v", err)),
		}, fmt.Errorf("failed to logout: %w", err)
	}

	return &pb.LogoutResponse{
		Status:  pb.LogoutStatus_LOGOUT_SUCCESS.Enum(),
		Message: proto.String("Successfully logged out"),
	}, nil
}
