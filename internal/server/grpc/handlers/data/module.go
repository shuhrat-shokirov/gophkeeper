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
	GetLoginList(ctx context.Context, request *pb.ListRequest) (*pb.ListResponse, error)
	GetLoginByID(ctx context.Context, request *pb.IDRequest) (*pb.LoginDataResponse, error)

	SaveText(ctx context.Context, request *pb.TextData) (*pb.Response, error)

	SaveCard(ctx context.Context, data *pb.CardData) (*pb.Response, error)

	SaveBinary(ctx context.Context, data *pb.BinaryData) (*pb.Response, error)

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
