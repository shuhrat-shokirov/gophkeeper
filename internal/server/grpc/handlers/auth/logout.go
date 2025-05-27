package auth

import (
	"context"
	"fmt"

	pb "gophkeeper/proto"
)

func (h *handler) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	err := h.authService.Logout(ctx, request.GetToken())
	if err != nil {
		h.logger.Error("failed to logout", "error", err)
		return &pb.LogoutResponse{
			Status:  pb.LogoutStatus_LOGOUT_ERROR,
			Message: fmt.Sprintf("failed to logout: %v", err),
		}, fmt.Errorf("failed to logout: %w", err)
	}

	return &pb.LogoutResponse{
		Message: "Successfully logged out",
	}, nil
}
