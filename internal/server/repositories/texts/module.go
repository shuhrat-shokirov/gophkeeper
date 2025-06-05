package texts

import (
	"context"

	"go.uber.org/fx"

	"gophkeeper/pkg/db"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In

	Conn db.Conn
}

type Repo interface {
	Save(ctx context.Context, data *TextData) (int64, error)
}

type repo struct {
	conn db.Conn
}

func New(p Params) Repo {
	return &repo{
		conn: p.Conn,
	}
}
