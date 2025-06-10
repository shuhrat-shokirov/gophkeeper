package session

import (
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"gophkeeper/internal/server/errorx"
	"gophkeeper/pkg/db"
)

func Test_repo_Get(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		pool := new(db.MockPool)
		defer pool.AssertExpectations(t)
		row := new(db.MockRow)
		defer row.AssertExpectations(t)

		pool.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(
			row,
			nil)

		row.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)

		r := &repo{
			dbConn: pool,
		}

		session, err := r.Get(t.Context(), "test-session-id")
		require.NoError(t, err)
		require.NotNil(t, session)
	})

	t.Run("GetWithError", func(t *testing.T) {
		pool := new(db.MockPool)
		defer pool.AssertExpectations(t)
		row := new(db.MockRow)
		defer row.AssertExpectations(t)

		pool.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(row)
		row.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(assert.AnError)
		r := &repo{
			dbConn: pool,
		}

		session, err := r.Get(t.Context(), "test-session-id")
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
		require.Nil(t, session)
	})

	t.Run("GetNotFound", func(t *testing.T) {
		pool := new(db.MockPool)
		defer pool.AssertExpectations(t)
		row := new(db.MockRow)
		defer row.AssertExpectations(t)

		pool.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(row)
		row.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(pgx.ErrNoRows)

		r := &repo{
			dbConn: pool,
		}

		session, err := r.Get(t.Context(), "test-session-id")
		require.ErrorIs(t, err, errorx.ErrNotFound)
		require.Nil(t, session)
	})
}
