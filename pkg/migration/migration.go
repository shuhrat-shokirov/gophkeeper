package migration

import (
	"errors"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"gophkeeper/pkg/config"
	"gophkeeper/pkg/logger"
)

var Module = fx.Options(
	fx.Invoke(
		New,
	),
)

type Params struct {
	fx.In
	Logger logger.Logger
	Config config.Config
}

func New(p Params) {
	m, err := migrate.New(p.Config.GetString("migration.file"), p.Config.GetString("migration.dsn"))
	if err != nil {
		p.Logger.Error("err from migration.New", zap.Error(err))
		os.Exit(1)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		p.Logger.Error("err from up migration", zap.Error(err))
		os.Exit(1)
	}

	p.Logger.Info("migration completed successfully")
}
