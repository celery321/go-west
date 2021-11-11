package main

import (
	"context"
	"fmt"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/pkg/errors"
	pb "go-west/api/v1"
	"go-west/pkg/http/middleware/logging"
	"go-west/pkg/log"
	"time"
)


func main() {
	logger := log.DefaultLogger
	conn, err := transhttp.NewClient(
		context.Background(),
		transhttp.WithMiddleware(
			logging.Client(logger),
		),
		transhttp.WithTimeout(10 * time.Minute),
		transhttp.WithEndpoint("http://192.168.50.16:809"),
	)
	if err != nil {
		panic(errors.Wrap(err, "NewClient"))
	}

	client := pb.NewGreeterHTTPClient(conn)
	reply, err := client.GetHello(context.Background(), &pb.HelloRequest{
		Name: "demo",
	})
	if err != nil {
		fmt.Printf("err=%v", err)
	}
	fmt.Printf("[http] Hello [%v]\n", reply)


}
