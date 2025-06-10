//nolint:errcheck,gocritic
package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gophkeeper/pkg/db"
)

func Test_repo_CreateUser(t *testing.T) {
	t.Run("Error on create user", func(t *testing.T) {
		pool := new(db.MockPool)
		defer pool.AssertExpectations(t)
		row := new(db.MockRow)
		defer row.AssertExpectations(t)

		pool.On("QueryRow", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(row)
		row.On("Scan", mock.Anything).Return(assert.AnError)

		r := &repo{
			dbConn: pool,
		}

		createUser, err := r.CreateUser(t.Context(), &User{})
		assert.Zero(t, createUser)
		assert.Error(t, err)
	})

	t.Run("Success on create user", func(t *testing.T) {
		pool := new(db.MockPool)
		defer pool.AssertExpectations(t)
		row := new(db.MockRow)
		defer row.AssertExpectations(t)

		pool.On("QueryRow", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(row)
		row.On("Scan", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			id := args.Get(0).(*int)
			*id = 1
		})

		r := &repo{
			dbConn: pool,
		}

		createUser, err := r.CreateUser(t.Context(), &User{})
		assert.Equal(t, 1, createUser)
		assert.NoError(t, err)
	})
}
