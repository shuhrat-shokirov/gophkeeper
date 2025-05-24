package logger

import (
	"github.com/aliftechuz/pkg/logger"
	"go.uber.org/fx"

	"go-template/pkg/config"
)

var Module = fx.Provide(NewLogger)

type Params struct {
	fx.In

	Config config.Config
}

type Logger = logger.Logger

func NewLogger(p Params) Logger {
	return logger.New(p.Config.GetString("logger.level"))
}
