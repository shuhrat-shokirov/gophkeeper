//nolint:wrapcheck,gocritic
package session

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(ctx context.Context, session *Session) error {
	args := m.Called(ctx, session)
	if err := args.Error(0); err != nil {
		return err
	}
	return nil
}

func (m *MockRepository) Get(ctx context.Context, refreshToken string) (*Session, error) {
	args := m.Called(ctx, refreshToken)
	if session, ok := args.Get(0).(*Session); ok {
		return session, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockRepository) Delete(ctx context.Context, token string) error {
	args := m.Called(ctx, token)
	if err := args.Error(0); err != nil {
		return err
	}
	return nil
}

var _ Repo = (*MockRepository)(nil)

func NewMockRepository() *MockRepository {
	return &MockRepository{}
}
