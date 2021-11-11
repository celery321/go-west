package data

import (
	"context"
	pb "go-west/api/v1"
	"go-west/internal/biz"
	"go-west/pkg/log"
)

var _ biz.BattleRepo = (*battleRepo)(nil)

type battleRepo struct {
	data *Data
	log  *log.Helper
}

func (r *battleRepo) Start(ctx context.Context, request *biz.BattleRequest) (*biz.BattleReply, error) {
	panic("implement me")
}

func (r *battleRepo) NStart(ctx context.Context, g *biz.BattleRequest) (*biz.BattleReply, error) {
	httpCli := pb.NewBattleHTTPClient(r.data.httpclient)

	req := &pb.BattleRequest{
		MonsterA: g.MonsterA,
		MonsterB: g.MonsterB,
		Address: g.Address,
		BattleLevel: g.BattleLevel,
	}
	 rs , err := httpCli.Start(ctx, req)
	 if err != nil {
		return nil, err
	}
	d := &biz.BattleReply{
		Code:     rs.Code,
		Data:      biz.Data{
			BattleLevel:    rs.Data.BattleLevel,
			BpFragmentNum:  rs.Data.BpFragmentNum,
			BpPotionNum:    rs.Data.BpPotionNum,
			ChallengeExp:   rs.Data.ChallengeExp,
			ChallengeLevel: rs.Data.ChallengeLevel,
		},
		ErrorText: rs.ErrorText,
		Message:   rs.Message,
		Result:    rs.Result,
	}

	return d, nil

}

// NewBattleRepo NewGreeterRepo .
func NewBattleRepo(data *Data, logger log.Logger) biz.BattleRepo {
	return &battleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}


