//nolint:wrapcheck,gocritic
package logins

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) Save(ctx context.Context, login *LoginData) (int, error) {
	args := m.Called(ctx, login)
	if id, ok := args.Get(0).(int); ok {
		return id, args.Error(1)
	}
	return 0, args.Error(1)
}

func (m *MockRepo) List(ctx context.Context, userID int64, pg Pagination) ([]List, error) {
	args := m.Called(ctx, userID, pg)
	if list, ok := args.Get(0).([]List); ok {
		return list, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockRepo) GetByID(ctx context.Context, id int64) (*Info, error) {
	args := m.Called(ctx, id)
	if info, ok := args.Get(0).(*Info); ok {
		return info, args.Error(1)
	}
	return nil, args.Error(1)
}

var _ Repo = (*MockRepo)(nil)
