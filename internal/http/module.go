package http

import (
	"go.uber.org/fx"

	"go-template/internal/http/handlers"
	"go-template/internal/http/middleware"
	"go-template/internal/http/routes"
)

var Module = fx.Options(
	routes.Module,

	middleware.Module,

	handlers.Module,
)
