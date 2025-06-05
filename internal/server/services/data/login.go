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
