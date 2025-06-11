//nolint:dupl,gocritic
package data

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gophkeeper/internal/client/gateways/server"
	"gophkeeper/internal/client/services/auth"
)

func Test_service_SaveLogin(t *testing.T) {
	t.Run("error from GetUserID", func(t *testing.T) {
		authService := auth.NewMockService()
		defer authService.AssertExpectations(t)

		s := &service{authService: authService}

		authService.On("GetUserID", mock.Anything).Return(0, assert.AnError)

		err := s.SaveLogin(t.Context(), &LoginData{})
		assert.Error(t, err)
	})

	t.Run("error from SaveLoginAndPass", func(t *testing.T) {
		authService := auth.NewMockService()
		gateway := server.NewMockGateway()
		defer authService.AssertExpectations(t)
		defer gateway.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(1, nil)
		gateway.On("SaveLoginAndPass", mock.Anything, mock.Anything, mock.Anything).Return(assert.AnError)

		s := &service{authService: authService, serverGateway: gateway}

		err := s.SaveLogin(t.Context(), &LoginData{})
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		authService := auth.NewMockService()
		gateway := server.NewMockGateway()
		defer authService.AssertExpectations(t)
		defer gateway.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(1, nil)
		gateway.On("SaveLoginAndPass", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		s := &service{authService: authService, serverGateway: gateway}

		err := s.SaveLogin(t.Context(), &LoginData{})
		assert.NoError(t, err)
	})
}

func Test_service_GetLoginList(t *testing.T) {
	t.Run("error from GetUserID", func(t *testing.T) {
		authService := auth.NewMockService()
		defer authService.AssertExpectations(t)

		s := &service{authService: authService}

		authService.On("GetUserID", mock.Anything).Return(0, assert.AnError)

		list, err := s.GetLoginList(t.Context(), 10, 0)
		assert.Error(t, err)
		assert.Nil(t, list)
	})

	t.Run("error from gateway", func(t *testing.T) {
		authService := auth.NewMockService()
		gateway := server.NewMockGateway()
		defer authService.AssertExpectations(t)
		defer gateway.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(1, nil)
		gateway.On("GetLoginList", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, assert.AnError)

		s := &service{authService: authService, serverGateway: gateway}

		list, err := s.GetLoginList(t.Context(), 10, 0)
		assert.Error(t, err)
		assert.Nil(t, list)
	})

	t.Run("success", func(t *testing.T) {
		authService := auth.NewMockService()
		gateway := server.NewMockGateway()
		defer authService.AssertExpectations(t)
		defer gateway.AssertExpectations(t)

		now := time.Now().UnixNano()
		authService.On("GetUserID", mock.Anything).Return(1, nil)
		gateway.On("GetLoginList", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]server.ListItem{
			{ID: 1, Title: "Login 1", CreatedAt: now, UpdatedAt: now},
		}, nil)

		s := &service{authService: authService, serverGateway: gateway}

		list, err := s.GetLoginList(t.Context(), 10, 0)
		assert.NoError(t, err)
		assert.Len(t, list, 1)
		assert.Equal(t, "Login 1", list[0].Title)
	})
}

func Test_service_GetLoginByID(t *testing.T) {
	t.Run("error from GetUserID", func(t *testing.T) {
		authService := auth.NewMockService()
		defer authService.AssertExpectations(t)

		s := &service{authService: authService}

		authService.On("GetUserID", mock.Anything).Return(0, assert.AnError)

		info, err := s.GetLoginByID(t.Context(), 1)
		assert.Error(t, err)
		assert.Nil(t, info)
	})

	t.Run("error from gateway", func(t *testing.T) {
		authService := auth.NewMockService()
		gateway := server.NewMockGateway()
		defer authService.AssertExpectations(t)
		defer gateway.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(1, nil)
		gateway.On("GetLoginByID", mock.Anything, mock.Anything, mock.Anything).Return(nil, assert.AnError)

		s := &service{authService: authService, serverGateway: gateway}

		info, err := s.GetLoginByID(t.Context(), 1)
		assert.Error(t, err)
		assert.Nil(t, info)
	})

	t.Run("success", func(t *testing.T) {
		authService := auth.NewMockService()
		gateway := server.NewMockGateway()
		defer authService.AssertExpectations(t)
		defer gateway.AssertExpectations(t)

		now := time.Now()
		authService.On("GetUserID", mock.Anything).Return(1, nil)
		gateway.On("GetLoginByID", mock.Anything, mock.Anything, mock.Anything).Return(&server.LoginInfo{
			ID: 1,
			LoginAndPass: server.LoginAndPass{
				Login: "user1",
				Pass:  "secret",
				Title: "My login",
				Note:  "some note",
			},
			UpdatedAt: now,
		}, nil)

		s := &service{authService: authService, serverGateway: gateway}

		info, err := s.GetLoginByID(t.Context(), 1)
		assert.NoError(t, err)
		assert.NotNil(t, info)
		assert.Equal(t, "user1", info.Login)
		assert.Equal(t, "secret", info.Pass)
		assert.Equal(t, "My login", info.Title)
		assert.Equal(t, "some note", info.Note)
		assert.Equal(t, int64(1), info.ID)
	})
}
