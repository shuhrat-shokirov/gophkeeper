package data

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gophkeeper/internal/server/repositories/binary"
)

func Test_service_SaveBinary(t *testing.T) {
	t.Run("error from repository", func(t *testing.T) {
		repo := binary.NewMockRepo()
		defer repo.AssertExpectations(t)

		s := New(Params{
			BinaryRepo: repo,
		})

		repo.On("Save", mock.Anything, mock.Anything).Return(nil, assert.AnError)
		err := s.SaveBinary(context.Background(), &BinaryData{})
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		repo := binary.NewMockRepo()
		defer repo.AssertExpectations(t)

		s := New(Params{
			BinaryRepo: repo,
		})

		repo.On("Save", mock.Anything, mock.Anything).Return(int64(1), nil)
		err := s.SaveBinary(context.Background(), &BinaryData{Content: []byte("test content")})
		assert.NoError(t, err)
	})
}
