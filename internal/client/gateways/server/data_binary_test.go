package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"

	"gophkeeper/internal/client/errorx"
	pb "gophkeeper/proto"
)

func Test_gateway_SaveBinary(t *testing.T) {
	t.Run("err from server", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("SaveBinary", mock.Anything, mock.Anything).Return(nil, assert.AnError)

		g := &gateway{
			dataServiceClient: mockClient,
		}

		err := g.SaveBinary(t.Context(), 1, &Binary{})
		assert.Error(t, err, "expected error from server, got nil")
	})

	t.Run("not successful save", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("SaveBinary", mock.Anything, mock.Anything).Return(&pb.Response{
			Status: pb.ResponseStatus_ERROR.Enum(),
		}, nil)

		g := &gateway{
			dataServiceClient: mockClient,
		}

		err := g.SaveBinary(t.Context(), 1, &Binary{})
		assert.Error(t, err, "expected error for unsuccessful save, got nil")
	})

	t.Run("success", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("SaveBinary", mock.Anything, mock.Anything).Return(&pb.Response{
			Status: pb.ResponseStatus_SUCCESS.Enum(),
		}, nil)

		g := &gateway{
			dataServiceClient: mockClient,
		}

		err := g.SaveBinary(t.Context(), 1, &Binary{})
		assert.NoError(t, err, "expected successful save, got error")
	})
}

func Test_gateway_GetBinaryList(t *testing.T) {
	t.Run("err from server", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetBinaryList", mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		g := &gateway{
			dataServiceClient: mockClient,
		}

		binaries, err := g.GetBinaryList(t.Context(), 1, 10, 0)
		assert.Error(t, err, "expected error from server, got nil")
		assert.Nil(t, binaries, "expected nil response on error")
	})

	t.Run("not found", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetBinaryList", mock.Anything, mock.Anything).
			Return(&pb.ListResponse{
				Status: pb.ResponseListStatus_LIST_NOT_FOUND.Enum(),
			}, nil)

		g := &gateway{
			dataServiceClient: mockClient,
		}

		binaries, err := g.GetBinaryList(t.Context(), 1, 10, 0)
		assert.Error(t, err, "expected no error for not found status")
		assert.Empty(t, binaries, "expected empty list for not found status")
		assert.Equal(t, 0, len(binaries), "expected empty binary list")
		assert.Equal(t, err, errorx.ErrNotFound)
	})

	t.Run("not successful get", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetBinaryList", mock.Anything, mock.Anything).
			Return(&pb.ListResponse{
				Status: pb.ResponseListStatus_LIST_ERROR.Enum(),
			}, nil)

		g := &gateway{
			dataServiceClient: mockClient,
		}

		binaries, err := g.GetBinaryList(t.Context(), 1, 10, 0)
		assert.Error(t, err, "expected error for unsuccessful get, got nil")
		assert.Empty(t, binaries, "expected empty list for unsuccessful get")
	})

	t.Run("success", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetBinaryList", mock.Anything, mock.Anything).
			Return(&pb.ListResponse{
				Status: pb.ResponseListStatus_LIST_SUCCESS.Enum(),
				Items: []*pb.ListResp{
					{
						Id:    proto.Int64(1),
						Title: proto.String("Test Binary"),
					},
					{
						Id:    proto.Int64(2),
						Title: proto.String("Another Binary"),
					},
				},
			}, nil)

		g := &gateway{
			dataServiceClient: mockClient,
		}

		binaries, err := g.GetBinaryList(t.Context(), 1, 10, 0)
		assert.NoError(t, err, "expected successful get, got error")
		assert.NotNil(t, binaries, "expected non-nil response on success")
		assert.Len(t, binaries, 2, "expected two binaries in the list")
		assert.Equal(t, "Test Binary", binaries[0].Title, "expected first binary title to match")
		assert.Equal(t, "Another Binary", binaries[1].Title, "expected second binary title to match")
	})
}

func Test_gateway_GetBinaryByID(t *testing.T) {
	t.Run("err from server", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetBinaryByID", mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		g := &gateway{
			dataServiceClient: mockClient,
		}

		binary, err := g.GetBinaryByID(t.Context(), 1, 1)
		assert.Error(t, err, "expected error from server, got nil")
		assert.Nil(t, binary, "expected nil response on error")
	})

	t.Run("not found", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetBinaryByID", mock.Anything, mock.Anything).
			Return(&pb.BinaryDataResponse{
				Status: pb.ResponseListStatus_LIST_NOT_FOUND.Enum(),
			}, nil)

		g := &gateway{
			dataServiceClient: mockClient,
		}

		binary, err := g.GetBinaryByID(t.Context(), 1, 1)
		assert.Error(t, err, "expected error for not found status")
		assert.Nil(t, binary, "expected nil response for not found status")
		assert.Equal(t, err, errorx.ErrNotFound)
	})

	t.Run("not successful get", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetBinaryByID", mock.Anything, mock.Anything).
			Return(&pb.BinaryDataResponse{
				Status: pb.ResponseListStatus_LIST_ERROR.Enum(),
			}, nil)

		g := &gateway{
			dataServiceClient: mockClient,
		}

		binary, err := g.GetBinaryByID(t.Context(), 1, 1)
		assert.Error(t, err, "expected error for unsuccessful get, got nil")
		assert.Nil(t, binary, "expected nil response for unsuccessful get")
	})

	t.Run("success", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetBinaryByID", mock.Anything, mock.Anything).
			Return(&pb.BinaryDataResponse{
				Id:     proto.Int64(1),
				Status: pb.ResponseListStatus_LIST_SUCCESS.Enum(),
				Data: &pb.BinaryData{
					Meta: &pb.BaseData{
						UserId: proto.Int64(1),
						Title:  proto.String("Test Binary"),
						Note:   proto.String("This is a test binary"),
					},
					Data: []byte("test binary data"),
				},
			}, nil)

		g := &gateway{
			dataServiceClient: mockClient,
		}

		binary, err := g.GetBinaryByID(t.Context(), 1, 1)
		assert.NoError(t, err, "expected successful get, got error")
		assert.NotNil(t, binary, "expected non-nil response on success")
		assert.Equal(t, "Test Binary", binary.Title, "expected binary title to match")
		assert.Equal(t, "This is a test binary", binary.Note, "expected binary note to match")
		assert.Equal(t, []byte("test binary data"), binary.Content, "expected binary data to match")
		assert.Equal(t, int64(1), binary.ID, "expected binary user ID to match")
	})
}
