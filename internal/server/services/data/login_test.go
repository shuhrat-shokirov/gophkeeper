package data

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"gophkeeper/internal/server/repositories/logins"
	"gophkeeper/pkg/aes"
	"gophkeeper/pkg/config"
)

func Test_service_SaveLogin(t *testing.T) {
	t.Run("error from repository", func(t *testing.T) {
		repo := logins.NewMockRepo()
		defer repo.AssertExpectations(t)

		s := New(Params{
			LoginsRepo: repo,
		})

		repo.On("Save", t.Context(), mock.Anything).Return(nil, assert.AnError)
		err := s.SaveLogin(t.Context(), &LoginData{})
		assert.Error(t, err, "expected error from repository, got nil")
	})

	t.Run("success", func(t *testing.T) {
		repo := logins.NewMockRepo()
		defer repo.AssertExpectations(t)

		s := New(Params{
			LoginsRepo: repo,
		})

		repo.On("Save", t.Context(), mock.Anything).Return(int64(1), nil)
		err := s.SaveLogin(t.Context(), &LoginData{Login: "testuser", Password: "testpass"})
		assert.NoError(t, err, "expected successful save, got error")
	})
}

func Test_service_GetLoginList(t *testing.T) {
	t.Run("error from repository", func(t *testing.T) {
		repo := logins.NewMockRepo()
		defer repo.AssertExpectations(t)

		s := New(Params{
			LoginsRepo: repo,
		})

		repo.On("List", t.Context(), mock.Anything, mock.Anything).
			Return(nil, assert.AnError)
		list, err := s.GetLoginList(t.Context(), 1, 10, 0)
		assert.Error(t, err, "expected error from repository, got nil")
		assert.Nil(t, list, "expected nil login list on error")
	})

	t.Run("success", func(t *testing.T) {
		repo := logins.NewMockRepo()
		defer repo.AssertExpectations(t)

		s := New(Params{
			LoginsRepo: repo,
		})

		expectedLogins := []logins.List{
			{ID: 1, Title: "testuser"},
			{ID: 2, Title: "anotheruser"},
		}

		repo.On("List", t.Context(), mock.Anything, mock.Anything).
			Return(expectedLogins, nil)

		list, err := s.GetLoginList(t.Context(), 1, 10, 0)
		assert.NoError(t, err, "expected successful login list retrieval, got error")
		assert.Len(t, list, 2, "expected 2 logins in the list")
		assert.Equal(t, expectedLogins[0].ID, list[0].ID, "expected first login ID to match")
		assert.Equal(t, expectedLogins[0].Title, list[0].Title, "expected first login title to match")
	})
}

func Test_service_GetLoginByID(t *testing.T) {
	t.Run("error from repository", func(t *testing.T) {
		repo := logins.NewMockRepo()
		defer repo.AssertExpectations(t)

		s := New(Params{
			LoginsRepo: repo,
		})

		repo.On("GetByID", t.Context(), int64(1)).
			Return(nil, assert.AnError)

		login, err := s.GetLoginByID(t.Context(), 1)
		assert.Error(t, err, "expected error from repository, got nil")
		assert.Nil(t, login, "expected nil login on error")
	})

	t.Run("success", func(t *testing.T) {

		err := os.Setenv("AES_SECRET_KEY", "1234567890abcdef")
		require.NoError(t, err)

		newConfig, err := config.NewConfig()
		require.NoError(t, err)
		require.NotNil(t, newConfig)

		aes.New(aes.Params{Config: newConfig})

		repo := logins.NewMockRepo()
		defer repo.AssertExpectations(t)

		s := New(Params{
			LoginsRepo: repo,
		})

		expectedLogin := &logins.Info{
			ID: 1,
			LoginData: logins.LoginData{
				Login:    "DkknXKFwRWcs260j83q5AMUhfw==",
				Password: "DL5a6ImiFj2M3f/6kFZCGPvXKw==",
				Title:    "Test User",
				Note:     "dafuxvX4FaLyh9CT+Tvxwg==",
			},
		}

		repo.On("GetByID", t.Context(), int64(1)).
			Return(expectedLogin, nil)

		login, err := s.GetLoginByID(t.Context(), 1)
		require.NoError(t, err, "expected successful login retrieval, got error")
		require.NotNil(t, login, "expected non-nil login")
		require.NotNil(t, login.LoginData)

		assert.Equal(t, expectedLogin.ID, login.ID, "expected login ID to match")
		assert.Equal(t, login.Login, "qwe", "expected login to match")
		assert.Equal(t, login.Password, "qwe", "expected password to match")
		assert.Equal(t, login.Title, "Test User", "expected title to match")
		assert.Equal(t, login.Note, "", "expected note to match")
	})
}
