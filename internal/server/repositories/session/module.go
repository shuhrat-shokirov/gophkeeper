package session

import (
	"context"

	"go.uber.org/fx"

	"gophkeeper/pkg/db"
)

var Module = fx.Provide(New)

type Repo interface {
	Create(ctx context.Context, session *Session) error
	Get(ctx context.Context, refreshToken string) (*Session, error)
}

type Params struct {
	fx.In

	DBConn db.Conn
}

type repo struct {
	dbConn db.Conn
}

func New(p Params) Repo {
	return &repo{
		dbConn: p.DBConn,
	}
}
