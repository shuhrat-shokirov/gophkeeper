package auth

import (
	"context"
	"errors"
	"fmt"

	"gophkeeper/internal/server/exceptions"
	pb "gophkeeper/proto"
)

func (h *handler) RefreshToken(ctx context.Context, request *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	token, err := h.authService.RefreshToken(ctx, request.RefreshToken)
	if err != nil {
		if errors.Is(err, exceptions.ErrNotFound) {
			return &pb.RefreshTokenResponse{
				Status:  pb.RefreshTokenStatus_INVALID_REFRESH_TOKEN,
				Message: "Invalid refresh token",
			}, nil
		}
		if errors.Is(err, exceptions.ErrRefreshTokenExpired) {
			return &pb.RefreshTokenResponse{
				Status:  pb.RefreshTokenStatus_EXPIRED_REFRESH_TOKEN,
				Message: "Refresh token has expired",
			}, nil
		}

		return &pb.RefreshTokenResponse{
			Status:  pb.RefreshTokenStatus_REFRESH_ERROR,
			Message: "Error refreshing token: " + err.Error(),
		}, fmt.Errorf("error refreshing token: %w", err)
	}

	return &pb.RefreshTokenResponse{
		Status: pb.RefreshTokenStatus_REFRESH_SUCCESS,
		Token:  token,
	}, nil
}
