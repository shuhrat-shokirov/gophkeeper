package gateways

import (
	"go.uber.org/fx"

	"gophkeeper/internal/server/gateways/emailtotp"
)

var Module = fx.Options(
	emailtotp.Module,
)
