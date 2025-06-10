package data

import (
	"context"
	"fmt"
	"time"

	"gophkeeper/internal/client/gateways/server"
)

func (s *service) SaveLogin(ctx context.Context, data *LoginData) error {
	userID, err := s.authService.GetUserID(ctx)
	if err != nil {
		return fmt.Errorf("get user id: %w", err)
	}

	err = s.serverGateway.SaveLoginAndPass(ctx, userID, &server.LoginAndPass{
		Login:      data.Login,
		Pass:       data.Pass,
		Title:      data.Title,
		Note:       data.Note,
		ModifiedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("save login: %w", err)
	}

	return nil
}
