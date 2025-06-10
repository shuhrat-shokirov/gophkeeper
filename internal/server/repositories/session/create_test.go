package session

import (
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"gophkeeper/pkg/db"
)

func Test_repo_Create(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		pool := new(db.MockPool)
		defer pool.AssertExpectations(t)

		pool.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(
			pgconn.CommandTag{},
			nil)
		r := &repo{
			dbConn: pool,
		}

		err := r.Create(t.Context(), &Session{})
		require.NoError(t, err)
	})
	t.Run("Error on create session", func(t *testing.T) {
		pool := new(db.MockPool)
		defer pool.AssertExpectations(t)

		pool.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(
			pgconn.CommandTag{},
			assert.AnError)
		r := &repo{
			dbConn: pool,
		}

		err := r.Create(t.Context(), &Session{})
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
}
