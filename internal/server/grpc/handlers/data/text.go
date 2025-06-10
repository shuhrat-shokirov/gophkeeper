package data

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	"gophkeeper/internal/server/services/data"
	pb "gophkeeper/proto"
)

func (h *handler) SaveText(ctx context.Context, request *pb.TextData) (*pb.Response, error) {
	err := h.dataService.SaveText(ctx, &data.TextData{
		UserID:    request.Meta.GetUserId(),
		Title:     request.Meta.GetTitle(),
		Content:   request.GetContent(),
		Note:      request.Meta.GetNote(),
		CreatedAt: request.Meta.GetCreatedAt(),
	})
	if err != nil {
		return &pb.Response{
			Status:  pb.ResponseStatus_ERROR.Enum(),
			Message: proto.String("Failed to save text data: " + err.Error()),
		}, fmt.Errorf("failed to save text data: %w", err)
	}

	return &pb.Response{
		Status:  pb.ResponseStatus_SUCCESS.Enum(),
		Message: proto.String("Text data saved successfully"),
	}, nil

}
