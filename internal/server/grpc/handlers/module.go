package handlers

import (
	"go.uber.org/fx"

	"gophkeeper/internal/server/grpc/handlers/auth"
	"gophkeeper/internal/server/grpc/handlers/data"
)

var Module = fx.Options(
	auth.Module,
	data.Module,
)
