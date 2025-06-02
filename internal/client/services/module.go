package services

import (
	"go.uber.org/fx"

	"gophkeeper/internal/client/services/auth"
)

var Module = fx.Options(
	auth.Module,
)
