//nolint:dupl,gocritic
package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gophkeeper/internal/server/services/data"
	pb "gophkeeper/proto"
)

func Test_handler_SaveText(t *testing.T) {
	t.Run("err from service", func(t *testing.T) {
		service := data.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			DataService: service,
		})

		service.On("SaveText", t.Context(), mock.Anything).Return(assert.AnError)

		resp, err := h.SaveText(t.Context(), &pb.TextData{
			Meta: &pb.BaseData{},
		})
		assert.Error(t, err, "expected error from service, got nil")
		assert.NotNil(t, resp)
		assert.Equal(t, pb.ResponseStatus_ERROR, *resp.Status, "expected STATUS_ERROR status")
	})

	t.Run("success", func(t *testing.T) {
		service := data.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			DataService: service,
		})

		service.On("SaveText", t.Context(), mock.Anything).Return(nil)

		resp, err := h.SaveText(t.Context(), &pb.TextData{
			Meta:    &pb.BaseData{},
			Content: nil,
		})
		assert.NoError(t, err, "expected successful save, got error")
		assert.NotNil(t, resp)
		assert.Equal(t, pb.ResponseStatus_SUCCESS, *resp.Status, "expected STATUS_SUCCESS status")
	})
}
