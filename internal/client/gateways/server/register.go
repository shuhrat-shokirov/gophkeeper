package server

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	"gophkeeper/internal/client/errorx"
	pb "gophkeeper/proto"
)

func (g *gateway) Register(ctx context.Context, email, password string) (string, error) {
	resp, err := g.authServiceClient.Register(ctx, &pb.RegisterRequest{
		Email:    proto.String(email),
		Password: proto.String(password),
	})
	if err != nil {
		return "", fmt.Errorf("failed to register: %w", err)
	}

	if resp.GetStatus() == pb.RegisterStatus_USER_ALREADY_EXISTS {
		return "", errorx.ErrUserAlreadyExists
	}

	if resp.GetStatus() != pb.RegisterStatus_OTP_SENT {
		return "", fmt.Errorf("unexpected response status: %s", resp.GetMessage())
	}

	return resp.GetOtpId(), nil
}
