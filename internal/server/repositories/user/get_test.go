//nolint:errcheck,gocritic
package user

import (
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gophkeeper/internal/server/errorx"
	"gophkeeper/pkg/db"
)

func Test_repo_GetUserByEmail(t *testing.T) {
	t.Run("Error on get user by email", func(t *testing.T) {
		pool := new(db.MockPool)
		defer pool.AssertExpectations(t)
		row := new(db.MockRow)
		defer row.AssertExpectations(t)

		pool.On("QueryRow", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(row)
		row.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(assert.AnError)

		r := &repo{
			dbConn: pool,
		}

		user, err := r.GetUserByEmail(t.Context(), "123")
		assert.Nil(t, user)
		assert.Error(t, err)
	})

	t.Run("Err not found on get user by email", func(t *testing.T) {
		pool := new(db.MockPool)
		defer pool.AssertExpectations(t)
		row := new(db.MockRow)
		defer row.AssertExpectations(t)

		pool.On("QueryRow", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(row)
		row.On("Scan", mock.Anything, mock.Anything, mock.Anything,
			mock.Anything, mock.Anything).Return(pgx.ErrNoRows)

		r := &repo{
			dbConn: pool,
		}

		user, err := r.GetUserByEmail(t.Context(), "123")
		assert.Nil(t, user)
		assert.ErrorIs(t, err, errorx.ErrNotFound)
	})

	t.Run("Success on get user by email", func(t *testing.T) {
		pool := new(db.MockPool)
		defer pool.AssertExpectations(t)
		row := new(db.MockRow)
		defer row.AssertExpectations(t)

		pool.On("QueryRow", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(row)
		row.On("Scan", mock.Anything, mock.Anything, mock.Anything,
			mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			id := args.Get(0).(*int)
			email := args.Get(1).(*string)
			password := args.Get(2).(*string)
			*id = 1
			*email = "123@gmail.com"
			*password = "password123"
		})

		r := &repo{
			dbConn: pool,
		}

		user, err := r.GetUserByEmail(t.Context(), "")
		assert.NoError(t, err)
		assert.NotNil(t, user)

		assert.Equal(t, 1, user.ID)
		assert.Equal(t, "123@gmail.com", user.Email)
		assert.Equal(t, "password123", user.Password)
	})
}

func Test_repo_GetUserByID(t *testing.T) {
	t.Run("Error on get user by id", func(t *testing.T) {
		pool := new(db.MockPool)
		defer pool.AssertExpectations(t)
		row := new(db.MockRow)
		defer row.AssertExpectations(t)

		pool.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(row)
		row.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(assert.AnError)

		r := &repo{
			dbConn: pool,
		}

		user, err := r.GetUserByID(t.Context(), 1)
		assert.Nil(t, user)
		assert.Error(t, err)
	})

	t.Run("Err not found on get user by id", func(t *testing.T) {
		pool := new(db.MockPool)
		defer pool.AssertExpectations(t)
		row := new(db.MockRow)
		defer row.AssertExpectations(t)

		pool.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(row)
		row.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(pgx.ErrNoRows)

		r := &repo{
			dbConn: pool,
		}

		user, err := r.GetUserByID(t.Context(), 1)
		assert.Nil(t, user)
		assert.ErrorIs(t, err, errorx.ErrNotFound)
	})

	t.Run("Success on get user by id", func(t *testing.T) {
		pool := new(db.MockPool)
		defer pool.AssertExpectations(t)
		row := new(db.MockRow)
		defer row.AssertExpectations(t)

		pool.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(row)
		row.On("Scan", mock.Anything, mock.Anything,
			mock.Anything, mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			id := args.Get(0).(*int)
			email := args.Get(1).(*string)
			password := args.Get(2).(*string)
			*id = 1
			*email = "123@gmail.com"
			*password = "password123"
		})

		r := &repo{
			dbConn: pool,
		}

		user, err := r.GetUserByID(t.Context(), 1)
		assert.NoError(t, err)
		assert.NotNil(t, user)

		assert.Equal(t, 1, user.ID)
		assert.Equal(t, "123@gmail.com", user.Email)
		assert.Equal(t, "password123", user.Password)
	})
}
