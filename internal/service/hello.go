package service

import (
	"context"
	"github.com/SkyAPM/go2sky"
	v1 "go-west/api/v1"
	"go-west/internal/biz"
)

// SayHello implements helloworld.GreeterServer
func (s *Service) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	if in.Name == "error" {
		s.log.Error("error")
		return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	}
	s.log.WithContext(ctx).Info("aaaaaaaaa")
	//go2sky.ctxKey
	b := &biz.Greeter{
		Hello: "abc",
	}
	err := s.uc.Create(ctx, b)
 	if err != nil {

	}

	req := &v1.HelloReply{}
	req.ReqId = go2sky.TraceID(ctx)
	req.Message = "Hello" + in.GetName()
	return req, nil
}

