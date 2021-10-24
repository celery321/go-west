package sky

import (
	"context"
	"fmt"
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"strconv"
	"time"
	agentv3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
)
// Option is tracing option.
type Option func(*options)

type options struct {
	tracerProvider trace.TracerProvider
	propagator     propagation.TextMapPropagator
}

func getOperationName(ht http.Transporter) string {
	return fmt.Sprintf("%s", ht.PathTemplate())
}

//func getOperationName(ht http.Transporter) string {
//	return fmt.Sprintf("/%s%s", ht.Request().Method, ht.Request().URL)
//}



// WithTracerProvider with tracer provider.
// Deprecated: use otel.SetTracerProvider(provider) instead.
func WithTracerProvider(provider trace.TracerProvider) Option {
	return func(opts *options) {
		opts.tracerProvider = provider
	}
}


// Server returns a new server middleware for OpenTelemetry.
func Server(opts ...Option) middleware.Middleware {
	rp, err := reporter.NewGRPCReporter("114.67.201.131:11800", reporter.WithCheckInterval(time.Second))
	if err != nil{
		fmt.Printf("err===%v\n", err)
	}

	tracer, err := go2sky.NewTracer("go-west", go2sky.WithReporter(rp))
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
					if ht, ok := tr.(http.Transporter); ok {
						if tr.Kind() == transport.KindHTTP {
							span, _, _ := tracer.CreateEntrySpan(ctx, getOperationName(ht), func(key string) (string, error) {
								return tr.RequestHeader().Get(key), nil
							})
						fmt.Printf("endpint=[%v] oper=[%v] kind=[%v] request=[%v] requesheader=[%v] respone=[%v] temp=[%v]\n", ht.Endpoint(), ht.Operation(), ht.Kind(), ht.Request(), ht.RequestHeader(), ht.ReplyHeader() , ht.PathTemplate())
						span.SetComponent(5006)
						span.Tag(go2sky.TagHTTPMethod, ht.Request().Method)
						span.Tag(go2sky.TagURL, fmt.Sprintf("/%s%s", ht.Request().Method, ht.Request().URL))
						span.SetSpanLayer(agentv3.SpanLayer_Http)
						span.Tag(go2sky.TagStatusCode, strconv.Itoa(200))
					  	span.End()
					}
				}
			}
			return handler(ctx, req)
		}
	}
}
