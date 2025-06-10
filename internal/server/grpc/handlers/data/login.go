package data

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	"gophkeeper/internal/server/services/data"
	pb "gophkeeper/proto"
)

func (h *handler) SaveLogin(ctx context.Context, request *pb.LoginData) (*pb.Response, error) {
	err := h.dataService.SaveLogin(ctx, &data.LoginData{
		UserId:    request.Meta.GetUserId(),
		Login:     request.GetLogin(),
		Password:  request.GetPassword(),
		Title:     request.Meta.GetTitle(),
		Note:      request.Meta.GetNote(),
		CreatedAt: request.Meta.GetCreatedAt(),
	})
	if err != nil {
		return &pb.Response{
			Status:  pb.ResponseStatus_ERROR.Enum(),
			Message: proto.String("Failed to save login: " + err.Error()),
		}, fmt.Errorf("failed to save login: %w", err)
	}

	return &pb.Response{
		Status:  pb.ResponseStatus_SUCCESS.Enum(),
		Message: proto.String("Login saved successfully"),
	}, nil
}
