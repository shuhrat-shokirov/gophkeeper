//nolint:wrapcheck,gocritic
package data

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) SaveLogin(ctx context.Context, data *LoginData) error {
	args := m.Called(ctx, data)
	if err := args.Error(0); err != nil {
		return err
	}
	return nil
}

func (m *MockService) SaveText(ctx context.Context, data *TextData) error {
	args := m.Called(ctx, data)
	if err := args.Error(0); err != nil {
		return err
	}
	return nil
}

func (m *MockService) SaveCard(ctx context.Context, data *CardData) error {
	args := m.Called(ctx, data)
	if err := args.Error(0); err != nil {
		return err
	}
	return nil
}

func (m *MockService) SaveBinary(ctx context.Context, data *BinaryData) error {
	args := m.Called(ctx, data)
	if err := args.Error(0); err != nil {
		return err
	}
	return nil
}

func (m *MockService) GetLoginList(ctx context.Context, userID, limit, offset int64) ([]LoginListItem, error) {
	args := m.Called(ctx, userID, limit, offset)
	if items, ok := args.Get(0).([]LoginListItem); ok {
		return items, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockService) GetLoginByID(ctx context.Context, id int64) (*LoginInfo, error) {
	args := m.Called(ctx, id)
	if info, ok := args.Get(0).(*LoginInfo); ok {
		return info, args.Error(1)
	}
	return nil, args.Error(1)
}

var _ Service = (*MockService)(nil)

func NewMockService() *MockService {
	return &MockService{}
}
