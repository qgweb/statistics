package main

import (
	"github.com/astaxie/beego"
	_ "github.com/qgweb/statistics/routers"
)

func main() {
	//beego.BConfig.WebConfig.AutoRender = false
	beego.Run()
}
