package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"

	"gophkeeper/internal/server/errorx"
	"gophkeeper/internal/server/services/auth"
	pb "gophkeeper/proto"
)

func Test_handler_RefreshToken(t *testing.T) {
	t.Run("token not found", func(t *testing.T) {
		service := auth.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			AuthService: service,
		})

		service.On("RefreshToken", t.Context(), mock.Anything).
			Return("", errorx.ErrNotFound)
		resp, err := h.RefreshToken(t.Context(), &pb.RefreshTokenRequest{})
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, pb.RefreshTokenStatus_INVALID_REFRESH_TOKEN, *resp.Status)
	})

	t.Run("error from service", func(t *testing.T) {
		service := auth.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			AuthService: service,
		})

		service.On("RefreshToken", t.Context(), mock.Anything).
			Return("", assert.AnError)
		resp, err := h.RefreshToken(t.Context(), &pb.RefreshTokenRequest{})
		assert.Error(t, err, "expected error from service, got nil")
		assert.NotNil(t, resp)
		assert.Equal(t, pb.RefreshTokenStatus_REFRESH_ERROR, *resp.Status)
	})

	t.Run("success", func(t *testing.T) {
		service := auth.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			AuthService: service,
		})

		service.On("RefreshToken", t.Context(), mock.Anything).
			Return("new-token", nil)
		resp, err := h.RefreshToken(t.Context(), &pb.RefreshTokenRequest{
			RefreshToken: proto.String("123"),
		})
		assert.NoError(t, err, "expected successful refresh, got error")
		assert.NotNil(t, resp)
		assert.Equal(t, pb.RefreshTokenStatus_REFRESH_SUCCESS, *resp.Status)
		assert.Equal(t, "new-token", *resp.Token)
	})
}
