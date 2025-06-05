package data

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	serviceData "gophkeeper/internal/server/services/data"
	pb "gophkeeper/proto"
)

func (h *handler) SaveBinary(ctx context.Context, data *pb.BinaryData) (*pb.Response, error) {
	err := h.dataService.SaveBinary(ctx, &serviceData.BinaryData{
		UserID:    data.Meta.GetUserId(),
		Title:     data.Meta.GetTitle(),
		Content:   data.GetData(),
		Note:      data.Meta.GetNote(),
		CreatedAt: data.Meta.GetCreatedAt(),
	})
	if err != nil {
		return &pb.Response{
			Status:  pb.ResponseStatus_ERROR.Enum(),
			Message: proto.String("failed to save binary data: " + err.Error()),
		}, fmt.Errorf("failed to save binary data: %w", err)
	}

	return &pb.Response{
		Status:  pb.ResponseStatus_SUCCESS.Enum(),
		Message: proto.String("binary data saved successfully"),
	}, nil

}
