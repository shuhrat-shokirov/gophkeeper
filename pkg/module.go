package pkg

import (
	"go.uber.org/fx"

	"go-template/pkg/config"
	"go-template/pkg/logger"
	"go-template/pkg/reply"
	"go-template/pkg/translation"
)

var Module = fx.Options(
	config.Module,
	logger.Module,
	reply.Module,
	translation.Module,
)
