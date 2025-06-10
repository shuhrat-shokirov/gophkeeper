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

	List(ctx context.Context, userID int64, pg Pagination) ([]List, error)
	GetByID(ctx context.Context, id int64) (*Info, error)
}

type repo struct {
	dbConn db.Conn
}

func New(p Params) Repo {
	return &repo{
		dbConn: p.DBConn,
	}
}
