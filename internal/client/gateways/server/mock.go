//nolint:wrapcheck,gocritic
package server

import (
	"context"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	pb "gophkeeper/proto"
)

type MockAuthServiceClient struct {
	mock.Mock
}

func (m *MockAuthServiceClient) Register(ctx context.Context, in *pb.RegisterRequest,
	opts ...grpc.CallOption) (*pb.RegisterResponse, error) {
	args := m.Called(ctx, in)
	if response, ok := args.Get(0).(*pb.RegisterResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAuthServiceClient) ConfirmOTP(ctx context.Context, in *pb.ConfirmOTPRequest,
	opts ...grpc.CallOption) (*pb.ConfirmOTPResponse, error) {
	args := m.Called(ctx, in)
	if response, ok := args.Get(0).(*pb.ConfirmOTPResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAuthServiceClient) Login(ctx context.Context, in *pb.LoginRequest,
	opts ...grpc.CallOption) (*pb.LoginResponse, error) {
	args := m.Called(ctx, in)
	if response, ok := args.Get(0).(*pb.LoginResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAuthServiceClient) Logout(ctx context.Context, in *pb.LogoutRequest,
	opts ...grpc.CallOption) (*pb.LogoutResponse, error) {
	args := m.Called(ctx, in)
	if response, ok := args.Get(0).(*pb.LogoutResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAuthServiceClient) RefreshToken(ctx context.Context, in *pb.RefreshTokenRequest,
	opts ...grpc.CallOption) (*pb.RefreshTokenResponse, error) {
	args := m.Called(ctx, in)
	if response, ok := args.Get(0).(*pb.RefreshTokenResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}
