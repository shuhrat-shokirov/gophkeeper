package data

import (
	"context"

	"go.uber.org/fx"

	"gophkeeper/internal/client/gateways/server"
	"gophkeeper/internal/client/services/auth"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In

	ServerGateway server.Gateway
	AuthService   auth.Service
}

type Service interface {
	SaveLogin(ctx context.Context, data *LoginData) error
	SaveText(ctx context.Context, data *TextData) error
	SaveCard(ctx context.Context, data *CardData) error
	SaveFile(ctx context.Context, data *FileData) error

	GetLoginList(ctx context.Context, limit, offset int64) ([]ListItem, error)
	GetLoginByID(ctx context.Context, id int64) (*LoginInfo, error)

	GetTextList(ctx context.Context, limit, offset int64) ([]ListItem, error)
	GetTextByID(ctx context.Context, id int64) (*TextInfo, error)

	GetCardList(ctx context.Context, limit, offset int64) ([]ListItem, error)
	GetCardByID(ctx context.Context, id int64) (*CardInfo, error)

	GetBinaryList(ctx context.Context, limit, offset int64) ([]ListItem, error)
	GetBinaryByID(ctx context.Context, id int64) (*BinaryInfo, error)
}

type service struct {
	serverGateway server.Gateway
	authService   auth.Service
}

func New(p Params) Service {
	return &service{
		serverGateway: p.ServerGateway,
		authService:   p.AuthService,
	}
}
