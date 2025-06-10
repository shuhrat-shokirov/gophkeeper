//nolint:errcheck,gocritic
package logins

import (
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gophkeeper/internal/server/errorx"
	"gophkeeper/pkg/db"
)

func Test_repo_GetByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		pool := new(db.MockPool)
		defer pool.AssertExpectations(t)
		row := new(db.MockRow)
		defer row.AssertExpectations(t)

		pool.On("QueryRow", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(row)
		row.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything,
			mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).
			Run(func(args mock.Arguments) {
				id := args.Get(0).(*int64)
				*id = 1
				userID := args.Get(1).(*int64)
				*userID = 1
				login := args.Get(2).(*string)
				*login = "test_login"
				password := args.Get(3).(*string)
				*password = "test_password"
			})

		r := &repo{
			dbConn: pool,
		}

		loginData, err := r.GetByID(t.Context(), 1)
		assert.NoError(t, err)
		assert.NotNil(t, loginData)

		assert.Equal(t, int64(1), loginData.ID)
		assert.Equal(t, "test_login", loginData.Login)
		assert.Equal(t, "test_password", loginData.Password)
	})

	t.Run("failure", func(t *testing.T) {
		pool := new(db.MockPool)
		defer pool.AssertExpectations(t)
		row := new(db.MockRow)
		defer row.AssertExpectations(t)

		pool.On("QueryRow", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(row)
		row.On("Scan", mock.Anything, mock.Anything, mock.Anything,
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(assert.AnError)

		r := &repo{
			dbConn: pool,
		}

		loginData, err := r.GetByID(t.Context(), 1)
		assert.Error(t, err)
		assert.Nil(t, loginData)
		assert.ErrorIs(t, err, assert.AnError)
	})

	t.Run("NotFound", func(t *testing.T) {
		pool := new(db.MockPool)
		defer pool.AssertExpectations(t)
		row := new(db.MockRow)
		defer row.AssertExpectations(t)

		pool.On("QueryRow", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(row)
		row.On("Scan", mock.Anything, mock.Anything, mock.Anything,
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(pgx.ErrNoRows)

		r := &repo{
			dbConn: pool,
		}

		loginData, err := r.GetByID(t.Context(), 1)
		assert.ErrorIs(t, err, errorx.ErrNotFound)
		assert.Nil(t, loginData)
	})
}

func Test_repo_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		pool := new(db.MockPool)
		defer pool.AssertExpectations(t)
		rows := new(db.MockRows)
		defer rows.AssertExpectations(t)

		pool.On("Query", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(rows, nil)
		rows.On("Next").Return(true).Once()
		rows.On("Next").Return(false)

		rows.On("Scan", mock.Anything, mock.Anything, mock.Anything,
			mock.Anything, mock.Anything, mock.Anything,
			mock.Anything, mock.Anything).Return(nil).Once()
		rows.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

		rows.On("Close").Return(nil)

		r := &repo{
			dbConn: pool,
		}

		loginDataList, err := r.List(t.Context(), 1, Pagination{})
		assert.NoError(t, err)
		assert.NotNil(t, loginDataList)
	})

	t.Run("failure", func(t *testing.T) {
		pool := new(db.MockPool)
		defer pool.AssertExpectations(t)
		rows := new(db.MockRows)
		defer rows.AssertExpectations(t)

		pool.On("Query", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(rows, assert.AnError)

		r := &repo{
			dbConn: pool,
		}

		loginDataList, err := r.List(t.Context(), 1, Pagination{})
		assert.Error(t, err)
		assert.Nil(t, loginDataList)
		assert.ErrorIs(t, err, assert.AnError)
	})

	t.Run("empty", func(t *testing.T) {
		pool := new(db.MockPool)
		defer pool.AssertExpectations(t)
		rows := new(db.MockRows)
		defer rows.AssertExpectations(t)

		pool.On("Query", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(rows, nil)
		rows.On("Next").Return(false)

		rows.On("Close").Return(nil)

		r := &repo{
			dbConn: pool,
		}

		loginDataList, err := r.List(t.Context(), 1, Pagination{})
		assert.ErrorIs(t, err, errorx.ErrNotFound)
		assert.Empty(t, loginDataList)
	})

	t.Run("scan error", func(t *testing.T) {
		pool := new(db.MockPool)
		defer pool.AssertExpectations(t)
		rows := new(db.MockRows)
		defer rows.AssertExpectations(t)

		pool.On("Query", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(rows, nil)
		rows.On("Next").Return(true).Once()
		rows.On("Next").Return(false)

		rows.On("Scan", mock.Anything, mock.Anything, mock.Anything,
			mock.Anything, mock.Anything, mock.Anything,
			mock.Anything, mock.Anything).Return(assert.AnError).Once()

		rows.On("Close").Return(nil)

		r := &repo{
			dbConn: pool,
		}

		loginDataList, err := r.List(t.Context(), 1, Pagination{})
		assert.Error(t, err)
		assert.Nil(t, loginDataList)
		assert.ErrorIs(t, err, assert.AnError)
	})
}
