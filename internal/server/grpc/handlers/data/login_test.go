package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"

	"gophkeeper/internal/server/errorx"
	"gophkeeper/internal/server/services/data"
	pb "gophkeeper/proto"
)

func Test_handler_SaveLogin(t *testing.T) {
	t.Run("err from service", func(t *testing.T) {
		service := data.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			DataService: service,
		})

		service.On("SaveLogin", t.Context(), mock.Anything).Return(assert.AnError)
		resp, err := h.SaveLogin(t.Context(), &pb.LoginData{})
		assert.Error(t, err, "expected error from service, got nil")
		assert.NotNil(t, resp)
		assert.Equal(t, pb.ResponseStatus_ERROR, *resp.Status, "expected error status in response")
	})

	t.Run("success", func(t *testing.T) {
		service := data.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			DataService: service,
		})

		service.On("SaveLogin", t.Context(), mock.Anything).Return(nil)
		resp, err := h.SaveLogin(t.Context(),
			&pb.LoginData{Login: proto.String("testuser"), Password: proto.String("testpass")})
		assert.NoError(t, err, "expected successful save, got error")
		assert.NotNil(t, resp)
		assert.Equal(t, pb.ResponseStatus_SUCCESS, *resp.Status, "expected OK status in response")
	})
}

func Test_handler_GetLoginList(t *testing.T) {
	t.Run("not found", func(t *testing.T) {
		service := data.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			DataService: service,
		})

		service.On("GetLoginList", t.Context(), int64(1), int64(10), int64(0)).
			Return(nil, errorx.ErrNotFound)
		resp, err := h.GetLoginList(t.Context(), &pb.ListRequest{
			UserId: proto.Int64(1),
			Limit:  proto.Int64(10),
			Offset: proto.Int64(0),
		})
		assert.NoError(t, err, "expected no error on not found")
		assert.NotNil(t, resp)
		assert.Equal(t, pb.ResponseListStatus_LIST_NOT_FOUND, *resp.Status, "expected LIST_NOT_FOUND status")
	})

	t.Run("err from service", func(t *testing.T) {
		service := data.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			DataService: service,
		})

		service.On("GetLoginList", t.Context(), int64(1), int64(10), int64(0)).
			Return(nil, assert.AnError)
		resp, err := h.GetLoginList(t.Context(), &pb.ListRequest{
			UserId: proto.Int64(1),
			Limit:  proto.Int64(10),
			Offset: proto.Int64(0),
		})
		assert.Error(t, err, "expected error from service, got nil")
		assert.NotNil(t, resp)
		assert.Equal(t, pb.ResponseListStatus_LIST_ERROR, *resp.Status, "expected LIST_ERROR status")
	})

	t.Run("success", func(t *testing.T) {
		service := data.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			DataService: service,
		})

		expectedLogins := []data.LoginListItem{
			{ID: 1, Title: "testuser", CreatedAt: 1234567890, UpdatedAt: 1234567890},
			{ID: 2, Title: "anotheruser", CreatedAt: 1234567891, UpdatedAt: 1234567891},
		}

		service.On("GetLoginList", t.Context(), int64(1), int64(10), int64(0)).
			Return(expectedLogins, nil)
		resp, err := h.GetLoginList(t.Context(), &pb.ListRequest{
			UserId: proto.Int64(1),
			Limit:  proto.Int64(10),
			Offset: proto.Int64(0),
		})
		assert.NoError(t, err, "expected successful login list retrieval, got error")
		assert.NotNil(t, resp)
		assert.Equal(t, pb.ResponseListStatus_LIST_SUCCESS, *resp.Status, "expected LIST_SUCCESS status")
		assert.Len(t, resp.Items, len(expectedLogins), "expected correct number of logins in response")
		for i, item := range resp.Items {
			assert.Equal(t, expectedLogins[i].ID, *item.Id)
			assert.Equal(t, expectedLogins[i].Title, *item.Title)
			assert.Equal(t, expectedLogins[i].CreatedAt, *item.CreatedAt)
			assert.Equal(t, expectedLogins[i].UpdatedAt, *item.UpdatedAt)
		}
	})
}

func Test_handler_GetLoginByID(t *testing.T) {
	t.Run("not found", func(t *testing.T) {
		service := data.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			DataService: service,
		})

		service.On("GetLoginByID", t.Context(), int64(1)).
			Return(nil, errorx.ErrNotFound)
		resp, err := h.GetLoginByID(t.Context(), &pb.IDRequest{Id: proto.Int64(1)})
		assert.NoError(t, err, "expected no error on not found")
		assert.NotNil(t, resp)
		assert.Equal(t, pb.ResponseListStatus_LIST_NOT_FOUND, *resp.Status, "expected LIST_NOT_FOUND status")
	})

	t.Run("err from service", func(t *testing.T) {
		service := data.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			DataService: service,
		})

		service.On("GetLoginByID", t.Context(), int64(1)).
			Return(nil, assert.AnError)
		resp, err := h.GetLoginByID(t.Context(), &pb.IDRequest{Id: proto.Int64(1)})
		assert.Error(t, err, "expected error from service, got nil")
		assert.NotNil(t, resp)
		assert.Equal(t, pb.ResponseListStatus_LIST_ERROR, *resp.Status, "expected LIST_ERROR status")
	})

	t.Run("success", func(t *testing.T) {
		service := data.NewMockService()
		defer service.AssertExpectations(t)

		h := New(Params{
			DataService: service,
		})

		expectedLogin := &data.LoginInfo{
			ID: 1,
			LoginData: data.LoginData{
				UserId:    1,
				Login:     "testuser",
				Password:  "testpass",
				Title:     "Test User",
				Note:      "This is a test note",
				CreatedAt: 1234567890,
			},
			UpdatedAt: 1234567890,
		}

		service.On("GetLoginByID", t.Context(), int64(1)).
			Return(expectedLogin, nil)
		resp, err := h.GetLoginByID(t.Context(), &pb.IDRequest{Id: proto.Int64(1)})
		assert.NoError(t, err, "expected successful login retrieval, got error")
		assert.NotNil(t, resp)
		assert.Equal(t, pb.ResponseListStatus_LIST_SUCCESS, *resp.Status, "expected LIST_SUCCESS status")
		assert.Equal(t, expectedLogin.ID, *resp.Id)
		assert.Equal(t, expectedLogin.UserId, *resp.Data.Meta.UserId)
		assert.Equal(t, expectedLogin.Title, *resp.Data.Meta.Title)
		assert.Equal(t, expectedLogin.Note, *resp.Data.Meta.Note)
		assert.Equal(t, expectedLogin.CreatedAt, *resp.Data.Meta.CreatedAt)
	})
}
