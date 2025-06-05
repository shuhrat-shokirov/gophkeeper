package data

import (
	"context"
	"fmt"
	"time"

	"gophkeeper/internal/server/repositories/texts"
	"gophkeeper/pkg/aes"
)

func (s *service) SaveText(ctx context.Context, data *TextData) error {
	_, err := s.textRepo.Save(ctx, &texts.TextData{
		UserID:    data.UserID,
		Title:     data.Title,
		Content:   aes.MustEncrypt(data.Content),
		Note:      aes.MustEncrypt(data.Note),
		CreatedAt: time.Unix(0, data.CreatedAt),
	})
	if err != nil {
		return fmt.Errorf("save text: %w", err)
	}

	return nil
}
