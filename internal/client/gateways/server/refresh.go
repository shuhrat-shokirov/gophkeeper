package server

import (
	"context"
	"fmt"

	"gophkeeper/internal/client/exceptions"
	pb "gophkeeper/proto"
)

func (g *gateway) RefreshToken(ctx context.Context, refreshToken string) (*Token, error) {
	resp, err := g.client.RefreshToken(ctx, &pb.RefreshTokenRequest{
		RefreshToken: refreshToken,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}

	if resp.GetStatus() == pb.RefreshTokenStatus_EXPIRED_REFRESH_TOKEN {
		return nil, exceptions.ErrRefreshTokenExpired
	}

	if resp.GetStatus() == pb.RefreshTokenStatus_INVALID_REFRESH_TOKEN {
		return nil, exceptions.ErrRefreshTokenInvalid
	}

	if resp.GetStatus() != pb.RefreshTokenStatus_REFRESH_SUCCESS {
		return nil, fmt.Errorf("failed to refresh token: %s", resp.GetMessage())
	}

	return &Token{
		AccessToken: resp.GetToken(),
	}, nil

}
