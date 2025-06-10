package migration

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/fx/fxtest"

	"gophkeeper/pkg/config"
	"gophkeeper/pkg/logger"
)

func TestNew(t *testing.T) {
	t.Run("can't find file", func(t *testing.T) {
		err := os.Setenv("MIGRATION_FILE", "file://./testdata/migrations")
		require.NoError(t, err)

		newConfig, err := config.NewConfig()
		require.NoError(t, err)
		require.NotNil(t, newConfig)

		l, err := logger.New(logger.Params{
			Lifecycle: fxtest.NewLifecycle(t),
			Config:    newConfig,
		})
		require.NoError(t, err)
		require.NotNil(t, l)

		err = New(Params{
			Logger: l,
			Config: newConfig,
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "err from migration.New")
	})
	t.Run("success", func(t *testing.T) {
		err := os.Setenv("MIGRATION_FILE", "file://../../migrations/server")
		require.NoError(t, err)

		err = os.Setenv("MIGRATION_DSN", "postgres://postgres:postgres@localhost:5432/goph_keeper_db?sslmode=disable")
		require.NoError(t, err)

		newConfig, err := config.NewConfig()
		require.NoError(t, err)
		require.NotNil(t, newConfig)

		l, err := logger.New(logger.Params{
			Lifecycle: fxtest.NewLifecycle(t),
			Config:    newConfig,
		})
		require.NoError(t, err)
		require.NotNil(t, l)

		err = New(Params{
			Logger: l,
			Config: newConfig,
		})
		require.NoError(t, err)
	})
}
