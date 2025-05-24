package middleware

import (
	"go.uber.org/fx"

	"go-template/pkg/logger"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In

	Logger logger.Logger
}

type Middleware interface {
}

type middleware struct {
	logger logger.Logger
}

func New(p Params) Middleware {
	return &middleware{
		logger: p.Logger,
	}
}
