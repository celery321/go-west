package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "gowest/api/v1"
	"gowest/internal/biz"
)

type Service struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase

	log *log.Helper
}

// New NewGreeterService new a greeter service.
func New(logger log.Logger) *Service {
	return &Service{log: log.NewHelper(logger)}
}

// SayHello implements helloworld.GreeterServer
func (s *Service) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}
