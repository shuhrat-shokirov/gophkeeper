package user

import (
	"context"

	"go.uber.org/fx"

	"gophkeeper/pkg/db"
)

var Module = fx.Provide(New)

type Repo interface {
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id int) (*User, error)

	CreateUser(ctx context.Context, user *User) (int, error)
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
