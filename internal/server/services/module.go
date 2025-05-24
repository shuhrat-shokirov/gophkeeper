package services

import (
	"go.uber.org/fx"

	"gophkeeper/internal/server/services/auth"
)

var Module = fx.Options(
	auth.Module,
)
