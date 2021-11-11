package service

import (
	v1 "go-west/api/v1"
	"go-west/internal/biz"
	"go-west/internal/conf"
	"go-west/internal/data"
	"go-west/pkg/log"
)



type Service struct {
	v1.UnimplementedGreeterServer
	log *log.Helper
	data *data.Data
	repo biz.GreeterRepo
	uc  *biz.GreeterUsecase
	//repo biz.BattleRepo
	//uc  *biz.BattleUsecase
}



// New NewGreeterService new a greeter service.
func New(c *conf.Data, logger log.Logger) *Service {
	s :=  &Service{
		log: log.NewHelper(log.With(logger, "module", "service/go-west")),
	}
	s.data, _, _  = data.NewData(c, logger)
	//// 实例化hello
	s.repo = data.NewGreeterRepo(s.data, logger)
	s.uc = biz.NewGreeterUsecase(s.repo, logger)
	// 实例化bettle
	//s.repo = data.NewBattleRepo(s.data, logger)
	//s.uc = biz.NewBattleUsecase(s.repo, logger)



	return s
}

