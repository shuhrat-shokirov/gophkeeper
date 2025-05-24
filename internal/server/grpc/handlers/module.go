package handlers

import (
	"go.uber.org/fx"

	"gophkeeper/internal/server/grpc/handlers/auth"
)

var Module = fx.Options(
	auth.Module,
)
