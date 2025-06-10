package repositories

import (
	"go.uber.org/fx"

	"gophkeeper/internal/server/repositories/binary"
	"gophkeeper/internal/server/repositories/card"
	"gophkeeper/internal/server/repositories/logins"
	"gophkeeper/internal/server/repositories/session"
	"gophkeeper/internal/server/repositories/texts"
	"gophkeeper/internal/server/repositories/user"
)

var Module = fx.Options(
	user.Module,
	session.Module,
	logins.Module,
	texts.Module,
	card.Module,
	binary.Module,
)
