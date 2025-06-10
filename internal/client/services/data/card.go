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
		Title:     data.Title,
		Pan:       data.Pan,
		Cvv:       data.Cvv,
		Expiry:    data.Expiry,
		Note:      data.Note,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("save card: %w", err)
	}

	return nil
}

func (s *service) GetCardList(ctx context.Context, limit, offset int64) ([]ListItem, error) {
	userID, err := s.authService.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get user ID: %w", err)
	}

	list, err := s.serverGateway.GetCardList(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get card list: %w", err)
	}

	var cardDataList = make([]ListItem, 0, len(list))
	for _, item := range list {
		cardDataList = append(cardDataList, ListItem{
			ID:        item.ID,
			Title:     item.Title,
			CreatedAt: time.Unix(0, item.CreatedAt),
			UpdatedAt: time.Unix(0, item.UpdatedAt),
		})
	}

	return cardDataList, nil
}

func (s *service) GetCardByID(ctx context.Context, id int64) (*CardInfo, error) {
	userID, err := s.authService.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get user ID: %w", err)
	}

	card, err := s.serverGateway.GetCardByID(ctx, userID, id)
	if err != nil {
		return nil, fmt.Errorf("get card by ID: %w", err)
	}

	return &CardInfo{
		ID: card.ID,
		CardData: CardData{
			Title:  card.Title,
			Pan:    card.Pan,
			Cvv:    card.Cvv,
			Expiry: card.Expiry,
			Note:   card.Note,
		},
		CreatedAt: card.CreatedAt,
		UpdatedAt: card.UpdatedAt,
	}, nil
}
