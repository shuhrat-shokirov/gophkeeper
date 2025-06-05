package data

import (
	"context"
	"fmt"
	"time"

	"gophkeeper/internal/server/repositories/card"
	"gophkeeper/pkg/aes"
)

func (s *service) SaveCard(ctx context.Context, data *CardData) error {
	_, err := s.cardRepo.Save(ctx, &card.Data{
		UserID:    data.UserID,
		Pan:       aes.MustEncrypt(data.Pan),
		Cvv:       aes.MustEncrypt(data.Cvv),
		Expiry:    aes.MustEncrypt(data.Expiry),
		Title:     data.Title,
		Note:      aes.MustEncrypt(data.Note),
		CreatedAt: time.Unix(0, data.CreatedAt),
	})
	if err != nil {
		return fmt.Errorf("save card: %w", err)
	}

	return nil

}
