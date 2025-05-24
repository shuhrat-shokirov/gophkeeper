package main

import (
	"go.uber.org/fx"

	"gophkeeper/internal/server/gateways"
	"gophkeeper/internal/server/grpc"
	"gophkeeper/internal/server/grpc/handlers"
	"gophkeeper/internal/server/repositories"
	"gophkeeper/internal/server/services"
	"gophkeeper/pkg/cache"
	"gophkeeper/pkg/config"
	"gophkeeper/pkg/db"
	"gophkeeper/pkg/jwt"
	"gophkeeper/pkg/logger"
	"gophkeeper/pkg/migration"
)

func main() {
	fx.New(
		grpc.Module,
		handlers.Module,

		repositories.Module,
		services.Module,
		gateways.Module,

		config.Module,
		logger.Module,
		migration.Module,
		db.Module,
		cache.Module,
		jwt.Module,
	).Run()
}
