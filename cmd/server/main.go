package main

import (
	"go.uber.org/fx"

	"gophkeeper/internal/server/grpc"
	"gophkeeper/internal/server/grpc/handlers"
	"gophkeeper/internal/server/repositories"
	"gophkeeper/internal/server/services"
	"gophkeeper/pkg"
)

func main() {
	fx.New(
		grpc.Module,
		handlers.Module,

		repositories.Module,
		services.Module,

		pkg.Module,
	).Run()
}
