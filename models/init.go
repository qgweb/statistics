package models

import (
	"github.com/astaxie/beego"
	"github.com/qgweb/gossdb"
	"os"
)

var (
	dbpool *gossdb.Connectors
)

func init() {
	var (
		dbhost    = beego.AppConfig.String("db.host")
		dbport, _ = beego.AppConfig.Int("db.port")
		err       error
	)

	dbpool, err = gossdb.NewPool(&gossdb.Config{
		Host:             dbhost,
		Port:             dbport,
		MinPoolSize:      5,
		MaxPoolSize:      50,
		AcquireIncrement: 5,
		GetClientTimeout: 10,
		MaxWaitSize:      1000,
		MaxIdleTime:      1,
		HealthSecond:     2,
	})

	if err != nil {
		beego.Error(err)
		os.Exit(-2)
	}
}
