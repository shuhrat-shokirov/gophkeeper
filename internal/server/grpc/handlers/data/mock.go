//nolint:wrapcheck,gocritic
package data

import (
	"context"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	pb "gophkeeper/proto"
)

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) SaveLogin(context context.Context, request *pb.LoginData) (*pb.Response, error) {
	args := m.Called(context, request)
	if response, ok := args.Get(0).(*pb.Response); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockHandler) RegisterService(srv *grpc.Server) {
	args := m.Called(srv)
	if len(args) > 0 {
		if fn, ok := args.Get(0).(func(*grpc.Server)); ok {
			fn(srv)
		}
	}
}

var _ Handler = (*MockHandler)(nil)

func NewMockHandler() *MockHandler {
	return &MockHandler{}
}
