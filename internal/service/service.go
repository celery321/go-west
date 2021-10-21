package service

import (
	v1 "go-west/api/v1"
	"go-west/internal/biz"
	"go-west/pkg/logger"
)



type Service struct {
	v1.UnimplementedGreeterServer
	uc  *biz.GreeterUsecase
	log logger.Logger
}

// New NewGreeterService new a greeter service.
func New(logger logger.Logger) *Service {
	return &Service{log: logger}
}
