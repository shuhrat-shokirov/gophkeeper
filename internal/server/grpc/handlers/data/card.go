package data

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	serviceData "gophkeeper/internal/server/services/data"
	pb "gophkeeper/proto"
)

func (h *handler) SaveCard(ctx context.Context, data *pb.CardData) (*pb.Response, error) {
	err := h.dataService.SaveCard(ctx, &serviceData.CardData{
		UserID:    data.Meta.GetUserId(),
		Pan:       data.GetPan(),
		Cvv:       data.GetCvv(),
		Expiry:    data.GetExpiry(),
		Title:     data.Meta.GetTitle(),
		Note:      data.Meta.GetNote(),
		CreatedAt: data.Meta.GetCreatedAt(),
	})
	if err != nil {
		return &pb.Response{
			Status:  pb.ResponseStatus_ERROR.Enum(),
			Message: proto.String("failed to save card data: " + err.Error()),
		}, fmt.Errorf("failed to save card data: %w", err)
	}

	return &pb.Response{
		Status:  pb.ResponseStatus_SUCCESS.Enum(),
		Message: proto.String("card data saved successfully"),
	}, nil
}
