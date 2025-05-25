package server

import (
	"context"
	"fmt"

	"gophkeeper/internal/client/exceptions"
	pb "gophkeeper/proto"
)

func (g *gateway) Register(ctx context.Context, email, password string) (string, error) {
	resp, err := g.client.Register(ctx, &pb.RegisterRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", fmt.Errorf("failed to register: %w", err)
	}

	if resp.GetStatus() == pb.RegisterStatus_USER_ALREADY_EXISTS {
		return "", exceptions.ErrUserAlreadyExists
	}

	if resp.GetStatus() != pb.RegisterStatus_OTP_SENT {
		return "", fmt.Errorf("unexpected response status: %s", resp.GetMessage())
	}

	return resp.GetOtpId(), nil
}
