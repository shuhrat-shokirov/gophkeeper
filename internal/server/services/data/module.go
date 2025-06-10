package data

import (
	"context"

	"go.uber.org/fx"

	"gophkeeper/internal/server/repositories/binary"
	"gophkeeper/internal/server/repositories/card"
	"gophkeeper/internal/server/repositories/logins"
	"gophkeeper/internal/server/repositories/texts"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In

	LoginsRepo logins.Repo
	TextRepo   texts.Repo
	CardRepo   card.Repo
	BinaryRepo binary.Repo
}

type Service interface {
	SaveLogin(ctx context.Context, data *LoginData) error
	SaveText(ctx context.Context, data *TextData) error
	SaveCard(ctx context.Context, data *CardData) error
	SaveBinary(ctx context.Context, data *BinaryData) error

	GetLoginList(ctx context.Context, userID, limit, offset int64) ([]LoginListItem, error)
	GetLoginByID(ctx context.Context, id int64) (*LoginInfo, error)
}

type service struct {
	loginsRepo logins.Repo
	textRepo   texts.Repo
	cardRepo   card.Repo
	binaryRepo binary.Repo
}

func New(p Params) Service {
	return &service{
		loginsRepo: p.LoginsRepo,
		textRepo:   p.TextRepo,
		cardRepo:   p.CardRepo,
		binaryRepo: p.BinaryRepo,
	}
}
