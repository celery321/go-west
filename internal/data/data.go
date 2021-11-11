package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go-west/internal/conf"
	"go-west/pkg/cache/redis"
	"go-west/pkg/container/pool"
	"go-west/pkg/database/sql"
	"go-west/pkg/http/client"
	"go-west/pkg/log"
	xtime "go-west/pkg/time"
)



// Data .
type Data struct {
	db *sql.DB
	rd *redis.Pool
	httpclient *http.Client
}


// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	mysqlConf := &sql.Config{
		DSN:          c.Mysql.Dsn,
		Active:       int(c.Mysql.Active),
		Idle:         int(c.Mysql.Idle),
		IdleTimeout:  xtime.Str2Time(c.Mysql.IdleTimeout) ,
		QueryTimeout: xtime.Str2Time(c.Mysql.QueryTimeout),
		ExecTimeout : xtime.Str2Time(c.Mysql.ExecTimeout),
		TranTimeout: xtime.Str2Time(c.Mysql.TranTimeout),
	}
	redisConf := &redis.Config{
		Name:          c.Redis.Name,
		Proto:         c.Redis.Proto,
		Addr:          c.Redis.Addr,
		DialTimeout  : xtime.Str2Time(c.Redis.DialTimeout),
		ReadTimeout  : xtime.Str2Time(c.Redis.ReadTimeout),
		WriteTimeout : xtime.Str2Time(c.Redis.WriteTimeout),
		Config: &pool.Config{
				Idle: int(c.Redis.Config.Idle),
				Active: int(c.Redis.Config.Active),
				IdleTimeout: xtime.Str2Time(c.Redis.Config.IdleTimeout),
		},
	}


	d := &Data {
		 db: sql.NewMySQL(mysqlConf),
		 rd: redis.NewPool(redisConf),
		 httpclient: client.NewHttpClient(logger),
	}

	return d, cleanup, nil
}

// Ping ping dao
func (d *Data) Ping(c context.Context) (err error) {
	if err := d.PingRedis(c); err != nil {
		return err
	}

	if err := d.PingMysql(c); err != nil {
		return err
	}

	return
}

func (d *Data) PingMysql(c context.Context) (err error) {
	if err = d.db.Ping(c); err != nil {
		return err
	}
	return
}

func (d *Data) PingRedis(c context.Context) (err error) {
	conn := d.rd.Get(c)
	_, err = conn.Do("SET", "PING", "PONG")
	if err != nil {
		fmt.Printf("err===%v\n", err)
	}
	defer conn.Close()
	return
}


