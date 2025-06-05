package handlers

import (
	"go.uber.org/fx"

	"gophkeeper/internal/client/tui/handlers/render"
)

var Module = fx.Options(
	render.Module,
)
