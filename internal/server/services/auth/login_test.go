package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"gophkeeper/internal/server/errorx"
	"gophkeeper/internal/server/repositories/user"
)

func Test_service_Login(t *testing.T) {
	t.Run("err from user repository", func(t *testing.T) {
		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		userRepo.On("GetUserByEmail", t.Context(),
			mock.Anything).Return(nil, errorx.ErrNotFound)

		_, err := service.Login(t.Context(), "", "")
		require.Error(t, err)
		require.ErrorIs(t, err, errorx.ErrNotFound)
	})

	t.Run("password mismatch", func(t *testing.T) {
		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		userRepo.On("GetUserByEmail", t.Context(),
			mock.Anything).Return(&user.User{Password: "hashed_password"}, nil)

		_, err := service.Login(t.Context(), "", "wrong_password")
		require.Error(t, err)
		require.ErrorIs(t, err, errorx.ErrInvalidCredentials)
	})

	t.Run("send email err", func(t *testing.T) {
		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		userRepo.On("GetUserByEmail", t.Context(),
			mock.Anything).Return(&user.User{ID: 1,
			Password: "$2a$10$7vtvuhvGGAXib9qc0a2lbOVkk1o3oaeR7qSTrMsw6taSVxIkfGaWW"}, nil)

		emailGateway.On("SendEmail", mock.Anything).Return(assert.AnError)

		_, err := service.Login(t.Context(), "", "qweqwe")
		require.Error(t, err)
	})
	t.Run("successful login", func(t *testing.T) {
		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		userRepo.On("GetUserByEmail", t.Context(),
			mock.Anything).Return(&user.User{ID: 1,
			Password: "$2a$10$7vtvuhvGGAXib9qc0a2lbOVkk1o3oaeR7qSTrMsw6taSVxIkfGaWW"}, nil)

		emailGateway.On("SendEmail", mock.Anything).Return(nil)
		cache.On("Save", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

		otpId, err := service.Login(t.Context(), "", "qweqwe")
		require.NoError(t, err)
		require.NotEmpty(t, otpId)
	})
}
