package server

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	"gophkeeper/internal/client/errorx"
	pb "gophkeeper/proto"
)

func (g *gateway) RefreshToken(ctx context.Context, refreshToken string) (*Token, error) {
	resp, err := g.authServiceClient.RefreshToken(ctx, &pb.RefreshTokenRequest{
		RefreshToken: proto.String(refreshToken),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}

	if resp.GetStatus() == pb.RefreshTokenStatus_INVALID_REFRESH_TOKEN {
		return nil, errorx.ErrRefreshTokenInvalid
	}

	if resp.GetStatus() != pb.RefreshTokenStatus_REFRESH_SUCCESS {
		return nil, fmt.Errorf("failed to refresh token: %s", resp.GetMessage())
	}

	return &Token{
		AccessToken: resp.GetToken(),
	}, nil

}
