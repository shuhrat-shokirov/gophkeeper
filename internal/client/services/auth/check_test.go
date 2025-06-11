//nolint:lll,gocritic,gosec
package auth

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"gophkeeper/internal/client/errorx"
	"gophkeeper/internal/client/gateways/server"
)

const (
	testPublicKey = "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUF6eUtPY2R1c0dJTy9VLzZnRTFsNAp0Z3Z4KzR5TUExQ1pQRkYwRnhKbzJaS1ZTY3I4SU41RVVDUDlUeEsxRTJLc2xnS01zQkhnYldJY3ZHMFBpTXZUClo0dTB3SWlPaTVSMDNlK3I5V1NqOG1xSCs3UjU1VndybUZVbFdMRWxuT1E4MnYveWNpV2hPZFJURWJ5cTZYQWcKaU5BckZyL3NFRTBacHFPdlVSeEFmeG5Qb1ZGd3M4NUplU0FYR1c2aG9HK3FoeEIvZ3diYkZOVitpbXViZUZ6dApyd1NJUGdmNjR2d2RoWnpDY1JZOVRUK1dyRm16Yk5uZmNwSzNvZEVnOCszdVJaWDdBb2R0U2E5OTZPTFFOcFNJClFnUDRiMnBXem5hc0NlRHU3ZlBWME5GNkJmSG1hcCs3RHZORTk1blcxUUh2MzZRRzNoVVRmZ0ZZQVJoc0NrOTAKcndJREFRQUIKLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg=="
	testToken     = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDk1NDA2NTcsImlhdCI6MTc0OTUzODg1NywiaXNzIjoiZ29waGtlZXBlciIsIm5iZiI6MTc0OTUzODg1NywidXNlcl9pZCI6MTF9.lOg4GN_4ybxxWlFY4e_ZR_bHeHQ1R9z1xmmXxCUvkBZ0vZRQPXsm0yuyUteKEYJP3B7vheuP8JFXtvADP4fukXxzVTxdubpDZTxN-zlntbRAdOZmHb29BtPgqhIDnpSwUtucPxgqMlhrzR6ZL1zJF4xl9BeYEAI3-8EFMfzVItd_pDD3_pvI_Vd8-Gg7Slrm9rjlHec60OugTVuWKJ-_MNVQsp-_cIvxOj8Mipz8oB2_UsWTJH-vK5DngaMvsxnHcq6gA7aDzwfSqjb2hur2oSqXX_xCblex2ntxxzNjbJpzJGQkeprIvVogksoGkqRcjhBbdcNKtqmG_hqSf8GPkg"
)

func Test_service_CheckAuth(t *testing.T) {
	t.Run("access token is empty", func(t *testing.T) {
		s := &service{
			accessToken: "",
		}

		err := s.CheckAuth(t.Context())
		assert.Error(t, err)
		assert.Equal(t, err, errorx.ErrTokenNotFound)
	})

	t.Run("err from jwt", func(t *testing.T) {
		s := &service{
			accessToken:  "123",
			refreshToken: "456",
		}

		err := s.CheckAuth(t.Context())
		assert.Error(t, err)
	})

	t.Run("err from gateway", func(t *testing.T) {
		bytes, err := base64.StdEncoding.DecodeString(testPublicKey)
		require.NoError(t, err)
		require.NotNil(t, bytes)

		serverGateway := server.NewMockGateway()
		defer serverGateway.AssertExpectations(t)

		s := &service{
			serverGateway: serverGateway,
			accessToken:   testToken,
			refreshToken:  "123",
			publicKey:     bytes,
		}

		serverGateway.On("RefreshToken", mock.Anything, mock.Anything).
			Return("", assert.AnError)

		err = s.CheckAuth(t.Context())
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		bytes, err := base64.StdEncoding.DecodeString(testPublicKey)
		require.NoError(t, err)
		require.NotNil(t, bytes)

		serverGateway := server.NewMockGateway()
		defer serverGateway.AssertExpectations(t)

		s := &service{
			serverGateway: serverGateway,
			accessToken:   testToken,
			refreshToken:  "123",
			publicKey:     bytes,
		}

		serverGateway.On("RefreshToken", mock.Anything, mock.Anything).
			Return(&server.Token{AccessToken: "new-access-token"}, nil)

		err = s.CheckAuth(t.Context())
		assert.NoError(t, err)
		assert.Equal(t, "new-access-token", s.accessToken)
	})
}

func Test_service_GetUserID(t *testing.T) {
	t.Run("user ID is set", func(t *testing.T) {
		s := &service{
			userID: 42,
		}

		id, err := s.GetUserID(t.Context())
		assert.NoError(t, err)
		assert.Equal(t, int64(42), id)
	})

	t.Run("err from jwt", func(t *testing.T) {
		s := &service{}

		id, err := s.GetUserID(t.Context())
		assert.Error(t, err)
		assert.Equal(t, int64(0), id)
	})

	t.Run("token expired", func(t *testing.T) {
		bytes, err := base64.StdEncoding.DecodeString(testPublicKey)
		require.NoError(t, err)
		require.NotNil(t, bytes)

		s := &service{
			accessToken: testToken,
			publicKey:   bytes,
		}

		id, err := s.GetUserID(t.Context())
		assert.Error(t, err)
		assert.Equal(t, int64(0), id)
	})
}
