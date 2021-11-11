package sql

import (
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/go-kratos/kratos/pkg/net/netutil/breaker"
	"github.com/go-kratos/kratos/pkg/time"
	// database driver
	_ "github.com/go-sql-driver/mysql"
)

// Config mysql config.
type Config struct {
	DSN          string          // write data source name.
	ReadDSN      []string        // read data source name.
	Active       int             // pool
	Idle         int             // pool
	IdleTimeout  time.Duration   // connect max life time.
	QueryTimeout time.Duration   // query sql timeout
	ExecTimeout  time.Duration   // execute sql timeout
	TranTimeout  time.Duration   // transaction sql timeout
	Breaker      *breaker.Config // breaker
}

type Option func(o *Config)

// DSN Address with server address.
func DSN(dsn string) Option {
	return func(s *Config) {
		s.DSN = dsn
	}
}


// NewMySQL new db and retry connection when has error.
func NewMySQL(opts ...Option) (db *DB) {
	options := &Config{
			IdleTimeout: 400,
			QueryTimeout: 50000000,
			ExecTimeout: 50000000,
			TranTimeout: 50000000,
	}
	for _, o := range opts {
		o(options)
	}
	if options.QueryTimeout == 0 || options.ExecTimeout == 0 || options.TranTimeout == 0 {
		panic("mysql must be set query/execute/transction timeout")
	}
	db, err := Open(options)
	if err != nil {
		log.Error("open mysql error(%v)", err)
		panic(err)
	}
	return
}
