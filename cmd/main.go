package main

import (
	"go.uber.org/fx"

	"go-template/internal/gateways"
	"go-template/internal/http"
	"go-template/internal/repositories"
	"go-template/internal/services"
	"go-template/pkg"
)

func main() {
	fx.New(
		http.Module,

		gateways.Module,
		repositories.Module,
		services.Module,

		pkg.Module,
	).Run()
}
