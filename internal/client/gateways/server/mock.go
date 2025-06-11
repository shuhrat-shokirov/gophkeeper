//nolint:wrapcheck,gocritic,lll
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

func (m *MockAuthServiceClient) Register(ctx context.Context, in *pb.RegisterRequest, opts ...grpc.CallOption) (*pb.RegisterResponse, error) {
	args := m.Called(ctx, in)
	if response, ok := args.Get(0).(*pb.RegisterResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAuthServiceClient) ConfirmOTP(ctx context.Context, in *pb.ConfirmOTPRequest, opts ...grpc.CallOption) (*pb.ConfirmOTPResponse, error) {
	args := m.Called(ctx, in)
	if response, ok := args.Get(0).(*pb.ConfirmOTPResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAuthServiceClient) Login(ctx context.Context, in *pb.LoginRequest, opts ...grpc.CallOption) (*pb.LoginResponse, error) {
	args := m.Called(ctx, in)
	if response, ok := args.Get(0).(*pb.LoginResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAuthServiceClient) Logout(ctx context.Context, in *pb.LogoutRequest, opts ...grpc.CallOption) (*pb.LogoutResponse, error) {
	args := m.Called(ctx, in)
	if response, ok := args.Get(0).(*pb.LogoutResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAuthServiceClient) RefreshToken(ctx context.Context, in *pb.RefreshTokenRequest, opts ...grpc.CallOption) (*pb.RefreshTokenResponse, error) {
	args := m.Called(ctx, in)
	if response, ok := args.Get(0).(*pb.RefreshTokenResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

type MockDataServiceClient struct {
	mock.Mock
}

func (m *MockDataServiceClient) SaveLogin(ctx context.Context, in *pb.LoginData, opts ...grpc.CallOption) (*pb.Response, error) {
	args := m.Called(ctx, in)
	if response, ok := args.Get(0).(*pb.Response); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDataServiceClient) SaveText(ctx context.Context, in *pb.TextData, opts ...grpc.CallOption) (*pb.Response, error) {
	args := m.Called(ctx, in)
	if response, ok := args.Get(0).(*pb.Response); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDataServiceClient) SaveBinary(ctx context.Context, in *pb.BinaryData, opts ...grpc.CallOption) (*pb.Response, error) {
	args := m.Called(ctx, in)
	if response, ok := args.Get(0).(*pb.Response); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDataServiceClient) SaveCard(ctx context.Context, in *pb.CardData, opts ...grpc.CallOption) (*pb.Response, error) {
	args := m.Called(ctx, in)
	if response, ok := args.Get(0).(*pb.Response); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDataServiceClient) GetLoginList(ctx context.Context, in *pb.ListRequest, opts ...grpc.CallOption) (*pb.ListResponse, error) {
	args := m.Called(ctx, in)
	if response, ok := args.Get(0).(*pb.ListResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDataServiceClient) GetLoginByID(ctx context.Context, in *pb.IDRequest, opts ...grpc.CallOption) (*pb.LoginDataResponse, error) {
	args := m.Called(ctx, in)
	if response, ok := args.Get(0).(*pb.LoginDataResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDataServiceClient) GetTextList(ctx context.Context, in *pb.ListRequest, opts ...grpc.CallOption) (*pb.ListResponse, error) {
	args := m.Called(ctx, in)
	if response, ok := args.Get(0).(*pb.ListResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDataServiceClient) GetTextByID(ctx context.Context, in *pb.IDRequest, opts ...grpc.CallOption) (*pb.TextDataResponse, error) {
	args := m.Called(ctx, in)
	if response, ok := args.Get(0).(*pb.TextDataResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDataServiceClient) GetBinaryList(ctx context.Context, in *pb.ListRequest, opts ...grpc.CallOption) (*pb.ListResponse, error) {
	args := m.Called(ctx, in)
	if response, ok := args.Get(0).(*pb.ListResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDataServiceClient) GetBinaryByID(ctx context.Context, in *pb.IDRequest, opts ...grpc.CallOption) (*pb.BinaryDataResponse, error) {
	args := m.Called(ctx, in)
	if response, ok := args.Get(0).(*pb.BinaryDataResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDataServiceClient) GetCardList(ctx context.Context, in *pb.ListRequest, opts ...grpc.CallOption) (*pb.ListResponse, error) {
	args := m.Called(ctx, in)
	if response, ok := args.Get(0).(*pb.ListResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDataServiceClient) GetCardByID(ctx context.Context, in *pb.IDRequest, opts ...grpc.CallOption) (*pb.CardDataResponse, error) {
	args := m.Called(ctx, in)
	if response, ok := args.Get(0).(*pb.CardDataResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}
