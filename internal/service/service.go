package service

import (
	"context"
	v1 "gowest/api/v1"
	"gowest/internal/biz"
	"gowest/pkg/logger"
)

type Service struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase

	log logger.Logger
}

// New NewGreeterService new a greeter service.
func New(logger logger.Logger) *Service {
	return &Service{log: logger}
}

// SayHello implements helloworld.GreeterServer
func (s *Service) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.Infof("SayHello")
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}
