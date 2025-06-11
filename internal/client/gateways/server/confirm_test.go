package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"

	"gophkeeper/internal/client/errorx"
	pb "gophkeeper/proto"
)

func Test_gateway_ConfirmOtp(t *testing.T) {
	t.Run("err from server", func(t *testing.T) {
		mockClient := new(MockAuthServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("ConfirmOTP", mock.Anything, mock.Anything).Return(&pb.ConfirmOTPResponse{
			Status: pb.ConfirmOTPStatus_CONFIRM_ERROR.Enum(),
		}, assert.AnError)

		g := &gateway{
			authServiceClient: mockClient,
		}

		token, err := g.ConfirmOtp(t.Context(), "otp_id", "123")
		assert.Error(t, err, "expected error from server, got nil")
		assert.Nil(t, token, "expected nil response on error")
	})

	t.Run("invalid code", func(t *testing.T) {
		mockClient := new(MockAuthServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("ConfirmOTP", mock.Anything, mock.Anything).Return(&pb.ConfirmOTPResponse{
			Status: pb.ConfirmOTPStatus_INVALID_CODE.Enum(),
		}, nil)

		g := &gateway{
			authServiceClient: mockClient,
		}

		token, err := g.ConfirmOtp(t.Context(), "otp_id", "invalid_code")
		assert.Error(t, err, "expected error for invalid code, got nil")
		assert.Nil(t, token, "expected nil response for invalid code")
		assert.Equal(t, err, errorx.ErrOtpInvalid)
	})

	t.Run("code expired", func(t *testing.T) {
		mockClient := new(MockAuthServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("ConfirmOTP", mock.Anything, mock.Anything).Return(&pb.ConfirmOTPResponse{
			Status: pb.ConfirmOTPStatus_CODE_EXPIRED.Enum(),
		}, nil)

		g := &gateway{
			authServiceClient: mockClient,
		}

		token, err := g.ConfirmOtp(t.Context(), "otp_id", "expired_code")
		assert.Error(t, err, "expected error for expired code, got nil")
		assert.Nil(t, token, "expected nil response for expired code")
		assert.Equal(t, err, errorx.ErrOtpExpired)
	})

	t.Run("user not found", func(t *testing.T) {
		mockClient := new(MockAuthServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("ConfirmOTP", mock.Anything, mock.Anything).Return(&pb.ConfirmOTPResponse{
			Status: pb.ConfirmOTPStatus_USER_NOT_FOUND_CODE.Enum(),
		}, nil)

		g := &gateway{
			authServiceClient: mockClient,
		}

		token, err := g.ConfirmOtp(t.Context(), "otp_id", "user_not_found_code")
		assert.Error(t, err, "expected error for user not found, got nil")
		assert.Nil(t, token, "expected nil response for user not found")
		assert.Equal(t, err, errorx.ErrUserNotFound)
	})

	t.Run("error code", func(t *testing.T) {
		mockClient := new(MockAuthServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("ConfirmOTP", mock.Anything, mock.Anything).Return(&pb.ConfirmOTPResponse{
			Status:  pb.ConfirmOTPStatus_CONFIRM_ERROR.Enum(),
			Message: proto.String("unexpected error"),
		}, nil)

		g := &gateway{
			authServiceClient: mockClient,
		}

		token, err := g.ConfirmOtp(t.Context(), "otp_id", "error_code")
		assert.Error(t, err, "expected error for unexpected status, got nil")
		assert.Nil(t, token, "expected nil response for unexpected status")
		assert.Equal(t, err.Error(), "unexpected response status: unexpected error")
	})

	t.Run("success", func(t *testing.T) {
		mockClient := new(MockAuthServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("ConfirmOTP", mock.Anything, mock.Anything).Return(&pb.ConfirmOTPResponse{
			Status:       pb.ConfirmOTPStatus_CONFIRM_SUCCESS.Enum(),
			Token:        proto.String("access_token"),
			RefreshToken: proto.String("refresh_token"),
		}, nil)

		g := &gateway{
			authServiceClient: mockClient,
		}

		token, err := g.ConfirmOtp(t.Context(), "otp_id", "valid_code")
		assert.NoError(t, err, "expected successful confirmation, got error")
		assert.NotNil(t, token, "expected non-nil token on success")
		assert.Equal(t, "access_token", token.AccessToken, "expected correct access token")
		assert.Equal(t, "refresh_token", token.RefreshToken, "expected correct refresh token")
	})
}
