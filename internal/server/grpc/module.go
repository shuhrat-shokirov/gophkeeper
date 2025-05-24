package grpc

import "go.uber.org/fx"

var Module = fx.Invoke(New)

type Params struct {
	fx.In
}

func New(p Params) {
}
