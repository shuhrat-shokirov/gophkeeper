package data

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gophkeeper/internal/client/gateways/server"
	"gophkeeper/internal/client/services/auth"
)

func Test_service_SaveText(t *testing.T) {
	t.Run("err from get user id", func(t *testing.T) {
		authService := auth.NewMockService()
		defer authService.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(0, assert.AnError)

		s := &service{
			authService: authService,
		}

		err := s.SaveText(t.Context(), &TextData{})
		assert.Error(t, err, "expected error from GetUserID, got nil")
	})

	t.Run("err from save text", func(t *testing.T) {
		authService := auth.NewMockService()
		defer authService.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(1, nil)

		gateway := server.NewMockGateway()
		defer gateway.AssertExpectations(t)

		gateway.On("SaveText", mock.Anything, mock.Anything, mock.Anything).Return(assert.AnError)

		s := &service{
			authService:   authService,
			serverGateway: gateway,
		}

		err := s.SaveText(t.Context(), &TextData{})
		assert.Error(t, err, "expected error from SaveText, got nil")
	})

	t.Run("success", func(t *testing.T) {
		authService := auth.NewMockService()
		defer authService.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(1, nil)

		gateway := server.NewMockGateway()
		defer gateway.AssertExpectations(t)

		gateway.On("SaveText", mock.Anything, mock.Anything, mock.Anything).
			Return(nil)

		s := &service{
			authService:   authService,
			serverGateway: gateway,
		}

		err := s.SaveText(t.Context(), &TextData{})
		assert.NoError(t, err, "expected no error from SaveText")
	})
}

func Test_service_GetTextList(t *testing.T) {
	t.Run("error from GetUserID", func(t *testing.T) {
		authService := auth.NewMockService()
		defer authService.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(0, assert.AnError)

		s := &service{
			authService: authService,
		}

		list, err := s.GetTextList(t.Context(), 10, 0)
		assert.Error(t, err)
		assert.Nil(t, list)
	})

	t.Run("error from gateway", func(t *testing.T) {
		authService := auth.NewMockService()
		gateway := server.NewMockGateway()
		defer authService.AssertExpectations(t)
		defer gateway.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(1, nil)
		gateway.On("GetTextList", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		s := &service{
			authService:   authService,
			serverGateway: gateway,
		}

		list, err := s.GetTextList(t.Context(), 10, 0)
		assert.Error(t, err)
		assert.Nil(t, list)
	})

	t.Run("success", func(t *testing.T) {
		authService := auth.NewMockService()
		gateway := server.NewMockGateway()
		defer authService.AssertExpectations(t)
		defer gateway.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(1, nil)

		gateway.On("GetTextList", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return([]server.ListItem{
				{
					ID:        1,
					Title:     "Text 1",
					CreatedAt: time.Now().UnixNano(),
					UpdatedAt: time.Now().UnixNano(),
				},
			}, nil)

		s := &service{
			authService:   authService,
			serverGateway: gateway,
		}

		list, err := s.GetTextList(t.Context(), 10, 0)
		assert.NoError(t, err)
		assert.Len(t, list, 1)
		assert.Equal(t, "Text 1", list[0].Title)
	})
}

func Test_service_GetTextByID(t *testing.T) {
	t.Run("error from GetUserID", func(t *testing.T) {
		authService := auth.NewMockService()
		defer authService.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(0, assert.AnError)

		s := &service{
			authService: authService,
		}

		text, err := s.GetTextByID(t.Context(), 1)
		assert.Error(t, err)
		assert.Nil(t, text)
	})

	t.Run("error from gateway", func(t *testing.T) {
		authService := auth.NewMockService()
		gateway := server.NewMockGateway()
		defer authService.AssertExpectations(t)
		defer gateway.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(1, nil)
		gateway.On("GetTextByID", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		s := &service{
			authService:   authService,
			serverGateway: gateway,
		}

		text, err := s.GetTextByID(t.Context(), 1)
		assert.Error(t, err)
		assert.Nil(t, text)
	})

	t.Run("success", func(t *testing.T) {
		authService := auth.NewMockService()
		gateway := server.NewMockGateway()
		defer authService.AssertExpectations(t)
		defer gateway.AssertExpectations(t)

		created := time.Now()
		updated := time.Now()

		authService.On("GetUserID", mock.Anything).Return(1, nil)
		gateway.On("GetTextByID", mock.Anything, mock.Anything, mock.Anything).
			Return(&server.TextInfo{
				ID: 1,
				Text: server.Text{
					Title:     "Title",
					Content:   "Content",
					Note:      "Note",
					CreatedAt: created,
				},
				UpdatedAt: updated,
			}, nil)

		s := &service{
			authService:   authService,
			serverGateway: gateway,
		}

		text, err := s.GetTextByID(t.Context(), 1)
		assert.NoError(t, err)
		assert.NotNil(t, text)
		assert.Equal(t, "Title", text.Title)
		assert.Equal(t, "Content", text.Content)
		assert.Equal(t, "Note", text.Note)
		assert.Equal(t, created, text.CreatedAt)
		assert.Equal(t, updated, text.UpdatedAt)
	})
}
