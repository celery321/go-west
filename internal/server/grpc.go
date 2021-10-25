package server

import (
	"github.com/SkyAPM/go2sky"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "go-west/api/v1"
	"go-west/internal/conf"
	"go-west/internal/service"
	"go-west/pkg/http/middleware/logging"
	"go-west/pkg/http/middleware/metadata"
	skywalk "go-west/pkg/http/middleware/skywalking"
	"go-west/pkg/log"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.Service, logger log.Logger,tracer *go2sky.Tracer) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			metadata.Server(),
			skywalk.Server(tracer),
			logging.Server(logger),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServer(srv, greeter)
	return srv
}
