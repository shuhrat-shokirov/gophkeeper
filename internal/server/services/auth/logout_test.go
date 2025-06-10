package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_service_Logout(t *testing.T) {
	t.Run("empty refresh token", func(t *testing.T) {
		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		err := service.Logout(t.Context(), "")
		if err == nil {
			t.Error("expected error for empty refresh token, got nil")
		}
	})

	t.Run("error on session delete", func(t *testing.T) {
		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		sessionRepo.On("Delete", t.Context(), "refreshToken").
			Return(assert.AnError)

		err := service.Logout(t.Context(), "refreshToken")
		if err == nil {
			t.Error("expected error on session delete, got nil")
		}
	})

	t.Run("successful logout", func(t *testing.T) {
		service, cache, emailGateway, userRepo, sessionRepo := testInit(t)
		defer cache.AssertExpectations(t)
		defer emailGateway.AssertExpectations(t)
		defer userRepo.AssertExpectations(t)
		defer sessionRepo.AssertExpectations(t)

		sessionRepo.On("Delete", t.Context(), "refreshToken").
			Return(nil)

		err := service.Logout(t.Context(), "refreshToken")
		assert.NoError(t, err, "expected successful logout, got error")
	})
}
