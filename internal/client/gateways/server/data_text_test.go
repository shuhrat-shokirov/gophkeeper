//nolint:dupl,gocritic
package server

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"

	"gophkeeper/internal/client/errorx"
	pb "gophkeeper/proto"
)

func Test_gateway_SaveText(t *testing.T) {
	t.Run("error from server", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("SaveText", mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		g := &gateway{dataServiceClient: mockClient}

		err := g.SaveText(t.Context(), 1, &Text{})
		assert.Error(t, err)
	})

	t.Run("unsuccessful save", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("SaveText", mock.Anything, mock.Anything).
			Return(&pb.Response{Status: pb.ResponseStatus_ERROR.Enum()}, nil)

		g := &gateway{dataServiceClient: mockClient}

		err := g.SaveText(t.Context(), 1, &Text{})
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("SaveText", mock.Anything, mock.Anything).
			Return(&pb.Response{Status: pb.ResponseStatus_SUCCESS.Enum()}, nil)

		g := &gateway{dataServiceClient: mockClient}

		err := g.SaveText(t.Context(), 1, &Text{})
		assert.NoError(t, err)
	})
}

func Test_gateway_GetTextList(t *testing.T) {
	t.Run("error from server", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetTextList", mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		g := &gateway{dataServiceClient: mockClient}

		texts, err := g.GetTextList(t.Context(), 1, 10, 0)
		assert.Error(t, err)
		assert.Nil(t, texts)
	})

	t.Run("not found", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetTextList", mock.Anything, mock.Anything).
			Return(&pb.ListResponse{Status: pb.ResponseListStatus_LIST_NOT_FOUND.Enum()}, nil)

		g := &gateway{dataServiceClient: mockClient}

		texts, err := g.GetTextList(t.Context(), 1, 10, 0)
		assert.ErrorIs(t, err, errorx.ErrNotFound)
		assert.Nil(t, texts)
	})

	t.Run("unsuccessful get", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetTextList", mock.Anything, mock.Anything).
			Return(&pb.ListResponse{Status: pb.ResponseListStatus_LIST_ERROR.Enum()}, nil)

		g := &gateway{dataServiceClient: mockClient}

		texts, err := g.GetTextList(t.Context(), 1, 10, 0)
		assert.Error(t, err)
		assert.Nil(t, texts)
	})

	t.Run("success", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetTextList", mock.Anything, mock.Anything).
			Return(&pb.ListResponse{
				Status: pb.ResponseListStatus_LIST_SUCCESS.Enum(),
				Items: []*pb.ListResp{
					{Id: proto.Int64(1), Title: proto.String("Text1")},
					{Id: proto.Int64(2), Title: proto.String("Text2")},
				},
			}, nil)

		g := &gateway{dataServiceClient: mockClient}

		texts, err := g.GetTextList(t.Context(), 1, 10, 0)
		assert.NoError(t, err)
		assert.Len(t, texts, 2)
		assert.Equal(t, "Text1", texts[0].Title)
		assert.Equal(t, "Text2", texts[1].Title)
	})
}

func Test_gateway_GetTextByID(t *testing.T) {
	t.Run("error from server", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetTextByID", mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		g := &gateway{dataServiceClient: mockClient}

		text, err := g.GetTextByID(t.Context(), 1, 1)
		assert.Error(t, err)
		assert.Nil(t, text)
	})

	t.Run("not found", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetTextByID", mock.Anything, mock.Anything).
			Return(&pb.TextDataResponse{Status: pb.ResponseListStatus_LIST_NOT_FOUND.Enum()}, nil)

		g := &gateway{dataServiceClient: mockClient}

		text, err := g.GetTextByID(t.Context(), 1, 1)
		assert.ErrorIs(t, err, errorx.ErrNotFound)
		assert.Nil(t, text)
	})

	t.Run("unsuccessful get", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetTextByID", mock.Anything, mock.Anything).
			Return(&pb.TextDataResponse{Status: pb.ResponseListStatus_LIST_ERROR.Enum()}, nil)

		g := &gateway{dataServiceClient: mockClient}

		text, err := g.GetTextByID(t.Context(), 1, 1)
		assert.Error(t, err)
		assert.Nil(t, text)
	})

	t.Run("success", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		createdAt := time.Now().UnixNano()

		mockClient.On("GetTextByID", mock.Anything, mock.Anything).
			Return(&pb.TextDataResponse{
				Id:     proto.Int64(1),
				Status: pb.ResponseListStatus_LIST_SUCCESS.Enum(),
				Data: &pb.TextData{
					Meta: &pb.BaseData{
						Title:     proto.String("Test Title"),
						Note:      proto.String("Note content"),
						CreatedAt: proto.Int64(createdAt),
					},
					Content: proto.String("Text content"),
				},
			}, nil)

		g := &gateway{dataServiceClient: mockClient}

		text, err := g.GetTextByID(t.Context(), 1, 1)
		assert.NoError(t, err)
		assert.NotNil(t, text)
		assert.Equal(t, int64(1), text.ID)
		assert.Equal(t, "Test Title", text.Title)
		assert.Equal(t, "Note content", text.Note)
		assert.Equal(t, "Text content", text.Content)
	})
}
