//nolint:wrapcheck,gocritic
package user

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	args := m.Called(ctx, email)
	if user, ok := args.Get(0).(*User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockRepository) GetUserByID(ctx context.Context, id int) (*User, error) {
	args := m.Called(ctx, id)
	if user, ok := args.Get(0).(*User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockRepository) CreateUser(ctx context.Context, user *User) (int, error) {
	args := m.Called(ctx, user)
	if id, ok := args.Get(0).(int); ok {
		return id, args.Error(1)
	}
	return 0, args.Error(1)
}

var _ Repo = (*MockRepository)(nil)

func NewMockRepository() *MockRepository {
	return &MockRepository{}
}
