package config

import (
	"fmt"

	"github.com/aliftechuz/pkg/config"
	"go.uber.org/fx"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In
}

type Config = config.Config

func New(_ Params) (Config, error) {
	cf, err := config.NewConfig(
		config.SetConfigType(config.JSON),
		config.AddConfigPath("./configs"),
		config.SetConfigName("config"),
	)
	if err != nil {
		return nil, fmt.Errorf("can't init configs: %w", err)
	}

	return cf, nil
}
