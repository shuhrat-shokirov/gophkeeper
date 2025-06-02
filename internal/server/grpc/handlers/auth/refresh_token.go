package auth

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/protobuf/proto"

	"gophkeeper/internal/server/errorx"
	pb "gophkeeper/proto"
)

func (h *handler) RefreshToken(ctx context.Context, request *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	token, err := h.authService.RefreshToken(ctx, request.GetRefreshToken())
	if err != nil {
		if errors.Is(err, errorx.ErrNotFound) {
			return &pb.RefreshTokenResponse{
				Status:  pb.RefreshTokenStatus_INVALID_REFRESH_TOKEN.Enum(),
				Message: proto.String("Invalid refresh token"),
			}, nil
		}

		return &pb.RefreshTokenResponse{
			Status:  pb.RefreshTokenStatus_REFRESH_ERROR.Enum(),
			Message: proto.String("Error refreshing token: " + err.Error()),
		}, fmt.Errorf("error refreshing token: %w", err)
	}

	return &pb.RefreshTokenResponse{
		Status: pb.RefreshTokenStatus_REFRESH_SUCCESS.Enum(),
		Token:  proto.String(token),
	}, nil
}
