package data

import (
	"context"
	"fmt"
	"time"

	"gophkeeper/internal/client/gateways/server"
)

func (s *service) SaveCard(ctx context.Context, data *CardData) error {
	userID, err := s.authService.GetUserID(ctx)
	if err != nil {
		return fmt.Errorf("get user ID: %w", err)
	}

	err = s.serverGateway.SaveCard(ctx, userID, &server.Card{
		Title:      data.Title,
		Pan:        data.Pan,
		Cvv:        data.Cvv,
		Expiry:     data.Expiry,
		Note:       data.Note,
		ModifiedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("save card: %w", err)
	}

	return nil
}
