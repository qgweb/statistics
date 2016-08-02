package models

import (
	"github.com/astaxie/beego"
	"github.com/seefan/gossdb"
	"os"
)

var (
	dbpool *gossdb.Connectors
)

func init() {
	var (
		dbhost = beego.AppConfig.String("db.host")
		dbport, _ = beego.AppConfig.Int("db.port")
		err error
	)

	dbpool, err = gossdb.NewPool(&gossdb.Config{
		Host:             dbhost,
		Port:             dbport,
	})

	if err != nil {
		beego.Error(err)
		os.Exit(-2)
	}
}
