package auth

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gophkeeper/internal/client/errorx"
	"gophkeeper/internal/client/gateways/server"
	"gophkeeper/pkg/memorycache"
)

func Test_service_ConfirmOTP(t *testing.T) {
	t.Run("can't find on cache", func(t *testing.T) {
		cache := memorycache.New(memorycache.Params{})
		s := &service{
			cache: cache,
		}

		err := s.ConfirmOTP(t.Context(), "123456")
		assert.Error(t, err)
		assert.Equal(t, err, errorx.ErrOtpExpired)
	})

	t.Run("err from gateway", func(t *testing.T) {
		serv := server.NewMockGateway()
		defer serv.AssertExpectations(t)

		cache := memorycache.New(memorycache.Params{})

		s := &service{
			serverGateway: serv,
			cache:         cache,
		}

		cache.Set(otpCodeKey, "test-otp-id", time.Second)

		serv.On("ConfirmOtp", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		err := s.ConfirmOTP(t.Context(), "123456")
		assert.Error(t, err, "expected error from gateway, got nil")
	})

	t.Run("success", func(t *testing.T) {
		serv := server.NewMockGateway()
		defer serv.AssertExpectations(t)

		cache := memorycache.New(memorycache.Params{})

		s := &service{
			serverGateway: serv,
			cache:         cache,
		}

		cache.Set(otpCodeKey, "test-otp-id", time.Second)

		serv.On("ConfirmOtp", mock.Anything, mock.Anything, mock.Anything).
			Return(&server.Token{
				AccessToken:  "123",
				RefreshToken: "345",
			}, nil)

		err := s.ConfirmOTP(t.Context(), "123456")
		assert.NoError(t, err, "expected successful OTP confirmation, got error")

		assert.Equal(t, s.accessToken, "123")
		assert.Equal(t, s.refreshToken, "345")
	})
}

func Test_writeEnvFile(t *testing.T) {
	err := writeEnvFile("123", "345")
	assert.NoError(t, err, "expected no error writing env file")

	defer func() {
		_ = os.Remove(".env")
	}()

	bytes, err := os.ReadFile(".env")
	assert.NoError(t, err, "expected no error reading env file")

	assert.Contains(t, string(bytes), "GOPH_KEEPER_ACCESS_TOKEN=123")
	assert.Contains(t, string(bytes), "GOPH_KEEPER_REFRESH_TOKEN=345")
}
