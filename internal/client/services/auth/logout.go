package auth

import (
	"context"
)

func (s *service) Logout(ctx context.Context) {
	if s.refreshToken == "" {
		return // No tokens to invalidate
	}

	s.serverGateway.Logout(ctx, s.refreshToken)

	s.accessToken = ""
	s.refreshToken = ""
}
