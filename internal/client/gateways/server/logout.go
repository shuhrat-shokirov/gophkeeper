package server

import (
	"context"

	pb "gophkeeper/proto"
)

func (g *gateway) Logout(ctx context.Context, refreshToken string) {
	_, _ = g.client.Logout(ctx, &pb.LogoutRequest{
		Token: refreshToken,
	})
}
