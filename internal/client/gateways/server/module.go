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
	Login(ctx context.Context, email, password string) (string, error)
	Logout(ctx context.Context, refreshToken string)

	SaveLoginAndPass(ctx context.Context, userID int64, pass *LoginAndPass) error
	SaveText(ctx context.Context, userID int64, data *Text) error
	SaveCard(ctx context.Context, userID int64, data *Card) error
	SaveBinary(ctx context.Context, userID int64, data *Binary) error

	GetLoginList(ctx context.Context, userID int64, limit, offset int64) ([]ListItem, error)
	GetLoginByID(ctx context.Context, userID, id int64) (*LoginInfo, error)

	GetCardList(ctx context.Context, userID int64, limit, offset int64) ([]ListItem, error)
	GetCardByID(ctx context.Context, userID, id int64) (*CardInfo, error)

	GetTextList(ctx context.Context, userID int64, limit, offset int64) ([]ListItem, error)
	GetTextByID(ctx context.Context, userID, id int64) (*TextInfo, error)

	GetBinaryList(ctx context.Context, userID int64, limit, offset int64) ([]ListItem, error)
	GetBinaryByID(ctx context.Context, userID, id int64) (*BinaryInfo, error)
}

type gateway struct {
	authServiceClient pb.AuthServiceClient
	dataServiceClient pb.DataServiceClient
}

func New(p Params) (Gateway, error) {
	url := os.Getenv("GOPH_KEEPER_GRPC_ADDRESS")
	if url == "" {
		url = "localhost:9091"
	}

	conn, err := grpc.NewClient(url,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc authServiceClient: %w", err)
	}

	log.Printf("connecting to %s", url)

	return &gateway{
		authServiceClient: pb.NewAuthServiceClient(conn),
		dataServiceClient: pb.NewDataServiceClient(conn),
	}, nil
}
