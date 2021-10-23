package service

import (
	"context"
	v1 "go-west/api/v1"
)

// SayHello implements helloworld.GreeterServer
func (s *Service) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	if in.Name == "error" {
		s.log.Error("error")
		return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	}
	s.log.Info("info")
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}

