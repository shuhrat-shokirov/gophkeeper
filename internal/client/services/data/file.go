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
		Title:      data.Title,
		Content:    bytes,
		Note:       data.Note,
		ModifiedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("save file: %w", err)
	}

	return nil

}
