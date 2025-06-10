package data

import (
	"context"
	"fmt"
	"time"

	"gophkeeper/internal/client/gateways/server"
)

func (s *service) SaveText(ctx context.Context, data *TextData) error {
	userID, err := s.authService.GetUserID(ctx)
	if err != nil {
		return fmt.Errorf("get user id: %w", err)
	}

	err = s.serverGateway.SaveText(ctx, userID, &server.Text{
		Title:      data.Title,
		Content:    data.Content,
		Note:       data.Note,
		ModifiedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("save text: %w", err)
	}

	return nil
}
