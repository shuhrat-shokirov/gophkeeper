package tui

import (
	"go.uber.org/fx"

	"gophkeeper/internal/client/tui/handlers"
	"gophkeeper/internal/client/tui/router"
)

var Module = fx.Options(
	router.Module,
	handlers.Module,
)
