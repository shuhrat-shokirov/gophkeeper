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

func (g *gateway) SaveLoginAndPass(ctx context.Context, userID int64, pass *LoginAndPass) error {
	login, err := g.dataServiceClient.SaveLogin(ctx, &pb.LoginData{
		Meta: &pb.BaseData{
			UserId:    proto.Int64(userID),
			Title:     proto.String(pass.Title),
			Note:      proto.String(pass.Note),
			CreatedAt: proto.Int64(pass.CreatedAt.UnixNano()),
		},
		Login:    proto.String(pass.Login),
		Password: proto.String(pass.Pass),
	})
	if err != nil {
		return fmt.Errorf("save login: %w", err)
	}

	if login.GetStatus() != pb.ResponseStatus_SUCCESS {
		return fmt.Errorf("failed to save login: %s", login.GetMessage())
	}

	return nil
}

func (g *gateway) GetLoginList(ctx context.Context, userID int64, limit, offset int64) ([]ListItem, error) {
	loginList, err := g.dataServiceClient.GetLoginList(ctx, &pb.ListRequest{
		UserId: proto.Int64(userID),
		Limit:  proto.Int64(limit),
		Offset: proto.Int64(offset),
	})
	if err != nil {
		return nil, fmt.Errorf("get login list: %w", err)
	}

	if loginList.GetStatus() == pb.ResponseListStatus_LIST_NOT_FOUND {
		return nil, errorx.ErrNotFound
	}

	if loginList.GetStatus() != pb.ResponseListStatus_LIST_SUCCESS {
		return nil, fmt.Errorf("failed to get login list: %s", loginList.GetMessage())
	}

	var items = make([]ListItem, 0, len(loginList.GetItems()))
	for _, item := range loginList.GetItems() {
		items = append(items, ListItem{
			ID:        item.GetId(),
			Title:     item.GetTitle(),
			CreatedAt: item.GetCreatedAt(),
			UpdatedAt: item.GetUpdatedAt(),
		})
	}

	return items, nil
}

func (g *gateway) GetLoginByID(ctx context.Context, userID, id int64) (*LoginInfo, error) {
	loginInfo, err := g.dataServiceClient.GetLoginByID(ctx, &pb.IDRequest{
		Id:     proto.Int64(id),
		UserId: proto.Int64(userID),
	})
	if err != nil {
		return nil, fmt.Errorf("get login by ID: %w", err)
	}

	if loginInfo.GetStatus() == pb.ResponseListStatus_LIST_NOT_FOUND {
		return nil, errorx.ErrNotFound
	}

	if loginInfo.GetStatus() != pb.ResponseListStatus_LIST_SUCCESS {
		return nil, fmt.Errorf("failed to get login by ID: %s", loginInfo.GetMessage())
	}

	return &LoginInfo{
		ID: loginInfo.GetId(),
		LoginAndPass: LoginAndPass{
			Login:     loginInfo.Data.GetLogin(),
			Pass:      loginInfo.Data.GetPassword(),
			Title:     loginInfo.Data.Meta.GetTitle(),
			Note:      loginInfo.Data.Meta.GetNote(),
			CreatedAt: time.Unix(0, loginInfo.Data.Meta.GetCreatedAt()),
		},
		UpdatedAt: time.Unix(0, loginInfo.GetUpdatedAt()),
	}, nil
}
