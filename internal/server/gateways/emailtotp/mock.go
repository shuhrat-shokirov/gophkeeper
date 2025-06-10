//nolint:wrapcheck,gocritic
package emailtotp

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockGateway struct {
	mock.Mock
}

func (m *MockGateway) SendEmail(_ context.Context, request *Request) error {
	args := m.Called(request)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

var _ Gateway = (*MockGateway)(nil)

func NewMockGateway() *MockGateway {
	return &MockGateway{}
}
