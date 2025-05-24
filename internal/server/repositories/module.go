package repositories

import (
	"go.uber.org/fx"

	"gophkeeper/internal/server/repositories/session"
	"gophkeeper/internal/server/repositories/user"
)

var Module = fx.Options(
	user.Module,
	session.Module,
)
