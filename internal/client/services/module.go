package services

import (
	"go.uber.org/fx"

	"gophkeeper/internal/client/services/auth"
	"gophkeeper/internal/client/services/data"
)

var Module = fx.Options(
	auth.Module,
	data.Module,
)
