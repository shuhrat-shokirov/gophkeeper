package grpc

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"gophkeeper/internal/server/grpc/handlers"
	"gophkeeper/internal/server/grpc/handlers/auth"
	"gophkeeper/internal/server/grpc/handlers/data"
	"gophkeeper/pkg/config"
	"gophkeeper/pkg/logger"
)

var Module = fx.Options(
	handlers.Module,
	fx.Invoke(New),
)

type Params struct {
	fx.In
	fx.Lifecycle

	Config config.Config
	Logger logger.Logger

	AuthHandler auth.Handler
	DataHandler data.Handler
}

func New(p Params) error {
	server := grpc.NewServer()

	p.AuthHandler.RegisterService(server)
	p.DataHandler.RegisterService(server)

	reflection.Register(server)

	listen, err := net.Listen("tcp", p.Config.GetString("grpc.address"))
	if err != nil {
		p.Logger.Error("Failed to listen on address",
			zap.String("address", p.Config.GetString("grpc.address")),
			zap.Error(err))
		return fmt.Errorf("failed to listen %w", err)
	}

	p.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				p.Logger.Info("Starting gRPC server",
					zap.String("address", p.Config.GetString("grpc.address")))

				if err := server.Serve(listen); err != nil {
					p.Logger.Error("Failed to start gRPC server",
						zap.String("address", p.Config.GetString("grpc.address")),
						zap.Error(err))
					return
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.GracefulStop()
			if err := listen.Close(); err != nil {
				p.Logger.Error("Failed to close listener",
					zap.String("address", p.Config.GetString("grpc.address")),
					zap.Error(err))
				return fmt.Errorf("failed to close listener %w", err)
			}

			return nil
		},
	})

	return nil
}
