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

func Test_gateway_SaveCard(t *testing.T) {
	t.Run("error from server", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("SaveCard", mock.Anything, mock.Anything).Return(nil, assert.AnError)

		g := &gateway{dataServiceClient: mockClient}

		err := g.SaveCard(t.Context(), 1, &Card{})
		assert.Error(t, err)
	})

	t.Run("not successful save", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("SaveCard", mock.Anything, mock.Anything).
			Return(&pb.Response{Status: pb.ResponseStatus_ERROR.Enum()}, nil)

		g := &gateway{dataServiceClient: mockClient}

		err := g.SaveCard(t.Context(), 1, &Card{})
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("SaveCard", mock.Anything, mock.Anything).
			Return(&pb.Response{Status: pb.ResponseStatus_SUCCESS.Enum()}, nil)

		g := &gateway{dataServiceClient: mockClient}

		err := g.SaveCard(t.Context(), 1, &Card{})
		assert.NoError(t, err)
	})
}

func Test_gateway_GetCardList(t *testing.T) {
	t.Run("error from server", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetCardList", mock.Anything, mock.Anything).Return(nil, assert.AnError)

		g := &gateway{dataServiceClient: mockClient}

		cards, err := g.GetCardList(t.Context(), 1, 10, 0)
		assert.Error(t, err)
		assert.Nil(t, cards)
	})

	t.Run("not found", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetCardList", mock.Anything, mock.Anything).
			Return(&pb.ListResponse{Status: pb.ResponseListStatus_LIST_NOT_FOUND.Enum()}, nil)

		g := &gateway{dataServiceClient: mockClient}

		cards, err := g.GetCardList(t.Context(), 1, 10, 0)
		assert.ErrorIs(t, err, errorx.ErrNotFound)
		assert.Nil(t, cards)
	})

	t.Run("not successful list", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetCardList", mock.Anything, mock.Anything).
			Return(&pb.ListResponse{Status: pb.ResponseListStatus_LIST_ERROR.Enum()}, nil)

		g := &gateway{dataServiceClient: mockClient}

		cards, err := g.GetCardList(t.Context(), 1, 10, 0)
		assert.Error(t, err)
		assert.Nil(t, cards)
	})

	t.Run("success", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetCardList", mock.Anything, mock.Anything).
			Return(&pb.ListResponse{
				Status: pb.ResponseListStatus_LIST_SUCCESS.Enum(),
				Items: []*pb.ListResp{
					{Id: proto.Int64(1), Title: proto.String("Card 1")},
					{Id: proto.Int64(2), Title: proto.String("Card 2")},
				},
			}, nil)

		g := &gateway{dataServiceClient: mockClient}

		cards, err := g.GetCardList(t.Context(), 1, 10, 0)
		assert.NoError(t, err)
		assert.Len(t, cards, 2)
		assert.Equal(t, "Card 1", cards[0].Title)
		assert.Equal(t, "Card 2", cards[1].Title)
	})
}

func Test_gateway_GetCardByID(t *testing.T) {
	t.Run("error from server", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetCardByID", mock.Anything, mock.Anything).Return(nil, assert.AnError)

		g := &gateway{dataServiceClient: mockClient}

		card, err := g.GetCardByID(t.Context(), 1, 1)
		assert.Error(t, err)
		assert.Nil(t, card)
	})

	t.Run("not found", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetCardByID", mock.Anything, mock.Anything).
			Return(&pb.CardDataResponse{Status: pb.ResponseListStatus_LIST_NOT_FOUND.Enum()}, nil)

		g := &gateway{dataServiceClient: mockClient}

		card, err := g.GetCardByID(t.Context(), 1, 1)
		assert.ErrorIs(t, err, errorx.ErrNotFound)
		assert.Nil(t, card)
	})

	t.Run("not successful get", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetCardByID", mock.Anything, mock.Anything).
			Return(&pb.CardDataResponse{Status: pb.ResponseListStatus_LIST_ERROR.Enum()}, nil)

		g := &gateway{dataServiceClient: mockClient}

		card, err := g.GetCardByID(t.Context(), 1, 1)
		assert.Error(t, err)
		assert.Nil(t, card)
	})

	t.Run("success", func(t *testing.T) {
		mockClient := new(MockDataServiceClient)
		defer mockClient.AssertExpectations(t)

		mockClient.On("GetCardByID", mock.Anything, mock.Anything).
			Return(&pb.CardDataResponse{
				Id:     proto.Int64(1),
				Status: pb.ResponseListStatus_LIST_SUCCESS.Enum(),
				Data: &pb.CardData{
					Meta: &pb.BaseData{
						Title:     proto.String("My Card"),
						Note:      proto.String("Visa"),
						CreatedAt: proto.Int64(time.Now().UnixNano()),
					},
					Pan:    proto.String("1234567812345678"),
					Expiry: proto.String("12/30"),
					Cvv:    proto.String("999"),
				},
			}, nil)

		g := &gateway{dataServiceClient: mockClient}

		card, err := g.GetCardByID(t.Context(), 1, 1)
		assert.NoError(t, err)
		assert.NotNil(t, card)
		assert.Equal(t, "My Card", card.Title)
		assert.Equal(t, "1234567812345678", card.Pan)
		assert.Equal(t, "12/30", card.Expiry)
		assert.Equal(t, "999", card.Cvv)
		assert.Equal(t, "Visa", card.Note)
		assert.Equal(t, int64(1), card.ID)
	})
}
