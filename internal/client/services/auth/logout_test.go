package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"gophkeeper/internal/client/gateways/server"
)

func Test_service_Logout(t *testing.T) {
	t.Run("token empty", func(t *testing.T) {
		s := &service{
			refreshToken: "",
		}

		s.Logout(t.Context())
	})

	t.Run("successful logout", func(t *testing.T) {
		serv := server.NewMockGateway()
		defer serv.AssertExpectations(t)

		s := &service{
			serverGateway: serv,
			refreshToken:  "123",
			accessToken:   "123",
		}
		require.NotNil(t, s)

		serv.On("Logout", mock.Anything, mock.Anything).Return(nil)
		s.Logout(t.Context())

		assert.Empty(t, s.refreshToken)
		assert.Empty(t, s.accessToken)
	})
}
