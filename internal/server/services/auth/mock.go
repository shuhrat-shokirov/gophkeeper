//nolint:wrapcheck,gocritic
package auth

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) Registration(ctx context.Context, request Registration) (string, error) {
	args := m.Called(ctx, request)
	if token, ok := args.Get(0).(string); ok {
		return token, args.Error(1)
	}
	return "", args.Error(1)
}

func (m *MockService) ConfirmOTP(ctx context.Context, id, code string) (*ConfirmResponse, error) {
	args := m.Called(ctx, id, code)
	if response, ok := args.Get(0).(*ConfirmResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockService) Login(ctx context.Context, email, password string) (string, error) {
	args := m.Called(ctx, email, password)
	if token, ok := args.Get(0).(string); ok {
		return token, args.Error(1)
	}
	return "", args.Error(1)
}

func (m *MockService) Logout(ctx context.Context, accessToken string) error {
	args := m.Called(ctx, accessToken)
	return args.Error(0)
}

func (m *MockService) RefreshToken(ctx context.Context, token string) (string, error) {
	args := m.Called(ctx, token)
	if newToken, ok := args.Get(0).(string); ok {
		return newToken, args.Error(1)
	}
	return "", args.Error(1)
}

var _ Service = (*MockService)(nil)

func NewMockService() *MockService {
	return &MockService{}
}
