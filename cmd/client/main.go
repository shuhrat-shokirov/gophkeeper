package main

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"gophkeeper/internal/client/gateways"
	"gophkeeper/internal/client/services"
	"gophkeeper/internal/client/tui"
	"gophkeeper/pkg/memorycache"
)

func main() {
	fx.New(
		fx.WithLogger(func() fxevent.Logger {
			return fxevent.NopLogger
		}),

		memorycache.Module,

		gateways.Module,

		services.Module,

		tui.Module,
	).Run()
}
