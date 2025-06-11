package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"gophkeeper/internal/server/errorx"
	"gophkeeper/internal/server/services/auth"
	pb "gophkeeper/proto"
)

func Test_handler_ConfirmOTP(t *testing.T) {
	t.Run("err otp expired", func(t *testing.T) {
		service := auth.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			AuthService: service,
		})
		require.NotNil(t, h)

		service.On("ConfirmOTP", t.Context(), mock.Anything, mock.Anything).
			Return(nil, errorx.ErrOTPExpired)
		resp, err := h.ConfirmOTP(t.Context(), &pb.ConfirmOTPRequest{})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, pb.ConfirmOTPStatus_CODE_EXPIRED, *resp.Status)
	})

	t.Run("invalid otp", func(t *testing.T) {
		service := auth.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			AuthService: service,
		})
		require.NotNil(t, h)

		service.On("ConfirmOTP", t.Context(), mock.Anything, mock.Anything).
			Return(nil, errorx.ErrInvalidOTP)
		resp, err := h.ConfirmOTP(t.Context(), &pb.ConfirmOTPRequest{})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, pb.ConfirmOTPStatus_INVALID_CODE, *resp.Status)
	})

	t.Run("user not found", func(t *testing.T) {
		service := auth.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			AuthService: service,
		})
		require.NotNil(t, h)

		service.On("ConfirmOTP", t.Context(), mock.Anything, mock.Anything).
			Return(nil, errorx.ErrNotFound)
		resp, err := h.ConfirmOTP(t.Context(), &pb.ConfirmOTPRequest{})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, pb.ConfirmOTPStatus_USER_NOT_FOUND_CODE, *resp.Status)
	})

	t.Run("other error", func(t *testing.T) {
		service := auth.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			AuthService: service,
		})
		require.NotNil(t, h)

		service.On("ConfirmOTP", t.Context(), mock.Anything, mock.Anything).
			Return(nil, assert.AnError)
		resp, err := h.ConfirmOTP(t.Context(), &pb.ConfirmOTPRequest{})
		require.Error(t, err)
		require.NotNil(t, resp)
		require.Equal(t, pb.ConfirmOTPStatus_CONFIRM_ERROR, *resp.Status)
	})

	t.Run("success", func(t *testing.T) {
		service := auth.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			AuthService: service,
		})
		require.NotNil(t, h)

		service.On("ConfirmOTP", t.Context(), mock.Anything, mock.Anything).
			Return(&auth.ConfirmResponse{
				UserId:       1,
				Token:        "test-token",
				RefreshToken: "test-refresh-token",
			}, nil)
		resp, err := h.ConfirmOTP(t.Context(), &pb.ConfirmOTPRequest{
			OtpId: proto.String("test-otp-id"),
			Code:  proto.String("123456"),
		})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, pb.ConfirmOTPStatus_CONFIRM_SUCCESS, *resp.Status)
		assert.Equal(t, int64(1), *resp.UserId)
		assert.Equal(t, "test-token", *resp.Token)
		assert.Equal(t, "test-refresh-token", *resp.RefreshToken)
	})
}
