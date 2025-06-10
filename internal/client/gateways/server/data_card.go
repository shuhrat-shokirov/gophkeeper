//nolint:dupl,gocritic
package server

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/protobuf/proto"

	"gophkeeper/internal/client/errorx"
	pb "gophkeeper/proto"
)

func (g *gateway) SaveCard(ctx context.Context, userID int64, data *Card) error {
	card, err := g.dataServiceClient.SaveCard(ctx, &pb.CardData{
		Meta: &pb.BaseData{
			UserId:    proto.Int64(userID),
			Title:     proto.String(data.Title),
			Note:      proto.String(data.Note),
			CreatedAt: proto.Int64(data.CreatedAt.UnixNano()),
		},
		Pan:    proto.String(data.Pan),
		Expiry: proto.String(data.Expiry),
		Cvv:    proto.String(data.Cvv),
	})
	if err != nil {
		return fmt.Errorf("save card error: %w", err)
	}

	if card.GetStatus() != pb.ResponseStatus_SUCCESS {
		return fmt.Errorf("save card failed: %s", card.GetMessage())
	}

	return nil
}

func (g *gateway) GetCardList(ctx context.Context, userID, limit, offset int64) ([]ListItem, error) {
	cardList, err := g.dataServiceClient.GetCardList(ctx, &pb.ListRequest{
		UserId: proto.Int64(userID),
		Limit:  proto.Int64(limit),
		Offset: proto.Int64(offset),
	})
	if err != nil {
		return nil, fmt.Errorf("get card list error: %w", err)
	}

	if cardList.GetStatus() == pb.ResponseListStatus_LIST_NOT_FOUND {
		return nil, errorx.ErrNotFound
	}

	if cardList.GetStatus() != pb.ResponseListStatus_LIST_SUCCESS {
		return nil, fmt.Errorf("failed to get card list: %s", cardList.GetMessage())
	}

	var items = make([]ListItem, 0, len(cardList.GetItems()))
	for _, item := range cardList.GetItems() {
		items = append(items, ListItem{
			ID:        item.GetId(),
			Title:     item.GetTitle(),
			CreatedAt: item.GetCreatedAt(),
			UpdatedAt: item.GetUpdatedAt(),
		})
	}

	return items, nil
}

func (g *gateway) GetCardByID(ctx context.Context, userID, id int64) (*CardInfo, error) {
	card, err := g.dataServiceClient.GetCardByID(ctx, &pb.IDRequest{
		Id:     proto.Int64(id),
		UserId: proto.Int64(userID),
	})
	if err != nil {
		return nil, fmt.Errorf("get card by ID error: %w", err)
	}

	if card.GetStatus() == pb.ResponseListStatus_LIST_NOT_FOUND {
		return nil, errorx.ErrNotFound
	}

	if card.GetStatus() != pb.ResponseListStatus_LIST_SUCCESS {
		return nil, fmt.Errorf("failed to get card by ID: %s", card.GetMessage())
	}

	return &CardInfo{
		ID: card.GetId(),
		Card: Card{
			Title:     card.Data.Meta.GetTitle(),
			Pan:       card.Data.GetPan(),
			Cvv:       card.Data.GetCvv(),
			Expiry:    card.Data.GetExpiry(),
			Note:      card.Data.GetMeta().GetNote(),
			CreatedAt: time.Unix(0, card.Data.GetMeta().GetCreatedAt()),
		},
		UpdatedAt: time.Unix(0, card.GetUpdatedAt()),
	}, nil
}
