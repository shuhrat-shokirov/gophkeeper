package data

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/protobuf/proto"

	"gophkeeper/internal/server/errorx"
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

func (h *handler) GetLoginList(ctx context.Context, request *pb.ListRequest) (*pb.ListResponse, error) {
	logins, err := h.dataService.GetLoginList(ctx, request.GetUserId(), request.GetLimit(), request.GetOffset())
	if err != nil {
		if errors.Is(err, errorx.ErrNotFound) {
			return &pb.ListResponse{
				Status:  pb.ResponseListStatus_LIST_NOT_FOUND.Enum(),
				Message: proto.String("No logins found"),
			}, nil
		}

		return &pb.ListResponse{
			Status:  pb.ResponseListStatus_LIST_ERROR.Enum(),
			Message: proto.String("Failed to retrieve login list: " + err.Error()),
		}, fmt.Errorf("failed to retrieve login list: %w", err)
	}

	var list = make([]*pb.ListResp, 0, len(logins))
	for _, login := range logins {
		list = append(list, &pb.ListResp{
			Id:        proto.Int64(login.ID),
			Title:     proto.String(login.Title),
			CreatedAt: proto.Int64(login.CreatedAt),
			UpdatedAt: proto.Int64(login.UpdatedAt),
		})
	}

	return &pb.ListResponse{
		Status:  pb.ResponseListStatus_LIST_SUCCESS.Enum(),
		Items:   list,
		Message: proto.String("Login list retrieved successfully"),
	}, nil
}

func (h *handler) GetLoginByID(ctx context.Context, request *pb.IDRequest) (*pb.LoginDataResponse, error) {
	info, err := h.dataService.GetLoginByID(ctx, request.GetId())
	if err != nil {
		if errors.Is(err, errorx.ErrNotFound) {
			return &pb.LoginDataResponse{
				Status:  pb.ResponseListStatus_LIST_NOT_FOUND.Enum(),
				Message: proto.String("Login not found"),
			}, nil
		}

		return &pb.LoginDataResponse{
			Status:  pb.ResponseListStatus_LIST_ERROR.Enum(),
			Message: proto.String("Failed to retrieve login: " + err.Error()),
		}, fmt.Errorf("failed to retrieve login by ID: %w", err)
	}

	return &pb.LoginDataResponse{
		Id: proto.Int64(info.ID),
		Data: &pb.LoginData{
			Meta: &pb.BaseData{
				UserId:    proto.Int64(info.UserId),
				Title:     proto.String(info.Title),
				Note:      proto.String(info.Note),
				CreatedAt: proto.Int64(info.CreatedAt),
			},
			Login:    proto.String(info.Login),
			Password: proto.String(info.Password),
		},
		UpdatedAt: proto.Int64(info.UpdatedAt),
		Status:    pb.ResponseListStatus_LIST_SUCCESS.Enum(),
		Message:   proto.String("Login retrieved successfully"),
	}, nil
}
