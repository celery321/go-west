package service

import (
	v1 "go-west/api/v1"
	"go-west/internal/biz"
	"go-west/pkg/log"
)



type Service struct {
	v1.UnimplementedGreeterServer
	uc  *biz.GreeterUsecase
	log *log.Helper
}

// New NewGreeterService new a greeter service.
func New(logger log.Logger) *Service {
	return &Service{
		log: log.NewHelper(log.With(logger, "module", "service/go-west")),
	}
}
