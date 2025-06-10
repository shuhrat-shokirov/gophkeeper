package auth

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/fx"

	"gophkeeper/internal/client/gateways/server"
	"gophkeeper/pkg/memorycache"
)

var Module = fx.Provide(New)

type Service interface {
	Register(ctx context.Context, email, password string) error
	ConfirmOTP(ctx context.Context, code string) error
	CheckAuth(ctx context.Context) error

	GetUserID(_ context.Context) (int64, error)

	Login(ctx context.Context, email, password string) error
	Logout(ctx context.Context)
}

type Params struct {
	fx.In
	fx.Lifecycle

	ServerGateway server.Gateway
	Cache         memorycache.Cache
}

type service struct {
	serverGateway server.Gateway
	cache         memorycache.Cache

	accessToken  string
	refreshToken string
	userID       int
	publicKey    []byte
}

func New(p Params) (Service, error) {

	_ = godotenv.Load(".env")

	publicKey := os.Getenv("GOPH_KEEPER_PUBLIC_KEY")

	bytes, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("decode public key: %w", err)
	}

	s := &service{
		serverGateway: p.ServerGateway,
		cache:         p.Cache,
		publicKey:     bytes,
		accessToken:   os.Getenv("GOPH_KEEPER_ACCESS_TOKEN"),
		refreshToken:  os.Getenv("GOPH_KEEPER_REFRESH_TOKEN"),
	}

	p.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			timeout := 10 * time.Second
			newCtx, cancelFunc := context.WithTimeout(context.Background(), timeout)
			defer cancelFunc()

			_ = s.CheckAuth(newCtx)

			go func() {
				ticker := time.NewTicker(5 * time.Minute)
				defer ticker.Stop()

				for {
					select {
					case <-ticker.C:
						if err := s.CheckAuth(ctx); err != nil {
							log.Printf("Error checking auth: %v", err)
							continue
						}
					case <-ctx.Done():
						return
					}
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping auth service...")
			err := writeEnvFile(s.accessToken, s.refreshToken)
			if err != nil {
				log.Printf("Error writing tokens to .env file: %v", err)
			}
			return nil
		},
	})

	return s, nil
}
