package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"go-west/internal/conf"
	"go-west/internal/server"
	"go-west/internal/service"
	"go-west/pkg/boot"
	log2 "go-west/pkg/logger2"
	"gopkg.in/yaml.v3"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "go-west"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../configs", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	//logConfig := logger.LoadLogConfig()
	//lo0 := logger.New2(logConfig)
	//lo01 := log.With(lo0,
	//	"service.name", Name,
	//	"service.version", Version,
	//	"ts", log.DefaultTimestamp,
	//	"caller", log.DefaultCaller,
	//)
	lo1 := log2.NewZapLogger()
	lo11 := log.With(lo1,
		"service.name", Name,
		"service.version", Version,
		//"ts", log.DefaultTimestamp,
		//"caller", log.DefaultCaller,
	)

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

	svc := service.New(lo11)
	httpSrv := server.NewHTTPServer(bc.Server, svc, lo11)
	grpcSrv := server.NewGRPCServer(bc.Server, svc, lo11)

	app := boot.New(
		boot.Server(
			httpSrv,
			grpcSrv,
		))
	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
