package data

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gophkeeper/internal/client/gateways/server"
	"gophkeeper/internal/client/services/auth"
)

func Test_service_SaveFile(t *testing.T) {
	t.Run("ошибка чтения файла", func(t *testing.T) {
		s := &service{}

		err := s.SaveFile(t.Context(), &FileData{Path: "non-existent.txt"})
		assert.Error(t, err)
	})

	t.Run("ошибка от GetUserID", func(t *testing.T) {
		authService := auth.NewMockService()
		defer authService.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(0, assert.AnError)

		s := &service{authService: authService}

		// Чтобы ReadFile не падал, создаём временный файл
		tmpFile, _ := os.CreateTemp("", "example.txt")
		defer func() {
			_ = os.Remove(tmpFile.Name())
		}()

		err := s.SaveFile(t.Context(), &FileData{Path: tmpFile.Name()})
		assert.Error(t, err)
	})

	t.Run("ошибка от SaveBinary", func(t *testing.T) {
		authService := auth.NewMockService()
		gateway := server.NewMockGateway()
		defer authService.AssertExpectations(t)
		defer gateway.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(int64(1), nil)
		gateway.On("SaveBinary", mock.Anything, mock.Anything, mock.Anything).Return(assert.AnError)

		// Создаём временный файл
		tmpFile, _ := os.CreateTemp("", "binary.txt")
		defer func() {
			_ = os.Remove(tmpFile.Name())
			_ = tmpFile.Close()
		}()

		_, err := tmpFile.WriteString("test content")
		assert.NoError(t, err)

		s := &service{authService: authService, serverGateway: gateway}

		err = s.SaveFile(t.Context(), &FileData{
			Path:  tmpFile.Name(),
			Title: "file",
			Note:  "test",
		})
		assert.Error(t, err)
	})

	t.Run("успешный путь", func(t *testing.T) {
		authService := auth.NewMockService()
		gateway := server.NewMockGateway()
		defer authService.AssertExpectations(t)
		defer gateway.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(int64(1), nil)
		gateway.On("SaveBinary", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		tmpFile, _ := os.CreateTemp("", "binary.txt")
		defer func() {
			_ = os.Remove(tmpFile.Name())
			_ = tmpFile.Close()
		}()

		_, err := tmpFile.WriteString("valid content")
		assert.NoError(t, err)

		s := &service{authService: authService, serverGateway: gateway}

		err = s.SaveFile(t.Context(), &FileData{
			Path:  tmpFile.Name(),
			Title: "file",
			Note:  "test",
		})
		assert.NoError(t, err)
	})
}

func Test_service_GetBinaryList(t *testing.T) {
	t.Run("ошибка от GetUserID", func(t *testing.T) {
		authService := auth.NewMockService()
		defer authService.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(int64(0), assert.AnError)

		s := &service{authService: authService}

		list, err := s.GetBinaryList(t.Context(), 10, 0)
		assert.Error(t, err)
		assert.Nil(t, list)
	})

	t.Run("ошибка от gateway", func(t *testing.T) {
		authService := auth.NewMockService()
		gateway := server.NewMockGateway()
		defer authService.AssertExpectations(t)
		defer gateway.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(int64(1), nil)
		gateway.On("GetBinaryList", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("gateway error"))

		s := &service{authService: authService, serverGateway: gateway}

		list, err := s.GetBinaryList(t.Context(), 10, 0)
		assert.Error(t, err)
		assert.Nil(t, list)
	})

	t.Run("успешный путь", func(t *testing.T) {
		authService := auth.NewMockService()
		gateway := server.NewMockGateway()
		defer authService.AssertExpectations(t)
		defer gateway.AssertExpectations(t)

		now := time.Now().UnixNano()

		authService.On("GetUserID", mock.Anything).Return(int64(1), nil)
		gateway.On("GetBinaryList", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return([]server.ListItem{
				{ID: 1, Title: "file.bin", CreatedAt: now, UpdatedAt: now},
			}, nil)

		s := &service{authService: authService, serverGateway: gateway}

		list, err := s.GetBinaryList(t.Context(), 10, 0)
		assert.NoError(t, err)
		assert.Len(t, list, 1)
		assert.Equal(t, "file.bin", list[0].Title)
	})
}

func Test_service_GetBinaryByID(t *testing.T) {
	t.Run("ошибка от GetUserID", func(t *testing.T) {
		authService := auth.NewMockService()
		defer authService.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(int64(0), assert.AnError)

		s := &service{authService: authService}

		info, err := s.GetBinaryByID(t.Context(), 1)
		assert.Error(t, err)
		assert.Nil(t, info)
	})

	t.Run("ошибка от gateway", func(t *testing.T) {
		authService := auth.NewMockService()
		gateway := server.NewMockGateway()
		defer authService.AssertExpectations(t)
		defer gateway.AssertExpectations(t)

		authService.On("GetUserID", mock.Anything).Return(int64(1), nil)
		gateway.On("GetBinaryByID", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("not found"))

		s := &service{authService: authService, serverGateway: gateway}

		info, err := s.GetBinaryByID(t.Context(), 1)
		assert.Error(t, err)
		assert.Nil(t, info)
	})

	t.Run("успешный путь", func(t *testing.T) {
		authService := auth.NewMockService()
		gateway := server.NewMockGateway()
		defer authService.AssertExpectations(t)
		defer gateway.AssertExpectations(t)

		now := time.Now()
		authService.On("GetUserID", mock.Anything).Return(int64(1), nil)
		gateway.On("GetBinaryByID", mock.Anything, mock.Anything, mock.Anything).
			Return(&server.BinaryInfo{
				ID: 1,
				Binary: server.Binary{
					Title:     "My file",
					Content:   []byte("abc"),
					Note:      "note",
					CreatedAt: now,
				},
				UpdatedAt: now,
			}, nil)

		s := &service{authService: authService, serverGateway: gateway}

		info, err := s.GetBinaryByID(t.Context(), 1)
		assert.NoError(t, err)
		assert.NotNil(t, info)
		assert.Equal(t, "My file", info.Title)
		assert.Equal(t, []byte("abc"), info.Content)
		assert.Equal(t, now, info.CreatedAt)
		assert.Equal(t, now, info.UpdatedAt)
	})
}
