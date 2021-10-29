package biz

import (
	"context"
	"go-west/pkg/log"
)

type Greeter struct {
	Hello string
}

type GreeterRepo interface {
	GetGreeter(context.Context, *Greeter) ([]*Greeter,error)
	SetGreeter(context.Context, *Greeter)  error
}

type GreeterUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *GreeterUsecase) Create(ctx context.Context, g *Greeter)  ([]*Greeter,error) {
	return uc.repo.GetGreeter(ctx, g)
}

func (uc *GreeterUsecase) Set(ctx context.Context, g *Greeter)  error{
	return uc.repo.SetGreeter(ctx, g)
}
