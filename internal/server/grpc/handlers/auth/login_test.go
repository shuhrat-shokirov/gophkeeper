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

func Test_handler_Login(t *testing.T) {
	t.Run("err not found", func(t *testing.T) {
		service := auth.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			AuthService: service,
		})

		service.On("Login", t.Context(), mock.Anything, mock.Anything).
			Return(nil, errorx.ErrNotFound)
		resp, err := h.Login(t.Context(), &pb.LoginRequest{})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, pb.LoginStatus_INVALID_CREDENTIALS, *resp.Status)
	})

	t.Run("other error", func(t *testing.T) {
		service := auth.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			AuthService: service,
		})

		service.On("Login", t.Context(), mock.Anything, mock.Anything).
			Return(nil, assert.AnError)
		resp, err := h.Login(t.Context(), &pb.LoginRequest{})
		require.Error(t, err, "expected error from service, got nil")
		require.NotNil(t, resp)
		require.Equal(t, pb.LoginStatus_LOGIN_ERROR, *resp.Status)
	})

	t.Run("success", func(t *testing.T) {
		service := auth.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			AuthService: service,
		})

		service.On("Login", t.Context(), mock.Anything, mock.Anything).
			Return("123123", nil)
		resp, err := h.Login(t.Context(), &pb.LoginRequest{
			Email:    proto.String("123@gmail.com"),
			Password: proto.String("testpass"),
		})
		require.NoError(t, err, "expected successful login, got error")
		require.NotNil(t, resp)
		require.Equal(t, pb.LoginStatus_LOGIN_SUCCESS, *resp.Status)
		require.Equal(t, "123123", *resp.OtpId)
	})
}
