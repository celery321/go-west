package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"gopkg.in/yaml.v3"
	"gowest/internal/conf"
	"gowest/internal/server"
	"gowest/internal/service"
	"gowest/pkg/logger"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "beer.cart.service"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../configs", "config path, eg: -conf config.yaml")
}

func main() {

	//logger := log.With(lo,)
	//logger.Log(log.LevelInfo, "msg", "111")

	logConfig := logger.LoadLogConfig()
	lo := logger.New(logConfig)
	//logger := log.With(log.NewStdLogger(os.Stdout),
	//	"service.name", Name,
	//	"service.version", Version,
	//	"ts", log.DefaultTimestamp,
	//	"caller", log.DefaultCaller,
	//)
	//l := log.NewHelper(log.With(logger, "module", "main"))
	//l.Info("aaaaaaaaaaaaaaaaaaaaa")

	// init conf
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	// 服务注册
	//var rc conf.Registry
	//if err := c.Scan(&rc); err != nil {
	//	panic(err)
	//}
	//

	svc := service.New(lo)
	httpSrv := server.NewHTTPServer(bc.Server, svc, lo)
	grpcSrv := server.NewGRPCServer(bc.Server, svc, lo)

	app := kratos.New(
		//kratos.Logger(lo),
		kratos.Server(
			httpSrv,
			grpcSrv,
		))
	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
