package auth

import (
	"context"

	"go.uber.org/fx"

	"gophkeeper/internal/server/gateways/emailtotp"
	"gophkeeper/internal/server/repositories/session"
	"gophkeeper/internal/server/repositories/user"
	"gophkeeper/pkg/jwt"
	"gophkeeper/pkg/redis"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In

	Cache            redis.Cache
	EmailTotpGateway emailtotp.Gateway
	JWT              jwt.JWT

	UserRepo    user.Repo
	SessionRepo session.Repo
}

type Service interface {
	Registration(ctx context.Context, request Registration) (string, error)
	ConfirmOTP(ctx context.Context, id, code string) (*ConfirmResponse, error)

	RefreshToken(ctx context.Context, token string) (string, error)
}

type service struct {
	cache            redis.Cache
	emailTotpGateway emailtotp.Gateway
	jwt              jwt.JWT

	userRepo    user.Repo
	sessionRepo session.Repo
}

func New(p Params) Service { //nolint:gocritic
	return &service{
		cache:            p.Cache,
		emailTotpGateway: p.EmailTotpGateway,
		jwt:              p.JWT,

		userRepo:    p.UserRepo,
		sessionRepo: p.SessionRepo,
	}
}
