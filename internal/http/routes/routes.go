package routes

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	"go-template/internal/http/handlers/health"
	"go-template/pkg/config"
)

var Module = fx.Invoke(New)

type Params struct {
	fx.In
	fx.Lifecycle

	HealthHandler health.Handler
	Config        config.Config
}

func New(p Params) {
	var (
		basePath = p.Config.GetString("server.basePath")
	)

	engine := gin.New()

	engine.GET(basePath+"/health", p.HealthHandler.Health)

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	server := &http.Server{
		Addr:              p.Config.GetString("server.port"),
		Handler:           engine,
		ReadHeaderTimeout: 5 * time.Second, // Устанавливаем таймаут для чтения заголовков
	}

	p.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					panic(err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := server.Shutdown(ctx); err != nil {
				return fmt.Errorf("server shutdown: %w", err)
			}

			return nil
		},
	})
}
