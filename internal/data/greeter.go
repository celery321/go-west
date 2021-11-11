package data

import (
	"context"
	v1 "go-west/api/v1"
	"go-west/internal/biz"
	"go-west/pkg/database/sql"
	"go-west/pkg/log"
)

const (
	_addStu  = "INSERT INTO stu SET name=?"
	_getStus = "SELECT name FROM stu where name=?"
	_sleep = "select sleep(?)"
	_sleep1 = "select ?"
)
var _ biz.GreeterRepo = (*greeterRepo)(nil)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

func (r *greeterRepo) SetGreeter(ctx context.Context, g *biz.Greeter) (err error ){
	//go r.data.Ping(ctx)
	var tx           *sql.Tx

	if tx, err = r.data.db.Begin(ctx); err != nil {
		return
	}
	//time.Sleep(8 * time.Second)
	_, err = tx.Exec(_sleep1, g.Hello)
	if err != nil {
		tx.Rollback()
		r.log.Errorf("SetGreeter: rows.Scan() error(%v)", err)
		return v1.ErrorContentMissing("SetGreeter")
	}

	if err = tx.Commit(); err != nil {
		return
	}
	return


}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) GetGreeter(ctx context.Context, g *biz.Greeter) (l []*biz.Greeter, err error) {
	l = make([]*biz.Greeter, 0)
	rows, err := r.data.db.Query(ctx, _sleep1, g.Hello)
	if err != nil {
		return l, v1.ErrorUserNotFound("user %s not found", err, _sleep, g.Hello)
	}
	defer rows.Close()
	for rows.Next() {
		hs := new(biz.Greeter)
		if err = rows.Scan(&hs.Hello); err != nil {
			r.log.Error("TransfrerList: rows.Scan() error(%v)", err)
			return l, err
		}
		l = append(l, hs)
	}
	if err = rows.Err(); err != nil {
		r.log.Error("rows.Err() error(%v)", err)
	}

	// redis

	return l, err
}

