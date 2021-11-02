package main

import (
	"context"
	"fmt"
	hp "github.com/go-kratos/kratos/v2/transport/http"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"go-west/pkg/http/middleware/logging"
	"go-west/pkg/log"
)

func main() {
	logger := log.DefaultLogger

	conn, err := hp.NewClient(
		context.Background(),
		hp.WithMiddleware(
			logging.Client(logger),
		),
		hp.WithEndpoint("127.0.0.1:8000"),
	)
	if err != nil {
		panic(errors.Wrap(err, "NewClient"))
	}
	//req := &http.Request{}
	req, err := http.NewRequest("GET", "http://192.168.50.16:809/hello/123452333", strings.NewReader(""))
	if err != nil {
		// handle error
	}

	response, err := conn.Do(req)
	if err != nil {
		fmt.Printf("%v\n", errors.Wrap(err,"conn.Do"))
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}
