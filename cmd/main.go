package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2/log"
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
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func main() {
	log := log.NewHelper(log.DefaultLogger)
	log.Info("11111111111111111111111")


}
