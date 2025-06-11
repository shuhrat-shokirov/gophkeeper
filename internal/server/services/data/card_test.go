package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gophkeeper/internal/server/repositories/card"
)

func Test_service_SaveCard(t *testing.T) {
	t.Run("error from repository", func(t *testing.T) {
		repo := card.NewMockRepo()
		defer repo.AssertExpectations(t)

		s := New(Params{
			CardRepo: repo,
		})

		repo.On("Save", mock.Anything, mock.Anything).Return(nil, assert.AnError)
		err := s.SaveCard(t.Context(), &CardData{})
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		repo := card.NewMockRepo()
		defer repo.AssertExpectations(t)

		s := New(Params{
			CardRepo: repo,
		})

		repo.On("Save", mock.Anything, mock.Anything).Return(int64(1), nil)
		err := s.SaveCard(t.Context(), &CardData{})
		assert.NoError(t, err)
	})
}
