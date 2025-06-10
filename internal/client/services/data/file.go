package data

import (
	"context"
	"fmt"
	"os"
	"time"

	"gophkeeper/internal/client/gateways/server"
)

func (s *service) SaveFile(ctx context.Context, data *FileData) error {
	bytes, err := os.ReadFile(data.Path)
	if err != nil {
		return fmt.Errorf("ошибка чтения файла: %w", err)
	}

	userID, err := s.authService.GetUserID(ctx)
	if err != nil {
		return fmt.Errorf("get user ID: %w", err)
	}

	err = s.serverGateway.SaveBinary(ctx, userID, &server.Binary{
		Title:     data.Title,
		Content:   bytes,
		Note:      data.Note,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("save file: %w", err)
	}

	return nil
}

func (s *service) GetBinaryList(ctx context.Context, limit, offset int64) ([]ListItem, error) {
	userID, err := s.authService.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get user ID: %w", err)
	}

	list, err := s.serverGateway.GetBinaryList(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get file list: %w", err)
	}

	var fileDataList = make([]ListItem, 0, len(list))
	for _, item := range list {
		fileDataList = append(fileDataList, ListItem{
			ID:        item.ID,
			Title:     item.Title,
			CreatedAt: time.Unix(0, item.CreatedAt),
			UpdatedAt: time.Unix(0, item.UpdatedAt),
		})
	}

	return fileDataList, nil
}

func (s *service) GetBinaryByID(ctx context.Context, id int64) (*BinaryInfo, error) {
	userID, err := s.authService.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get user ID: %w", err)
	}

	binary, err := s.serverGateway.GetBinaryByID(ctx, userID, id)
	if err != nil {
		return nil, fmt.Errorf("get file by ID: %w", err)
	}

	return &BinaryInfo{
		ID:        binary.ID,
		Title:     binary.Title,
		Content:   binary.Content,
		Note:      binary.Note,
		CreatedAt: binary.CreatedAt,
		UpdatedAt: binary.UpdatedAt,
	}, nil
}
