package handlers

import (
	"go.uber.org/fx"

	"go-template/internal/http/handlers/health"
)

var Module = fx.Options(
	health.Module,
)
