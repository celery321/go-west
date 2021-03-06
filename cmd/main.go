package main

import (
	"flag"
	"fmt"
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/gorilla/mux"
	"go-west/internal/conf"
	"go-west/internal/server"
	"go-west/internal/service"
	"go-west/pkg/boot"
	"go-west/pkg/log"
	"gopkg.in/yaml.v3"
	"net/http"
	"time"
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

	// init trace
	rp, err := reporter.NewGRPCReporter(bc.Trace.Endpoint, reporter.WithCheckInterval(time.Second))
	if err != nil{
		fmt.Printf("create gosky reporter failed!")
	}
	tracer, err := go2sky.NewTracer(Name, go2sky.WithReporter(rp))

	loggerInstance := log.NewZapLogger(bc.Server.Log)
	logger := log.With(loggerInstance,
		"service.name", Name,
		"service.version", Version,
		//"trace.id",   skywalk.TraceID() ,
	)
	router := mux.NewRouter()
	router.HandleFunc("/home", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hello Gorilla Mux!")
	}).Methods("GET")
	svc := service.New(bc.Data, logger)
	httpSrv := server.NewHTTPServer(bc.Server, svc, logger, tracer)
	grpcSrv := server.NewGRPCServer(bc.Server, svc, logger, tracer)

	app := boot.New(
		boot.Name(Name),
		boot.Version(Version),
		//boot.Metadata(map[string]string{"a":"1"}),
		boot.Logger(logger),
		boot.Server(
			httpSrv,
			grpcSrv,
		))
	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
