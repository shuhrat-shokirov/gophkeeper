//nolint:errcheck,gocritic
package auth

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"gophkeeper/internal/server/errorx"
	"gophkeeper/internal/server/repositories/user"
)

func Test_service_ConfirmOTP(t *testing.T) {
	t.Run("not found on cache", func(t *testing.T) {
		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		cache.On("Find", mock.Anything, "otp:otp_id", mock.Anything).Return(errorx.ErrNotFound)

		_, err := service.ConfirmOTP(t.Context(), "otp_id", "123456")
		require.Error(t, err)
		require.ErrorIs(t, err, errorx.ErrOTPExpired)
	})

	t.Run("err from cache", func(t *testing.T) {
		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		cache.On("Find", mock.Anything, "otp:otp_id", mock.Anything).Return(assert.AnError)

		_, err := service.ConfirmOTP(t.Context(), "otp_id", "123456")
		require.Error(t, err)
	})

	t.Run("otp mismatch", func(t *testing.T) {
		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		cache.On("Find", mock.Anything, "otp:otp_id", mock.Anything).
			Return(nil, &otpVerification{Otp: "1234567"})

		_, err := service.ConfirmOTP(t.Context(), "otp_id", "123456")
		require.Error(t, err)
		require.ErrorIs(t, err, errorx.ErrInvalidOTP)
	})

	t.Run("err from user repository", func(t *testing.T) {
		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		cache.On("Find", mock.Anything, mock.Anything, mock.Anything).
			Run(func(args mock.Arguments) {
				dest := args.Get(2).(*otpVerification)
				dest.Otp = "123456"
			}).
			Return(nil, nil)

		userRepo.On("GetUserByID", mock.Anything, mock.Anything).
			Return(nil, errorx.ErrNotFound)

		_, err := service.ConfirmOTP(t.Context(), "otp_id", "123456")
		require.Error(t, err)
		require.ErrorIs(t, err, errorx.ErrNotFound)
	})

	t.Run("err from generate token pair", func(t *testing.T) {
		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		cache.On("Find", mock.Anything, mock.Anything, mock.Anything).
			Run(func(args mock.Arguments) {
				dest := args.Get(2).(*otpVerification)
				dest.Otp = "123456"
			}).
			Return(nil, nil)

		userRepo.On("GetUserByID", mock.Anything, mock.Anything).
			Return(&user.User{ID: 1}, nil)

		_, err := service.ConfirmOTP(t.Context(), "otp_id", "123456")
		require.Error(t, err)
	})

	t.Run("err from create session", func(t *testing.T) {
		err := os.Setenv("JWT_PRIVATE_KEY", testPrivateKey)
		require.NoError(t, err)

		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		cache.On("Find", mock.Anything, mock.Anything, mock.Anything).
			Run(func(args mock.Arguments) {
				dest := args.Get(2).(*otpVerification)
				dest.Otp = "123456"
			}).
			Return(nil, nil)

		userRepo.On("GetUserByID", mock.Anything, mock.Anything).
			Return(&user.User{ID: 1}, nil)

		sessionRepo.On("Create", mock.Anything, mock.Anything).
			Return(assert.AnError)

		_, err = service.ConfirmOTP(t.Context(), "otp_id", "123456")
		require.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		err := os.Setenv("JWT_PRIVATE_KEY", testPrivateKey)
		require.NoError(t, err)

		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		cache.On("Find", mock.Anything, mock.Anything, mock.Anything).
			Run(func(args mock.Arguments) {
				dest := args.Get(2).(*otpVerification)
				dest.Otp = "123456"
				dest.UserID = 1
			}).
			Return(nil, nil)

		userRepo.On("GetUserByID", mock.Anything, mock.Anything).
			Return(&user.User{ID: 1}, nil)

		sessionRepo.On("Create", mock.Anything, mock.Anything).
			Return(nil)

		cache.On("Delete", mock.Anything, "otp:otp_id").Return(nil)

		token, err := service.ConfirmOTP(t.Context(), "otp_id", "123456")
		require.NoError(t, err)
		assert.NotEmpty(t, token.Token)
		assert.NotEmpty(t, token.RefreshToken)
		assert.Equal(t, 1, token.UserId)
	})
}
