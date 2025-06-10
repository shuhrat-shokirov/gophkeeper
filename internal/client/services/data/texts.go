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
		Title:     data.Title,
		Content:   data.Content,
		Note:      data.Note,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("save text: %w", err)
	}

	return nil
}

func (s *service) GetTextList(ctx context.Context, limit, offset int64) ([]ListItem, error) {
	userID, err := s.authService.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get user id: %w", err)
	}

	list, err := s.serverGateway.GetTextList(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get list text: %w", err)
	}

	var textDataList = make([]ListItem, 0, len(list))
	for _, item := range list {
		textDataList = append(textDataList, ListItem{
			ID:        item.ID,
			Title:     item.Title,
			CreatedAt: time.Unix(0, item.CreatedAt),
			UpdatedAt: time.Unix(0, item.UpdatedAt),
		})
	}

	return textDataList, nil
}

func (s *service) GetTextByID(ctx context.Context, id int64) (*TextInfo, error) {
	userID, err := s.authService.GetUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get user id: %w", err)
	}

	text, err := s.serverGateway.GetTextByID(ctx, userID, id)
	if err != nil {
		return nil, fmt.Errorf("get text by ID: %w", err)
	}

	return &TextInfo{
		ID: text.ID,
		TextData: TextData{
			Title:   text.Title,
			Content: text.Content,
			Note:    text.Note,
		},
		CreatedAt: text.CreatedAt,
		UpdatedAt: text.UpdatedAt,
	}, nil
}
