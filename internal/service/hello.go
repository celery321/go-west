package service

import (
	"context"
	v1 "go-west/api/v1"
)

// SayHello implements helloworld.GreeterServer
func (s *Service) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.Infof("SayHello1")
	s.log.WithFields(map[string]interface{}{"aaaa":"aaaaa"}).Infof("111")
	s.log.Infof("SayHello2")
	//err := errors.New(200, "10002", v1.ErrorReason_USER_NOT_FOUND.String())

	//
	//err = errors.New(200, "10002", v1.ErrorReason_USER_NOT_FOUND.String()).WithMetadata(map[string]string{
	//	"a":"1"})
	err := v1.ErrorContentMissing("11111")
	//err := nil
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, err
}

