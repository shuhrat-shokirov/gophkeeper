package auth

import (
	"context"

	"go.uber.org/fx"
	"google.golang.org/grpc"

	"gophkeeper/internal/server/services/auth"
	"gophkeeper/pkg/logger"
	pb "gophkeeper/proto"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In

	Logger      logger.Logger
	AuthService auth.Service
}

type handler struct {
	pb.UnimplementedAuthServiceServer

	logger      logger.Logger
	authService auth.Service
}

type Handler interface {
	Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error)
	ConfirmRegistration(ctx context.Context,
		request *pb.ConfirmRegistrationRequest) (*pb.ConfirmRegistrationResponse, error)

	RefreshToken(ctx context.Context, request *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error)

	RegisterService(srv *grpc.Server)
}

func New(p Params) Handler {
	return &handler{
		logger:      p.Logger,
		authService: p.AuthService,
	}
}

func (h *handler) RegisterService(srv *grpc.Server) {
	pb.RegisterAuthServiceServer(srv, h)
}
