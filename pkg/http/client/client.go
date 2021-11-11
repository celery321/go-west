package client

import (
	"context"
	"fmt"
	hp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/pkg/errors"
	"go-west/pkg/http/middleware/logging"
	"go-west/pkg/log"
)

func NewHttpClient(logger log.Logger)  *hp.Client {
	conn := &hp.Client{}
	conn, err := hp.NewClient(
	 	context.Background(),
		hp.WithMiddleware(
			logging.Client(logger),
		),
		hp.WithEndpoint("metamon-api.radiocaca.com:80"),
	)

	if err != nil {
		fmt.Printf("error=%v\n", errors.Wrap(err, "NewClient"))
	}

	return conn

}

//// Client is http client.
//type Client struct {
//	conf      *ClientConfig
//	client    *xhttp.Client
//	dialer    *net.Dialer
//	transport xhttp.RoundTripper
//
//	urlConf  map[string]*ClientConfig
//	hostConf map[string]*ClientConfig
//	mutex    sync.RWMutex
//	breaker  *breaker.Group
//}
//
//// Raw sends an HTTP request and returns bytes response
//func (client *Client) Raw(c context.Context, req *xhttp.Request, v ...string) (bs []byte, err error) {
//	var (
//		ok      bool
//		code    string
//		cancel  func()
//		resp    *xhttp.Response
//		config  *ClientConfig
//		timeout time.Duration
//		uri     = fmt.Sprintf("%s://%s%s", req.URL.Scheme, req.Host, req.URL.Path)
//	)
//	// NOTE fix prom & config uri key.
//	if len(v) == 1 {
//		uri = v[0]
//	}
//	// breaker
//	brk := client.breaker.Get(uri)
//	if err = brk.Allow(); err != nil {
//		code = "breaker"
//		clientStats.Incr(uri, code)
//		return
//	}
//	defer client.onBreaker(brk, &err)
//	// stat
//	now := time.Now()
//	defer func() {
//		clientStats.Timing(uri, int64(time.Since(now)/time.Millisecond))
//		if code != "" {
//			clientStats.Incr(uri, code)
//		}
//	}()
//	// get config
//	// 1.url config 2.host config 3.default
//	client.mutex.RLock()
//	if config, ok = client.urlConf[uri]; !ok {
//		if config, ok = client.hostConf[req.Host]; !ok {
//			config = client.conf
//		}
//	}
//	client.mutex.RUnlock()
//	// timeout
//	deliver := true
//	timeout = time.Duration(config.Timeout)
//	if deadline, ok := c.Deadline(); ok {
//		if ctimeout := time.Until(deadline); ctimeout < timeout {
//			// deliver small timeout
//			timeout = ctimeout
//			deliver = false
//		}
//	}
//	if deliver {
//		c, cancel = context.WithTimeout(c, timeout)
//		defer cancel()
//	}
//	setTimeout(req, timeout)
//	req = req.WithContext(c)
//	setCaller(req)
//	if color := metadata.String(c, metadata.Color); color != "" {
//		setColor(req, color)
//	}
//	if resp, err = client.client.Do(req); err != nil {
//		err = pkgerr.Wrapf(err, "host:%s, url:%s", req.URL.Host, realURL(req))
//		code = "failed"
//		return
//	}
//	defer resp.Body.Close()
//	if resp.StatusCode >= xhttp.StatusBadRequest {
//		err = pkgerr.Errorf("incorrect http status:%d host:%s, url:%s", resp.StatusCode, req.URL.Host, realURL(req))
//		code = strconv.Itoa(resp.StatusCode)
//		return
//	}
//	if bs, err = readAll(resp.Body, _minRead); err != nil {
//		err = pkgerr.Wrapf(err, "host:%s, url:%s", req.URL.Host, realURL(req))
//		return
//	}
//	return
//}
//
//
//// Do sends an HTTP request and returns an HTTP json response.
//func (client *Client) Do(c context.Context, req *xhttp.Request, res interface{}, v ...string) (err error) {
//	var bs []byte
//	if bs, err = client.Raw(c, req, v...); err != nil {
//		return
//	}
//	if res != nil {
//		if err = json.Unmarshal(bs, res); err != nil {
//			err = pkgerr.Wrapf(err, "host:%s, url:%s", req.URL.Host, realURL(req))
//		}
//	}
//	return
//}