package migration

import (
	"errors"
	"fmt"

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

func New(p Params) error {
	m, err := migrate.New(p.Config.GetString("migration.file"), p.Config.GetString("migration.dsn"))
	if err != nil {
		p.Logger.Error("err from migration.New", zap.Error(err))
		return fmt.Errorf("err from migration.New: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		p.Logger.Error("err from up migration", zap.Error(err))
		return fmt.Errorf("err from up migration: %w", err)
	}

	p.Logger.Info("migration completed successfully")
	return nil
}
