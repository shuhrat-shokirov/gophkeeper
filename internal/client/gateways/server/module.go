package server

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "gophkeeper/proto"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In
}

type Gateway interface {
	Register(ctx context.Context, email, password string) (string, error)
	ConfirmOtp(ctx context.Context, otpId, otpCode string) (*Token, error)
	RefreshToken(ctx context.Context, refreshToken string) (*Token, error)
}

type gateway struct {
	client pb.AuthServiceClient
}

func New(p Params) (Gateway, error) {
	url := os.Getenv("GOPH_KEEPER_GRPC_ADDRESS")
	if url == "" {
		url = "localhost:9091"
	}

	conn, err := grpc.NewClient(url,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	log.Printf("connecting to %s", url)

	return &gateway{
		client: pb.NewAuthServiceClient(conn),
	}, nil
}
