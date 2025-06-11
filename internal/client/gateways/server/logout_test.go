package server

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

func Test_gateway_Logout(t *testing.T) {
	mockClient := new(MockAuthServiceClient)
	defer mockClient.AssertExpectations(t)

	mockClient.On("Logout", mock.Anything, mock.Anything).Return(nil, nil)

	g := &gateway{
		authServiceClient: mockClient,
	}

	g.Logout(t.Context(), "test-token")
}
