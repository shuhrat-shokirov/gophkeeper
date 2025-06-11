package redis

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockRedis struct {
	mock.Mock
}

func (m *MockRedis) Save(ctx context.Context, key string, value any, dur time.Duration) error {
	args := m.Called(ctx, key, value, dur)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (m *MockRedis) Find(ctx context.Context, key string, value any) error {
	args := m.Called(ctx, key, value)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	if v, ok := args.Get(1).(error); ok && v != nil {
		return v
	}
	return nil
}

func (m *MockRedis) Delete(ctx context.Context, key string) (err error) {
	args := m.Called(ctx, key)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

var _ Cache = (*MockRedis)(nil)

func NewMockRedis() *MockRedis {
	return &MockRedis{}
}
