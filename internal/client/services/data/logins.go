package data

import (
	"context"
	"fmt"
	"time"

	"gophkeeper/internal/client/gateways/server"
)

func (s *service) SaveLogin(ctx context.Context, data *LoginData) error {
	userID, err := s.authService.GetUserID(ctx)
	if err != nil {
		return fmt.Errorf("get user id: %w", err)
	}

	err = s.serverGateway.SaveLoginAndPass(ctx, userID, &server.LoginAndPass{
		Login:     data.Login,
		Pass:      data.Pass,
		Title:     data.Title,
		Note:      data.Note,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("save login: %w", err)
	}

	return nil
}

func (s *service) GetLoginList(ctx context.Context, limit, offset int64) ([]ListItem, error) {
	userID, err := s.authService.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get user id: %w", err)
	}

	list, err := s.serverGateway.GetLoginList(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get list login: %w", err)
	}

	var loginDataList = make([]ListItem, 0, len(list))
	for _, item := range list {
		loginDataList = append(loginDataList, ListItem{
			ID:        item.ID,
			Title:     item.Title,
			CreatedAt: time.Unix(0, item.CreatedAt),
			UpdatedAt: time.Unix(0, item.UpdatedAt),
		})
	}

	return loginDataList, nil
}

func (s *service) GetLoginByID(ctx context.Context, id int64) (*LoginInfo, error) {
	userID, err := s.authService.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get user id: %w", err)
	}

	info, err := s.serverGateway.GetLoginByID(ctx, userID, id)
	if err != nil {
		return nil, fmt.Errorf("get login by ID: %w", err)
	}

	return &LoginInfo{
		ID: info.ID,
		LoginData: LoginData{
			Login: info.Login,
			Pass:  info.Pass,
			Title: info.Title,
			Note:  info.Note,
		},
		CreatedAt: info.CreatedAt,
		UpdatedAt: info.UpdatedAt,
	}, nil
}
