package service

import (
	"context"
	"github.com/SkyAPM/go2sky"
	v1 "go-west/api/v1"
	"go-west/internal/biz"
)

// GetHello SayHello implements helloworld.GreeterServer
func (s *Service) GetHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	if in.Name == "error" {
		s.log.Error("error")
		return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	}
	s.log.WithContext(ctx).Info("aaaaaaaaa")
	//go2sky.ctxKey
	if err := s.data.Ping(ctx); err != nil {
		s.log.WithContext(ctx).Infof("ping", err)
	}

	b := &biz.Greeter{
		Hello: in.Name,
	}
	l, err := s.uc.Create(ctx, b)
 	if err != nil {
		s.log.WithContext(ctx).Errorf("s.uc.Create", err)
		return &v1.HelloReply{}, err
	}
	for _, v := range l {
		s.log.WithContext(ctx).Info("l", v.Hello)
	}

	req := &v1.HelloReply{}
	req.ReqId = go2sky.TraceID(ctx)
	req.Message = "Hello" + in.GetName()


	// redis

	s.data.PingRedis(ctx)
	return req, nil
}

func (s *Service) SetHello(ctx context.Context, in *v1.HelloRequest)  (*v1.CreateClusterRes, error){
	req := &biz.Greeter{
		Hello: in.Name,
	}
	err := s.uc.Set(ctx, req)
	if err != nil {
		s.log.WithContext(ctx).Errorf("s.uc.Set", err)
		return &v1.CreateClusterRes{}, err
	}

	return &v1.CreateClusterRes{}, nil
}