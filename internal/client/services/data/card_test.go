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

func Test_service_SaveCard(t *testing.T) {
	t.Run("error from GetUserID", func(t *testing.T) {
		authService := auth.NewMockService()
		defer authService.AssertExpectations(t)

		s := &service{authService: authService}

		authService.On("GetUserID", mock.Anything).Return(0, assert.AnError)

		err := s.SaveCard(t.Context(), &CardData{})
		assert.Error(t, err)
	})

	t.Run("error from SaveCard", func(t *testing.T) {
		authService := auth.NewMockService()
		gateway := server.NewMockGateway()
		defer authService.AssertExpectations(t)
		defer gateway.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(1, nil)
		gateway.On("SaveCard", mock.Anything, mock.Anything, mock.Anything).Return(assert.AnError)

		s := &service{authService: authService, serverGateway: gateway}

		err := s.SaveCard(t.Context(), &CardData{})
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		authService := auth.NewMockService()
		gateway := server.NewMockGateway()
		defer authService.AssertExpectations(t)
		defer gateway.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(1, nil)
		gateway.On("SaveCard", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		s := &service{authService: authService, serverGateway: gateway}

		err := s.SaveCard(t.Context(), &CardData{})
		assert.NoError(t, err)
	})
}

func Test_service_GetCardList(t *testing.T) {
	t.Run("error from GetUserID", func(t *testing.T) {
		authService := auth.NewMockService()
		defer authService.AssertExpectations(t)

		s := &service{authService: authService}

		authService.On("GetUserID", mock.Anything).Return(0, assert.AnError)

		list, err := s.GetCardList(t.Context(), 10, 0)
		assert.Error(t, err)
		assert.Nil(t, list)
	})

	t.Run("error from gateway", func(t *testing.T) {
		authService := auth.NewMockService()
		gateway := server.NewMockGateway()
		defer authService.AssertExpectations(t)
		defer gateway.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(1, nil)
		gateway.On("GetCardList", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, assert.AnError)

		s := &service{authService: authService, serverGateway: gateway}

		list, err := s.GetCardList(t.Context(), 10, 0)
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
		gateway.On("GetCardList", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]server.ListItem{
			{ID: 1, Title: "Card 1", CreatedAt: now, UpdatedAt: now},
		}, nil)

		s := &service{authService: authService, serverGateway: gateway}

		list, err := s.GetCardList(t.Context(), 10, 0)
		assert.NoError(t, err)
		assert.Len(t, list, 1)
		assert.Equal(t, "Card 1", list[0].Title)
	})
}

func Test_service_GetCardByID(t *testing.T) {
	t.Run("error from GetUserID", func(t *testing.T) {
		authService := auth.NewMockService()
		defer authService.AssertExpectations(t)

		s := &service{authService: authService}

		authService.On("GetUserID", mock.Anything).Return(0, assert.AnError)

		info, err := s.GetCardByID(t.Context(), 1)
		assert.Error(t, err)
		assert.Nil(t, info)
	})

	t.Run("error from gateway", func(t *testing.T) {
		authService := auth.NewMockService()
		gateway := server.NewMockGateway()
		defer authService.AssertExpectations(t)
		defer gateway.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(1, nil)
		gateway.On("GetCardByID", mock.Anything, mock.Anything, mock.Anything).Return(nil, assert.AnError)

		s := &service{authService: authService, serverGateway: gateway}

		info, err := s.GetCardByID(t.Context(), 1)
		assert.Error(t, err)
		assert.Nil(t, info)
	})

	t.Run("success", func(t *testing.T) {
		authService := auth.NewMockService()
		gateway := server.NewMockGateway()
		defer authService.AssertExpectations(t)
		defer gateway.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(1, nil)
		gateway.On("GetCardByID", mock.Anything, mock.Anything, mock.Anything).Return(&server.CardInfo{
			ID: 1,
			Card: server.Card{
				Title:  "Visa",
				Pan:    "4111111111111111",
				Cvv:    "123",
				Expiry: "12/24",
				Note:   "My card",
			},
		}, nil)

		s := &service{authService: authService, serverGateway: gateway}

		info, err := s.GetCardByID(t.Context(), 1)
		assert.NoError(t, err)
		assert.NotNil(t, info)
		assert.Equal(t, "Visa", info.Title)
		assert.Equal(t, "4111111111111111", info.Pan)
		assert.Equal(t, "123", info.Cvv)
		assert.Equal(t, "12/24", info.Expiry)
		assert.Equal(t, "My card", info.Note)
		assert.Equal(t, int64(1), info.ID)
	})
}
