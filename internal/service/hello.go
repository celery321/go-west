package service

import (
	"context"
	"fmt"
	"github.com/SkyAPM/go2sky"
	"github.com/go-kratos/kratos/v2/metadata"
	v1 "go-west/api/v1"
	skywalk "go-west/pkg/http/middleware/skywalking"
)

// SayHello implements helloworld.GreeterServer
func (s *Service) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	if in.Name == "error" {
		s.log.Error("error")
		return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	}
	s.log.WithContext(ctx).Info("aaaaaaaaa")
	//go2sky.ctxKey
	fmt.Printf("go2==%v\n", go2sky.TraceID(ctx))

	md, _ := metadata.FromServerContext(ctx)
	fmt.Printf("md===%v\n",  md)

	req := &v1.HelloReply{}
	req.ReqId = go2sky.TraceID(ctx)
	fmt.Printf("TraceID=%v real=%v\n", skywalk.TraceID(),go2sky.TraceID(ctx) )
	req.Message = "Hello" + in.GetName()

	return req, nil
}

