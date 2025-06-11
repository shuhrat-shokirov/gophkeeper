package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"

	"gophkeeper/internal/server/services/auth"
	pb "gophkeeper/proto"
)

func Test_handler_Logout(t *testing.T) {
	t.Run("err from service", func(t *testing.T) {
		service := auth.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			AuthService: service,
		})

		service.On("Logout", t.Context(), mock.Anything).Return(assert.AnError)
		resp, err := h.Logout(t.Context(), &pb.LogoutRequest{})
		assert.Error(t, err, "expected error from service, got nil")
		assert.NotNil(t, resp)
		assert.Equal(t, pb.LogoutStatus_LOGOUT_ERROR, *resp.Status, "expected error status in response")
	})

	t.Run("success", func(t *testing.T) {
		service := auth.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			AuthService: service,
		})

		service.On("Logout", t.Context(), mock.Anything).Return(nil)
		resp, err := h.Logout(t.Context(), &pb.LogoutRequest{
			Token: proto.String("test-token"),
		})
		assert.NoError(t, err, "expected successful logout, got error")
		assert.NotNil(t, resp)
		assert.Equal(t, pb.LogoutStatus_LOGOUT_SUCCESS, *resp.Status, "expected OK status in response")
	})
}
