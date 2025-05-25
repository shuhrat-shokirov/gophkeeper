package handlers

import (
	"go.uber.org/fx"

	"gophkeeper/internal/client/tui/handlers/auth"
)

var Module = fx.Options(
	auth.Module,
)
