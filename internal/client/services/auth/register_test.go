package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx/fxtest"

	"gophkeeper/internal/client/errorx"
	"gophkeeper/internal/client/gateways/server"
	"gophkeeper/pkg/memorycache"
)

func Test_service_Register(t *testing.T) {
	t.Run("invalid email format", func(t *testing.T) {
		s := &service{}
		err := s.Register(t.Context(), "invalid-email", "validPassword123")
		assert.Error(t, err, "expected error for invalid email format")
		assert.Equal(t, err, errorx.ErrEmailInvalidFormat)
	})

	t.Run("invalid password", func(t *testing.T) {
		s := &service{}
		err := s.Register(t.Context(), "123@gmail.com", "short")
		assert.Error(t, err, "expected error for invalid password")
		assert.Equal(t, err, errorx.ErrPasswordTooShort)
	})

	t.Run("err from gateway", func(t *testing.T) {
		serv := server.NewMockGateway()
		defer serv.AssertExpectations(t)

		s, err := New(Params{
			Lifecycle:     fxtest.NewLifecycle(t),
			ServerGateway: serv,
		})
		require.NoError(t, err)
		require.NotNil(t, s)

		serv.On("Register", t.Context(), mock.Anything, mock.Anything).Return("", assert.AnError)

		err = s.Register(t.Context(), "123@gmail.com", "validPassword123")
		assert.Error(t, err, "expected error from gateway, got nil")
	})

	t.Run("success", func(t *testing.T) {
		serv := server.NewMockGateway()
		defer serv.AssertExpectations(t)

		cache := memorycache.New(memorycache.Params{})

		s, err := New(Params{
			Lifecycle:     fxtest.NewLifecycle(t),
			ServerGateway: serv,
			Cache:         cache,
		})
		require.NoError(t, err)
		require.NotNil(t, s)

		expectedOtpID := "otp123"
		serv.On("Register", t.Context(), mock.Anything, mock.Anything).Return(expectedOtpID, nil)

		err = s.Register(t.Context(), "123@gmail.com", "validPassword123")
		assert.NoError(t, err, "expected no error from gateway")

		get, b := cache.Get(otpCodeKey)
		require.True(t, b, "expected otpCodeKey to be in cache")
		assert.Equal(t, expectedOtpID, get, "expected otpCodeKey to match the returned otpID")
	})
}
