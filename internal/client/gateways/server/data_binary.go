//nolint:dupl,gocritic
package server

import (
	"context"
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"

	"gophkeeper/internal/client/errorx"
	pb "gophkeeper/proto"
)

func (g *gateway) SaveBinary(ctx context.Context, userID int64, data *Binary) error {
	binary, err := g.dataServiceClient.SaveBinary(ctx, &pb.BinaryData{
		Meta: &pb.BaseData{
			UserId:    proto.Int64(userID),
			Title:     proto.String(data.Title),
			Note:      proto.String(data.Note),
			CreatedAt: proto.Int64(data.CreatedAt.UnixNano()),
		},
		Data: data.Content,
	})
	if err != nil {
		return fmt.Errorf("save binary error: %w", err)
	}

	if binary.GetStatus() != pb.ResponseStatus_SUCCESS {
		return fmt.Errorf("save binary failed: %s", binary.GetMessage())
	}

	return nil
}

func (g *gateway) GetBinaryList(ctx context.Context, userID, limit, offset int64) ([]ListItem, error) {
	binaryList, err := g.dataServiceClient.GetBinaryList(ctx, &pb.ListRequest{
		UserId: proto.Int64(userID),
		Limit:  proto.Int64(limit),
		Offset: proto.Int64(offset),
	})
	if err != nil {
		return nil, fmt.Errorf("get binary list error: %w", err)
	}

	if binaryList.GetStatus() == pb.ResponseListStatus_LIST_NOT_FOUND {
		return nil, errorx.ErrNotFound
	}

	if binaryList.GetStatus() != pb.ResponseListStatus_LIST_SUCCESS {
		return nil, fmt.Errorf("failed to get binary list: %s", binaryList.GetMessage())
	}

	var items = make([]ListItem, 0, len(binaryList.GetItems()))
	for _, item := range binaryList.GetItems() {
		items = append(items, ListItem{
			ID:        item.GetId(),
			Title:     item.GetTitle(),
			CreatedAt: item.GetCreatedAt(),
			UpdatedAt: item.GetUpdatedAt(),
		})
	}

	return items, nil
}

func (g *gateway) GetBinaryByID(ctx context.Context, userID, id int64) (*BinaryInfo, error) {
	binary, err := g.dataServiceClient.GetBinaryByID(ctx, &pb.IDRequest{
		Id:     proto.Int64(id),
		UserId: proto.Int64(userID),
	})
	if err != nil {
		return nil, fmt.Errorf("get binary by ID error: %w", err)
	}

	if binary.GetStatus() == pb.ResponseListStatus_LIST_NOT_FOUND {
		return nil, errorx.ErrNotFound
	}

	if binary.GetStatus() != pb.ResponseListStatus_LIST_SUCCESS {
		return nil, fmt.Errorf("failed to get binary by ID: %s", binary.GetMessage())
	}

	return &BinaryInfo{
		ID: binary.GetId(),
		Binary: Binary{
			Title:     binary.Data.Meta.GetTitle(),
			Content:   binary.Data.GetData(),
			Note:      binary.Data.Meta.GetNote(),
			CreatedAt: time.Unix(0, binary.Data.GetMeta().GetCreatedAt()),
		},
		UpdatedAt: time.Unix(0, binary.GetUpdatedAt()),
	}, nil
}
