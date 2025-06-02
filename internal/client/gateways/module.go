package gateways

import (
	"go.uber.org/fx"

	"gophkeeper/internal/client/gateways/server"
)

var Module = fx.Options(
	server.Module,
)
