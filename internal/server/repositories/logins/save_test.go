package logins

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"gophkeeper/pkg/db"
)

func Test_repo_Save(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		pool := new(db.MockPool)
		defer pool.AssertExpectations(t)
		row := new(db.MockRow)
		defer row.AssertExpectations(t)

		pool.On("QueryRow", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(row)
		row.On("Scan", mock.Anything).Return(nil)

		r := &repo{
			dbConn: pool,
		}

		_, err := r.Save(t.Context(), &LoginData{})
		require.NoError(t, err)
	})
	t.Run("fail", func(t *testing.T) {
		pool := new(db.MockPool)
		defer pool.AssertExpectations(t)
		row := new(db.MockRow)
		defer row.AssertExpectations(t)

		pool.On("QueryRow", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(row)
		row.On("Scan", mock.Anything).Return(assert.AnError)

		r := &repo{
			dbConn: pool,
		}

		_, err := r.Save(t.Context(), &LoginData{})
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
}
