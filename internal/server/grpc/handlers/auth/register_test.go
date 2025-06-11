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

func Test_handler_Register(t *testing.T) {
	t.Run("user already exists", func(t *testing.T) {
		service := auth.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			AuthService: service,
		})

		service.On("Registration", t.Context(), mock.Anything, mock.Anything).
			Return("", errorx.ErrAlreadyExists)

		resp, err := h.Register(t.Context(), &pb.RegisterRequest{})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, pb.RegisterStatus_USER_ALREADY_EXISTS, *resp.Status)
	})

	t.Run("error from service", func(t *testing.T) {
		service := auth.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			AuthService: service,
		})

		service.On("Registration", t.Context(), mock.Anything, mock.Anything).
			Return("", assert.AnError)

		resp, err := h.Register(t.Context(), &pb.RegisterRequest{})
		require.Error(t, err, "expected error from service, got nil")
		require.NotNil(t, resp)
		require.Equal(t, pb.RegisterStatus_REGISTER_ERROR, *resp.Status)
	})

	t.Run("success", func(t *testing.T) {
		service := auth.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			AuthService: service,
		})

		service.On("Registration", t.Context(), mock.Anything, mock.Anything).
			Return("otp_id", nil)

		resp, err := h.Register(t.Context(), &pb.RegisterRequest{
			Email:    proto.String(""),
			Password: proto.String(""),
		})
		require.NoError(t, err, "expected successful registration, got error")
		require.NotNil(t, resp)
		require.Equal(t, pb.RegisterStatus_OTP_SENT, *resp.Status)
		require.Equal(t, "otp_id", *resp.OtpId)
	})

}
