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

type MockGateway struct {
	mock.Mock
}

func (m *MockGateway) Register(ctx context.Context, email, password string) (string, error) {
	args := m.Called(ctx, email, password)
	if otpId, ok := args.Get(0).(string); ok {
		return otpId, args.Error(1)
	}
	return "", args.Error(1)
}

func (m *MockGateway) ConfirmOtp(ctx context.Context, otpId, otpCode string) (*Token, error) {
	args := m.Called(ctx, otpId, otpCode)
	if token, ok := args.Get(0).(*Token); ok {
		return token, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockGateway) RefreshToken(ctx context.Context, refreshToken string) (*Token, error) {
	args := m.Called(ctx, refreshToken)
	if token, ok := args.Get(0).(*Token); ok {
		return token, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockGateway) Login(ctx context.Context, email, password string) (string, error) {
	args := m.Called(ctx, email, password)
	if accessToken, ok := args.Get(0).(string); ok {
		return accessToken, args.Error(1)
	}
	return "", args.Error(1)
}

func (m *MockGateway) Logout(ctx context.Context, refreshToken string) {
	args := m.Called(ctx, refreshToken)
	if err := args.Error(0); err != nil {
		panic(err) // In a real implementation, you would handle this error properly
	}
	// No return value expected for logout
}

func (m *MockGateway) SaveLoginAndPass(ctx context.Context, userID int64, pass *LoginAndPass) error {
	args := m.Called(ctx, userID, pass)
	if err := args.Error(0); err != nil {
		return err
	}
	// No return value expected for save operation
	return nil
}

func (m *MockGateway) SaveText(ctx context.Context, userID int64, data *Text) error {
	args := m.Called(ctx, userID, data)
	if err := args.Error(0); err != nil {
		return err
	}
	// No return value expected for save operation
	return nil
}

func (m *MockGateway) SaveCard(ctx context.Context, userID int64, data *Card) error {
	args := m.Called(ctx, userID, data)
	if err := args.Error(0); err != nil {
		return err
	}
	// No return value expected for save operation
	return nil
}

func (m *MockGateway) SaveBinary(ctx context.Context, userID int64, data *Binary) error {
	args := m.Called(ctx, userID, data)
	if err := args.Error(0); err != nil {
		return err
	}
	// No return value expected for save operation
	return nil
}

func (m *MockGateway) GetLoginList(ctx context.Context, userID int64, limit, offset int64) ([]ListItem, error) {
	args := m.Called(ctx, userID, limit, offset)
	if items, ok := args.Get(0).([]ListItem); ok {
		return items, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockGateway) GetLoginByID(ctx context.Context, userID, id int64) (*LoginInfo, error) {
	args := m.Called(ctx, userID, id)
	if loginInfo, ok := args.Get(0).(*LoginInfo); ok {
		return loginInfo, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockGateway) GetCardList(ctx context.Context, userID int64, limit, offset int64) ([]ListItem, error) {
	args := m.Called(ctx, userID, limit, offset)
	if items, ok := args.Get(0).([]ListItem); ok {
		return items, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockGateway) GetCardByID(ctx context.Context, userID, id int64) (*CardInfo, error) {
	args := m.Called(ctx, userID, id)
	if cardInfo, ok := args.Get(0).(*CardInfo); ok {
		return cardInfo, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockGateway) GetTextList(ctx context.Context, userID int64, limit, offset int64) ([]ListItem, error) {
	args := m.Called(ctx, userID, limit, offset)
	if items, ok := args.Get(0).([]ListItem); ok {
		return items, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockGateway) GetTextByID(ctx context.Context, userID, id int64) (*TextInfo, error) {
	args := m.Called(ctx, userID, id)
	if textInfo, ok := args.Get(0).(*TextInfo); ok {
		return textInfo, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockGateway) GetBinaryList(ctx context.Context, userID int64, limit, offset int64) ([]ListItem, error) {
	args := m.Called(ctx, userID, limit, offset)
	if items, ok := args.Get(0).([]ListItem); ok {
		return items, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockGateway) GetBinaryByID(ctx context.Context, userID, id int64) (*BinaryInfo, error) {
	args := m.Called(ctx, userID, id)
	if binaryInfo, ok := args.Get(0).(*BinaryInfo); ok {
		return binaryInfo, args.Error(1)
	}
	return nil, args.Error(1)
}

var _ Gateway = (*MockGateway)(nil)

func NewMockGateway() *MockGateway {
	return &MockGateway{}
}
