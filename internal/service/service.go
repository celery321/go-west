package service

import (
	"github.com/go-kratos/kratos/v2/log"
	v1 "go-west/api/v1"
	"go-west/internal/biz"
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
