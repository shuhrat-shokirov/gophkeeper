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

func (g *gateway) SaveText(ctx context.Context, userID int64, data *Text) error {
	response, err := g.dataServiceClient.SaveText(ctx, &pb.TextData{
		Meta: &pb.BaseData{
			UserId:    proto.Int64(userID),
			Title:     proto.String(data.Title),
			Note:      proto.String(data.Note),
			CreatedAt: proto.Int64(data.CreatedAt.UnixNano()),
		},
		Content: proto.String(data.Content),
	})
	if err != nil {
		return fmt.Errorf("save text: %w", err)
	}

	if response.GetStatus() != pb.ResponseStatus_SUCCESS {
		return fmt.Errorf("save text: status %s", response.GetStatus().String())
	}

	return nil
}

func (g *gateway) GetTextList(ctx context.Context, userID int64, limit, offset int64) ([]ListItem, error) {
	textList, err := g.dataServiceClient.GetTextList(ctx, &pb.ListRequest{
		UserId: proto.Int64(userID),
		Limit:  proto.Int64(limit),
		Offset: proto.Int64(offset),
	})
	if err != nil {
		return nil, fmt.Errorf("get text list: %w", err)
	}

	if textList.GetStatus() == pb.ResponseListStatus_LIST_NOT_FOUND {
		return nil, errorx.ErrNotFound
	}

	if textList.GetStatus() != pb.ResponseListStatus_LIST_SUCCESS {
		return nil, fmt.Errorf("failed to get text list: %s", textList.GetMessage())
	}

	var items = make([]ListItem, 0, len(textList.GetItems()))
	for _, item := range textList.GetItems() {
		items = append(items, ListItem{
			ID:        item.GetId(),
			Title:     item.GetTitle(),
			CreatedAt: item.GetCreatedAt(),
			UpdatedAt: item.GetUpdatedAt(),
		})
	}

	return items, nil
}

func (g *gateway) GetTextByID(ctx context.Context, userID, id int64) (*TextInfo, error) {
	response, err := g.dataServiceClient.GetTextByID(ctx, &pb.IDRequest{
		Id:     proto.Int64(id),
		UserId: proto.Int64(userID),
	})
	if err != nil {
		return nil, fmt.Errorf("get text by ID: %w", err)
	}

	if response.GetStatus() == pb.ResponseListStatus_LIST_NOT_FOUND {
		return nil, errorx.ErrNotFound
	}

	if response.GetStatus() != pb.ResponseListStatus_LIST_SUCCESS {
		return nil, fmt.Errorf("failed to get text by ID: %s", response.GetMessage())
	}

	return &TextInfo{
		ID: response.GetId(),
		Text: Text{
			Title:     response.Data.Meta.GetTitle(),
			Content:   response.Data.GetContent(),
			Note:      response.Data.Meta.GetNote(),
			CreatedAt: time.Unix(0, response.Data.Meta.GetCreatedAt()),
		},
		UpdatedAt: time.Unix(0, response.GetUpdatedAt()),
	}, nil
}
