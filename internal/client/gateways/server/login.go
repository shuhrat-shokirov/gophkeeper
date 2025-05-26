package server

import (
	"context"
	"fmt"

	"gophkeeper/internal/client/exceptions"
	pb "gophkeeper/proto"
)

func (g *gateway) Login(ctx context.Context, email, password string) (string, error) {
	resp, err := g.client.Login(ctx, &pb.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", fmt.Errorf("login: %w", err)
	}

	if resp.GetStatus() == pb.LoginStatus_INVALID_CREDENTIALS {
		return "", exceptions.ErrInvalidCredentials
	}

	if resp.GetStatus() != pb.LoginStatus_LOGIN_SUCCESS {
		return "", fmt.Errorf("unexpected response status: %s", resp.GetMessage())
	}

	return resp.GetOtpId(), nil
}
