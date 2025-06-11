//nolint:dupl,gocritic
package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gophkeeper/internal/server/errorx"
	"gophkeeper/internal/server/repositories/user"
)

func Test_service_Registration(t *testing.T) {
	t.Run("get user error", func(t *testing.T) {
		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		userRepo.On("GetUserByEmail", t.Context(), mock.Anything).
			Return(nil, assert.AnError)

		registration, err := service.Registration(t.Context(), Registration{})
		assert.Error(t, err)
		assert.Empty(t, registration)
	})

	t.Run("user already exists", func(t *testing.T) {
		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		userRepo.On("GetUserByEmail", t.Context(), mock.Anything).
			Return(&user.User{ID: 1}, nil)

		registration, err := service.Registration(t.Context(), Registration{})
		assert.Error(t, err)
		assert.Empty(t, registration)
		assert.ErrorIs(t, err, errorx.ErrAlreadyExists)
	})

	t.Run("err from create user", func(t *testing.T) {
		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		userRepo.On("GetUserByEmail", t.Context(), mock.Anything).
			Return(nil, nil)
		userRepo.On("CreateUser", t.Context(), mock.Anything).
			Return(0, assert.AnError)

		registration, err := service.Registration(t.Context(), Registration{})
		assert.Error(t, err)
		assert.Empty(t, registration)
	})

	t.Run("err from email gateway", func(t *testing.T) {
		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		userRepo.On("GetUserByEmail", t.Context(), mock.Anything).
			Return(nil, nil)
		userRepo.On("CreateUser", t.Context(), mock.Anything).
			Return(1, nil)
		emailGateway.On("SendEmail", mock.Anything, mock.Anything).
			Return(assert.AnError)

		registration, err := service.Registration(t.Context(), Registration{})
		assert.Error(t, err)
		assert.Empty(t, registration)
	})

	t.Run("success", func(t *testing.T) {
		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		userRepo.On("GetUserByEmail", t.Context(), mock.Anything).
			Return(nil, nil)
		userRepo.On("CreateUser", t.Context(), mock.Anything).
			Return(1, nil)
		emailGateway.On("SendEmail", mock.Anything, mock.Anything).
			Return(nil)
		cache.On("Save", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)

		registration, err := service.Registration(t.Context(), Registration{})
		assert.NoError(t, err)
		assert.NotEmpty(t, registration)
	})
}

func Test_service_senOtp(t *testing.T) {
	t.Run("err from email gateway", func(t *testing.T) {
		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		userRepo.On("GetUserByEmail", t.Context(), mock.Anything).
			Return(nil, nil)
		userRepo.On("CreateUser", t.Context(), mock.Anything).
			Return(1, nil)
		emailGateway.On("SendEmail", mock.Anything, mock.Anything).
			Return(assert.AnError)

		otpID, err := service.Registration(t.Context(), Registration{})
		assert.Error(t, err)
		assert.Empty(t, otpID)
	})

	t.Run("err on cache set", func(t *testing.T) {
		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		userRepo.On("GetUserByEmail", t.Context(), mock.Anything).
			Return(nil, nil)
		userRepo.On("CreateUser", t.Context(), mock.Anything).
			Return(1, nil)
		emailGateway.On("SendEmail", mock.Anything, mock.Anything).
			Return(nil)
		cache.On("Save", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(assert.AnError)

		otpID, err := service.Registration(t.Context(), Registration{})
		assert.Error(t, err)
		assert.Empty(t, otpID)
	})

	t.Run("success", func(t *testing.T) {
		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		userRepo.On("GetUserByEmail", t.Context(), mock.Anything).
			Return(nil, nil)
		userRepo.On("CreateUser", t.Context(), mock.Anything).
			Return(1, nil)
		emailGateway.On("SendEmail", mock.Anything, mock.Anything).
			Return(nil)
		cache.On("Save", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)

		otpID, err := service.Registration(t.Context(), Registration{})
		assert.NoError(t, err)
		assert.NotEmpty(t, otpID)
	})
}
