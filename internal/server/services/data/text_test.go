package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gophkeeper/internal/server/repositories/texts"
)

func Test_service_SaveText(t *testing.T) {
	t.Run("error from repository", func(t *testing.T) {
		repo := texts.NewMockRepo()
		defer repo.AssertExpectations(t)

		s := New(Params{
			TextRepo: repo,
		})

		repo.On("Save", mock.Anything, mock.Anything).Return(nil, assert.AnError)
		err := s.SaveText(t.Context(), &TextData{})
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		repo := texts.NewMockRepo()
		defer repo.AssertExpectations(t)

		s := New(Params{
			TextRepo: repo,
		})

		repo.On("Save", mock.Anything, mock.Anything).Return(int64(1), nil)
		err := s.SaveText(t.Context(), &TextData{Content: "test content"})
		assert.NoError(t, err)
	})
}
