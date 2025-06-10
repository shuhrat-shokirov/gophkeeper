//nolint:wrapcheck,gocritic
package texts

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) Save(ctx context.Context, data *TextData) (int64, error) {
	args := m.Called(ctx, data)
	if id, ok := args.Get(0).(int64); ok {
		return id, args.Error(1)
	}
	return 0, args.Error(1)
}

var _ Repo = (*MockRepo)(nil)

func NewMockRepo() *MockRepo {
	return &MockRepo{}
}
