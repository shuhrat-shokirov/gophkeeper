//nolint:wrapcheck,gocritic
package auth

import (
	"context"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	pb "gophkeeper/proto"
)

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	args := m.Called(ctx, request)
	if response, ok := args.Get(0).(*pb.RegisterResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockHandler) ConfirmOTP(ctx context.Context,
	request *pb.ConfirmOTPRequest) (*pb.ConfirmOTPResponse, error) {
	args := m.Called(ctx, request)
	if response, ok := args.Get(0).(*pb.ConfirmOTPResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockHandler) RefreshToken(ctx context.Context,
	request *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	args := m.Called(ctx, request)
	if response, ok := args.Get(0).(*pb.RefreshTokenResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockHandler) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	args := m.Called(ctx, request)
	if response, ok := args.Get(0).(*pb.LoginResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockHandler) RegisterService(srv *grpc.Server) {
	args := m.Called(srv)
	if len(args) > 0 {
		if fn, ok := args.Get(0).(func(*grpc.Server)); ok {
			fn(srv)
		}
	}
}

var _ Handler = (*MockHandler)(nil)

func NewMockHandler() *MockHandler {
	return &MockHandler{}
}
