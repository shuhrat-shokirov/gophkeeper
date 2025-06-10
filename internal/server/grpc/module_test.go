package grpc

import (
	"crypto/rand"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx/fxtest"

	"gophkeeper/internal/server/grpc/handlers/auth"
	"gophkeeper/internal/server/grpc/handlers/data"
	"gophkeeper/pkg/config"
	"gophkeeper/pkg/logger"
)

func TestNew(t *testing.T) {
	t.Run("err on listen", func(t *testing.T) {
		err := os.Setenv("GRPC_ADDRESS", "invalid_address")
		require.NoError(t, err)

		newConfig, err := config.NewConfig()
		require.NoError(t, err)
		require.NotNil(t, newConfig)

		authHandler := auth.NewMockHandler()
		defer authHandler.AssertExpectations(t)

		authHandler.On("RegisterService", mock.Anything)

		dataHandler := data.NewMockHandler()
		defer dataHandler.AssertExpectations(t)

		dataHandler.On("RegisterService", mock.Anything)

		lifecycle := fxtest.NewLifecycle(t)

		l, err := logger.New(logger.Params{
			Lifecycle: lifecycle,
			Config:    newConfig,
		})
		require.NoError(t, err)
		require.NotNil(t, l)

		err = New(Params{
			Lifecycle:   lifecycle,
			Config:      newConfig,
			Logger:      l,
			AuthHandler: authHandler,
			DataHandler: dataHandler,
		})
		require.Error(t, err, "expected error on invalid address, got nil")
	})

	t.Run("success", func(t *testing.T) {
		port, err := getRandomPort()
		require.NoError(t, err)

		err = os.Setenv("GRPC_ADDRESS", "localhost:"+strconv.Itoa(port))
		require.NoError(t, err)

		newConfig, err := config.NewConfig()
		require.NoError(t, err)
		require.NotNil(t, newConfig)

		authHandler := auth.NewMockHandler()
		defer authHandler.AssertExpectations(t)

		authHandler.On("RegisterService", mock.Anything)

		dataHandler := data.NewMockHandler()
		defer dataHandler.AssertExpectations(t)

		dataHandler.On("RegisterService", mock.Anything)

		lifecycle := fxtest.NewLifecycle(t)

		l, err := logger.New(logger.Params{
			Lifecycle: lifecycle,
			Config:    newConfig,
		})
		require.NoError(t, err)
		require.NotNil(t, l)

		err = New(Params{
			Lifecycle:   lifecycle,
			Config:      newConfig,
			Logger:      l,
			AuthHandler: authHandler,
			DataHandler: dataHandler,
		})
		require.NoError(t, err, "expected successful gRPC server initialization, got error")
	})

	t.Run("lifecycle start and stop", func(t *testing.T) {
		port, err := getRandomPort()
		require.NoError(t, err)

		err = os.Setenv("GRPC_ADDRESS", "localhost:"+strconv.Itoa(port))
		require.NoError(t, err)

		newConfig, err := config.NewConfig()
		require.NoError(t, err)
		require.NotNil(t, newConfig)

		authHandler := auth.NewMockHandler()
		defer authHandler.AssertExpectations(t)

		authHandler.On("RegisterService", mock.Anything)

		dataHandler := data.NewMockHandler()
		defer dataHandler.AssertExpectations(t)

		dataHandler.On("RegisterService", mock.Anything)

		lifecycle := fxtest.NewLifecycle(t)

		err = New(Params{
			Lifecycle:   lifecycle,
			Config:      newConfig,
			Logger:      &nopLogger{},
			AuthHandler: authHandler,
			DataHandler: dataHandler,
		})
		require.NoError(t, err, "expected successful gRPC server initialization, got error")

		err = lifecycle.Start(t.Context())
		require.NoError(t, err, "expected successful lifecycle start, got error")

		err = lifecycle.Stop(t.Context())
		require.NoError(t, err, "expected successful lifecycle stop, got error")
	})
}

type nopLogger struct{}

func (n *nopLogger) Error(msg string, fields ...interface{}) {}
func (n *nopLogger) Info(msg string, fields ...interface{})  {}
func (n *nopLogger) Debug(msg string, fields ...interface{}) {}
func (n *nopLogger) Warn(msg string, fields ...interface{})  {}

func getRandomPort() (int, error) {
	b := make([]byte, 2)
	_, err := rand.Read(b)
	if err != nil {
		return 0, fmt.Errorf("failed to generate random port: %w", err)
	}
	// Диапазон портов: 50000-50999
	return 50000 + int(b[0])%10*100 + int(b[1])%100, nil
}
