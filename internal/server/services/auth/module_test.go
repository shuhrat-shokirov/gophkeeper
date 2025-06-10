//nolint:gocritic
package auth

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/fx/fxtest"

	"gophkeeper/internal/server/gateways/emailtotp"
	"gophkeeper/internal/server/repositories/session"
	"gophkeeper/internal/server/repositories/user"
	"gophkeeper/pkg/config"
	"gophkeeper/pkg/jwt"
	"gophkeeper/pkg/logger"
	"gophkeeper/pkg/redis"
)

func TestNew(t *testing.T) {
	s := New(Params{})
	require.NotNil(t, s)
}

func testInit(t *testing.T) (Service, *redis.MockRedis,
	*emailtotp.MockGateway, *user.MockRepository, *session.MockRepository) {
	t.Helper()

	newConfig, err := config.NewConfig()
	require.NoError(t, err)
	require.NotNil(t, newConfig)

	lifecycle := fxtest.NewLifecycle(t)
	require.NotNil(t, lifecycle)

	l, err := logger.New(logger.Params{
		Lifecycle: lifecycle,
		Config:    newConfig,
	})
	require.NoError(t, err)
	require.NotNil(t, l)

	j, err := jwt.New(jwt.Params{
		Config: newConfig,
	})
	require.NoError(t, err)
	require.NotNil(t, j)

	mockRedis := redis.NewMockRedis()

	emailGateway := emailtotp.NewMockGateway()

	userRepo := user.NewMockRepository()

	sessionRepo := session.NewMockRepository()

	s := New(Params{
		Cache:            mockRedis,
		EmailTotpGateway: emailGateway,
		JWT:              j,
		UserRepo:         userRepo,
		SessionRepo:      sessionRepo,
	})
	require.NotNil(t, s)

	return s, mockRedis, emailGateway, userRepo, sessionRepo
}
