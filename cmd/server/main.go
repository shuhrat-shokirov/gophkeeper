package main

import (
	"go.uber.org/fx"

	"gophkeeper/internal/server/gateways"
	"gophkeeper/internal/server/grpc"
	"gophkeeper/internal/server/repositories"
	"gophkeeper/internal/server/services"
	"gophkeeper/pkg/config"
	"gophkeeper/pkg/db"
	"gophkeeper/pkg/jwt"
	"gophkeeper/pkg/logger"
	"gophkeeper/pkg/migration"
	"gophkeeper/pkg/redis"
)

func main() {
	fx.New(
		grpc.Module,

		repositories.Module,
		services.Module,
		gateways.Module,

		config.Module,
		logger.Module,
		migration.Module,
		db.Module,
		redis.Module,
		jwt.Module,
	).Run()
}
