package data

import (
	"context"

	"go.uber.org/fx"
	"google.golang.org/grpc"

	"gophkeeper/internal/server/services/data"
	pb "gophkeeper/proto"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In

	DataService data.Service
}

type Handler interface {
	SaveLogin(context context.Context, request *pb.LoginData) (*pb.Response, error)

	RegisterService(srv *grpc.Server)
}

type handler struct {
	pb.UnimplementedDataServiceServer

	dataService data.Service
}

func New(p Params) Handler {
	return &handler{
		dataService: p.DataService,
	}
}

func (h *handler) RegisterService(srv *grpc.Server) {
	pb.RegisterDataServiceServer(srv, h)
}
