package server

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	"gophkeeper/internal/client/errorx"
	pb "gophkeeper/proto"
)

func (g *gateway) Login(ctx context.Context, email, password string) (string, error) {
	resp, err := g.client.Login(ctx, &pb.LoginRequest{
		Email:    proto.String(email),
		Password: proto.String(password),
	})
	if err != nil {
		return "", fmt.Errorf("login: %w", err)
	}

	if resp.GetStatus() == pb.LoginStatus_INVALID_CREDENTIALS {
		return "", errorx.ErrInvalidCredentials
	}

	if resp.GetStatus() != pb.LoginStatus_LOGIN_SUCCESS {
		return "", fmt.Errorf("unexpected response status: %s", resp.GetMessage())
	}

	return resp.GetOtpId(), nil
}
