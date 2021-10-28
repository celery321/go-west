package data

import (
	"context"
	"fmt"
	"go-west/internal/conf"
	"go-west/pkg/database/sql"
	"go-west/pkg/log"
)



// Data .
type Data struct {
	// TODO wrapped database client
	db *sql.DB
}



// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	fmt.Printf("c.mysql.dsn==%v", c.Mysql)
	d := &Data{
		 db: sql.NewMySQL(sql.DSN(c.Mysql.Dsn)),
	}


	return d, cleanup, nil
}

// Ping ping dao
func (d *Data) Ping(c context.Context) (err error) {
	fmt.Printf("ctx===%v",c )
	if err = d.db.Ping(c); err != nil {
		return err
	}
	return
}