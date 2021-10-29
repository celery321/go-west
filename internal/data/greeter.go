package data

import (
	"context"
	"go-west/internal/biz"
	"go-west/pkg/log"
)

const (
	_addStu  = "INSERT INTO stu SET name=?"
	_getStus = "SELECT name FROM stu where name=?"
)
var _ biz.GreeterRepo = (*greeterRepo)(nil)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

func (r *greeterRepo) SetGreeter(ctx context.Context, g *biz.Greeter) error {
	_, err := r.data.db.Exec(ctx, _addStu, g.Hello)
	return err
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
	rows, err := r.data.db.Query(ctx, _getStus, g.Hello)
	if err != nil {
		r.log.Error("d.biliDM.Query(%v,%v) error(%v)", _getStus,  g.Hello, err)
		return l, err
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
	return l, err
}

