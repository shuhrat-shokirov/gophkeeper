package services

import (
	"go.uber.org/fx"

	"gophkeeper/internal/server/services/auth"
	"gophkeeper/internal/server/services/data"
)

var Module = fx.Options(
	auth.Module,
	data.Module,
)
