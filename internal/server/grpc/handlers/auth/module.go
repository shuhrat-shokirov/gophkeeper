package auth

import (
	"context"

	"go.uber.org/fx"
	"google.golang.org/grpc"

	"gophkeeper/internal/server/services/auth"
	pb "gophkeeper/proto"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In

	AuthService auth.Service
}

type handler struct {
	pb.UnimplementedAuthServiceServer

	authService auth.Service
}

type Handler interface {
	Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error)
	ConfirmOTP(ctx context.Context,
		request *pb.ConfirmOTPRequest) (*pb.ConfirmOTPResponse, error)

	RefreshToken(ctx context.Context, request *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error)

	Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error)
	Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error)

	RegisterService(srv *grpc.Server)
}

func New(p Params) Handler {
	return &handler{
		authService: p.AuthService,
	}
}

func (h *handler) RegisterService(srv *grpc.Server) {
	pb.RegisterAuthServiceServer(srv, h)
}
