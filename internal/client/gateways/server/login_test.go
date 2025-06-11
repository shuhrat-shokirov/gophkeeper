package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"

	"gophkeeper/internal/client/errorx"
	pb "gophkeeper/proto"
)

func Test_gateway_Login(t *testing.T) {
	t.Run("err from server", func(t *testing.T) {
		mockClient := new(MockAuthServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("Login", mock.Anything, mock.Anything).Return(nil, assert.AnError)

		g := &gateway{
			authServiceClient: mockClient,
		}

		token, err := g.Login(t.Context(), "username", "password")
		assert.Error(t, err, "expected error from server, got nil")
		assert.Empty(t, token, "expected nil response on error")
	})

	t.Run("invalid credentials", func(t *testing.T) {
		mockClient := new(MockAuthServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("Login", mock.Anything, mock.Anything).Return(&pb.LoginResponse{
			Status: pb.LoginStatus_INVALID_CREDENTIALS.Enum(),
		}, nil)

		g := &gateway{
			authServiceClient: mockClient,
		}

		token, err := g.Login(t.Context(), "invalid_user", "wrong_password")
		assert.Error(t, err, "expected error for invalid credentials, got nil")
		assert.Empty(t, token, "expected empty token for invalid credentials")
		assert.Equal(t, err, errorx.ErrInvalidCredentials)
	})

	t.Run("err from server with status", func(t *testing.T) {
		mockClient := new(MockAuthServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("Login", mock.Anything, mock.Anything).Return(&pb.LoginResponse{
			Status: pb.LoginStatus_LOGIN_ERROR.Enum(),
		}, nil)

		g := &gateway{
			authServiceClient: mockClient,
		}

		token, err := g.Login(t.Context(), "username", "password")
		assert.Error(t, err, "expected error for login error status, got nil")
		assert.Empty(t, token, "expected empty token for login error status")
	})

	t.Run("success", func(t *testing.T) {
		mockClient := new(MockAuthServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("Login", mock.Anything, mock.Anything).Return(&pb.LoginResponse{
			Status: pb.LoginStatus_LOGIN_SUCCESS.Enum(),
			OtpId:  proto.String("otp_id"),
		}, nil)

		g := &gateway{
			authServiceClient: mockClient,
		}

		otpId, err := g.Login(t.Context(), "valid_user", "correct_password")
		assert.NoError(t, err, "expected successful login, got error")
		assert.Equal(t, "otp_id", otpId, "expected valid token on successful login")
	})
}
