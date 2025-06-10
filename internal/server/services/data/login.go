package data

import (
	"context"
	"fmt"
	"time"

	"gophkeeper/internal/server/repositories/logins"
	"gophkeeper/pkg/aes"
)

func (s *service) SaveLogin(ctx context.Context, data *LoginData) error {
	_, err := s.loginsRepo.Save(ctx, &logins.LoginData{
		UserID:    data.UserId,
		Login:     aes.MustEncrypt(data.Login),
		Password:  aes.MustEncrypt(data.Password),
		Title:     data.Title,
		Note:      aes.MustEncrypt(data.Note),
		CreatedAt: time.Unix(0, data.CreatedAt),
	})
	if err != nil {
		return fmt.Errorf("failed to save login: %w", err)
	}

	return nil
}

func (s *service) GetLoginList(ctx context.Context, userID, limit, offset int64) ([]LoginListItem, error) {
	list, err := s.loginsRepo.List(ctx, userID, logins.Pagination{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get login list: %w", err)
	}

	var loginDataList = make([]LoginListItem, 0, len(list))
	for _, login := range list {
		loginDataList = append(loginDataList,
			LoginListItem{
				ID:        login.ID,
				Title:     login.Title,
				CreatedAt: login.CreatedAt.UnixNano(),
				UpdatedAt: login.UpdatedAt.UnixNano(),
			})
	}

	return loginDataList, nil
}

func (s *service) GetLoginByID(ctx context.Context, id int64) (*LoginInfo, error) {
	info, err := s.loginsRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get login by ID: %w", err)
	}

	return &LoginInfo{
		ID: info.ID,
		LoginData: LoginData{
			UserId:    info.UserID,
			Login:     aes.MustDecrypt(info.Login),
			Password:  aes.MustDecrypt(info.Password),
			Title:     info.Title,
			Note:      aes.MustDecrypt(info.Note),
			CreatedAt: info.CreatedAt.UnixNano(),
		},
		UpdatedAt: info.UpdatedAt.UnixNano(),
	}, nil
}
