package data

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestNew(t *testing.T) {
	h := New(Params{})
	require.NotNil(t, h)
}

func TestRegisterService(t *testing.T) {
	grpcServer := grpc.NewServer()

	h := New(Params{})
	require.NotNil(t, h)

	h.RegisterService(grpcServer)
}
