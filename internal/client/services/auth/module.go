package auth

import (
	"context"
	"encoding/base64"
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

	Login(ctx context.Context, email, password string) error
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
	publicKey    []byte
}

func New(p Params) Service {

	_ = godotenv.Load(".env")

	//nolint:lll,gocritic
	publicKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUF6eUtPY2R1c0dJTy9VLzZnRTFsNAp0Z3Z4KzR5TUExQ1pQRkYwRnhKbzJaS1ZTY3I4SU41RVVDUDlUeEsxRTJLc2xnS01zQkhnYldJY3ZHMFBpTXZUClo0dTB3SWlPaTVSMDNlK3I5V1NqOG1xSCs3UjU1VndybUZVbFdMRWxuT1E4MnYveWNpV2hPZFJURWJ5cTZYQWcKaU5BckZyL3NFRTBacHFPdlVSeEFmeG5Qb1ZGd3M4NUplU0FYR1c2aG9HK3FoeEIvZ3diYkZOVitpbXViZUZ6dApyd1NJUGdmNjR2d2RoWnpDY1JZOVRUK1dyRm16Yk5uZmNwSzNvZEVnOCszdVJaWDdBb2R0U2E5OTZPTFFOcFNJClFnUDRiMnBXem5hc0NlRHU3ZlBWME5GNkJmSG1hcCs3RHZORTk1blcxUUh2MzZRRzNoVVRmZ0ZZQVJoc0NrOTAKcndJREFRQUIKLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg=="
	bytes, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		log.Fatalf("failed to decode public key: %v", err)
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

	return s
}
