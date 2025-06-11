package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"

	"gophkeeper/internal/client/errorx"
	pb "gophkeeper/proto"
)

func Test_gateway_Register(t *testing.T) {
	t.Run("err from server", func(t *testing.T) {
		mockClient := new(MockAuthServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("Register", mock.Anything, mock.Anything).Return(nil, assert.AnError)

		g := &gateway{
			authServiceClient: mockClient,
		}

		otpID, err := g.Register(t.Context(), "", "")
		assert.Error(t, err, "expected error from server, got nil")
		assert.Empty(t, otpID, "expected empty otpID on error")
	})

	t.Run("user already exists", func(t *testing.T) {
		mockClient := new(MockAuthServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("Register", mock.Anything, mock.Anything).Return(&pb.RegisterResponse{
			Status: pb.RegisterStatus_USER_ALREADY_EXISTS.Enum(),
		}, nil)

		g := &gateway{
			authServiceClient: mockClient,
		}

		otpID, err := g.Register(t.Context(), "", "")
		assert.Error(t, err, "expected error for user already exists, got nil")
		assert.Empty(t, otpID, "expected empty otpID for user already exists")
		assert.Equal(t, err, errorx.ErrUserAlreadyExists, "expected ErrAlreadyExists error")
	})

	t.Run("err from service", func(t *testing.T) {
		mockClient := new(MockAuthServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("Register", mock.Anything, mock.Anything).Return(&pb.RegisterResponse{
			Status: pb.RegisterStatus_REGISTER_ERROR.Enum(),
		}, nil)

		g := &gateway{
			authServiceClient: mockClient,
		}

		otpID, err := g.Register(t.Context(), "", "")
		assert.Error(t, err, "expected error from service, got nil")
		assert.Empty(t, otpID, "expected empty otpID on service error")
	})

	t.Run("success", func(t *testing.T) {
		mockClient := new(MockAuthServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("Register", mock.Anything, mock.Anything).Return(&pb.RegisterResponse{
			Status: pb.RegisterStatus_OTP_SENT.Enum(),
			OtpId:  proto.String("otp_id"),
		}, nil)

		g := &gateway{
			authServiceClient: mockClient,
		}

		otpID, err := g.Register(t.Context(), "", "")
		assert.NoError(t, err, "expected successful registration, got error")
		assert.Equal(t, "otp_id", otpID, "expected otpID to match")
	})
}
