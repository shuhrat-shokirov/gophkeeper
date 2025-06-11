//nolint:wrapcheck,gocritic
package auth

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) Register(ctx context.Context, email, password string) error {
	args := m.Called(ctx, email, password)
	if err := args.Error(0); err != nil {
		return err
	}
	return nil
}

func (m *MockService) ConfirmOTP(ctx context.Context, code string) error {
	args := m.Called(ctx, code)
	if err := args.Error(0); err != nil {
		return err
	}
	return nil
}

func (m *MockService) CheckAuth(ctx context.Context) error {
	args := m.Called(ctx)
	if err := args.Error(0); err != nil {
		return err
	}
	return nil
}

func (m *MockService) GetUserID(_ context.Context) (int64, error) {
	args := m.Called()
	if userID, ok := args.Get(0).(int64); ok {
		return userID, args.Error(1)
	}
	return 0, args.Error(1)
}

func (m *MockService) Login(ctx context.Context, email, password string) error {
	args := m.Called(ctx, email, password)
	if err := args.Error(0); err != nil {
		return err
	}
	return nil
}

func (m *MockService) Logout(ctx context.Context) {
	args := m.Called(ctx)
	if err := args.Error(0); err != nil {
		panic(err) // In a real implementation, you would handle this error properly
	}
}

var _ Service = (*MockService)(nil)

func NewMockService() *MockService {
	return &MockService{}
}
