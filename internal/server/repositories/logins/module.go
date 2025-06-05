package logins

import (
	"context"

	"go.uber.org/fx"

	"gophkeeper/pkg/db"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In

	DBConn db.Conn
}

type Repo interface {
	Save(ctx context.Context, login *LoginData) (int, error)
}

type repo struct {
	dbConn db.Conn
}

func New(p Params) Repo {
	return &repo{
		dbConn: p.DBConn,
	}
}
