package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"

	pb "gophkeeper/proto"
)

func Test_gateway_RefreshToken(t *testing.T) {
	t.Run("err from server", func(t *testing.T) {
		mockClient := new(MockAuthServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("RefreshToken", mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		g := &gateway{
			authServiceClient: mockClient,
		}

		token, err := g.RefreshToken(t.Context(), "refresh_token")
		assert.Error(t, err, "expected error from server, got nil")
		assert.Empty(t, token, "expected empty token on error")
	})

	t.Run("invalid refresh token", func(t *testing.T) {
		mockClient := new(MockAuthServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("RefreshToken", mock.Anything, mock.Anything).
			Return(&pb.RefreshTokenResponse{
				Status: pb.RefreshTokenStatus_INVALID_REFRESH_TOKEN.Enum(),
			}, nil)

		g := &gateway{
			authServiceClient: mockClient,
		}

		token, err := g.RefreshToken(t.Context(), "invalid_refresh_token")
		assert.Error(t, err, "expected error for invalid refresh token, got nil")
		assert.Empty(t, token, "expected empty token for invalid refresh token")
	})

	t.Run("non success status", func(t *testing.T) {
		mockClient := new(MockAuthServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("RefreshToken", mock.Anything, mock.Anything).
			Return(&pb.RefreshTokenResponse{
				Status: pb.RefreshTokenStatus_REFRESH_ERROR.Enum(),
			}, nil)

		g := &gateway{
			authServiceClient: mockClient,
		}

		token, err := g.RefreshToken(t.Context(), "some_refresh_token")
		assert.Error(t, err, "expected error for non-success status, got nil")
		assert.Empty(t, token, "expected empty token for non-success status")
	})

	t.Run("success", func(t *testing.T) {
		mockClient := new(MockAuthServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("RefreshToken", mock.Anything, mock.Anything).
			Return(&pb.RefreshTokenResponse{
				Status: pb.RefreshTokenStatus_REFRESH_SUCCESS.Enum(),
				Token:  proto.String("new_access_token"),
			}, nil)

		g := &gateway{
			authServiceClient: mockClient,
		}

		token, err := g.RefreshToken(t.Context(), "valid_refresh_token")
		assert.NoError(t, err, "expected successful refresh, got error")
		assert.Equal(t, "new_access_token", token.AccessToken, "expected new access token")
	})
}
