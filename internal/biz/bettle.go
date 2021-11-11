package biz

import (
	"context"
	"go-west/pkg/log"
)

type BattleRequest struct {
	MonsterA string
	MonsterB  string
	Address string
	BattleLevel uint32

}

type Data struct {
	BattleLevel uint32
	BpFragmentNum uint32
	BpPotionNum uint32
	ChallengeExp uint32
	ChallengeLevel uint32
}
type BattleReply struct {
	Code string
	Data Data
	ErrorText string
	Message string
	Result int32


}

type BattleUsecase struct {
	repo BattleRepo
	log  *log.Helper
}

type BattleRepo interface {
	NStart(context.Context, *BattleRequest) (*BattleReply,error)
	Start(context.Context, *BattleRequest) (*BattleReply,error)
}

func NewBattleUsecase(repo BattleRepo, logger log.Logger) *BattleUsecase {
	return &BattleUsecase{repo: repo, log: log.NewHelper(logger)}
}


func (uc *BattleUsecase) NStart(ctx context.Context, g *BattleRequest)  (*BattleReply,error) {
	return uc.repo.NStart(ctx, g)
}
func (uc *BattleUsecase) Start(ctx context.Context, g *BattleRequest) (*BattleReply, error) {
	panic("implement me")
}