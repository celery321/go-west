package server

import (
	"github.com/SkyAPM/go2sky"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	v1 "go-west/api/v1"
	"go-west/internal/conf"
	"go-west/internal/service"
	"go-west/pkg/http/middleware/logging"
	"go-west/pkg/http/middleware/metadata"
	skywalk "go-west/pkg/http/middleware/skywalking"
	"go-west/pkg/http/middleware/validate"
	"go-west/pkg/log"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.Service, logger log.Logger, tracer *go2sky.Tracer) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			metadata.Server(),
			skywalk.Server(tracer),
			logging.Server(logger),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	// error
	//opts = append(opts, http.ErrorEncoder(response.ErrorEncoder))
	// response
	//opts = append(opts, http.ResponseEncoder(response.ResponseEncoder))

	srv := http.NewServer(opts...)
	srv.Handle("/metrics", promhttp.Handler())
	v1.RegisterGreeterHTTPServer(srv, greeter)

	return srv
}

