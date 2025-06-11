package logger

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/fx/fxtest"

	"gophkeeper/pkg/config"
)

func TestNew(t *testing.T) {
	t.Run("should return a new logger", func(t *testing.T) {
		newConfig, err := config.NewConfig()
		require.Nil(t, err)
		require.NotNil(t, newConfig)

		logger, err := New(Params{
			Config:    newConfig,
			Lifecycle: fxtest.NewLifecycle(t),
		})
		require.Nil(t, err)
		require.NotNil(t, logger)

		logger.Info("Test logger info message")
		logger.Debug("Test logger debug message")
		logger.Warn("Test logger warn message")
		logger.Error("Test logger error message")
	})
}

func Test_getLevel(t *testing.T) {
	t.Run("should return default level", func(t *testing.T) {
		newConfig, err := config.NewConfig()
		require.Nil(t, err)
		require.NotNil(t, newConfig)

		level := getLevel(newConfig)
		require.Equal(t, "debug", level.String())
	})

	t.Run("debug level", func(t *testing.T) {
		err := os.Setenv("LOGGER_LEVEL", "debug")
		require.Nil(t, err)

		newConfig, err := config.NewConfig()
		require.Nil(t, err)
		require.NotNil(t, newConfig)

		level := getLevel(newConfig)
		require.Equal(t, "debug", level.String())
	})

	t.Run("warn level", func(t *testing.T) {
		err := os.Setenv("LOGGER_LEVEL", "warning")
		require.Nil(t, err)

		newConfig, err := config.NewConfig()
		require.Nil(t, err)
		require.NotNil(t, newConfig)

		level := getLevel(newConfig)
		require.Equal(t, "warn", level.String())
	})

	t.Run("error level", func(t *testing.T) {
		err := os.Setenv("LOGGER_LEVEL", "error")
		require.Nil(t, err)

		newConfig, err := config.NewConfig()
		require.Nil(t, err)
		require.NotNil(t, newConfig)

		level := getLevel(newConfig)
		require.Equal(t, "error", level.String())
	})

	t.Run("info level", func(t *testing.T) {
		err := os.Setenv("LOGGER_LEVEL", "info")
		require.Nil(t, err)

		newConfig, err := config.NewConfig()
		require.Nil(t, err)
		require.NotNil(t, newConfig)

		level := getLevel(newConfig)
		require.Equal(t, "info", level.String())
	})
}
