package server

import (
	"context"

	"google.golang.org/protobuf/proto"

	pb "gophkeeper/proto"
)

func (g *gateway) Logout(ctx context.Context, refreshToken string) {
	_, _ = g.client.Logout(ctx, &pb.LogoutRequest{
		Token: proto.String(refreshToken),
	})
}
