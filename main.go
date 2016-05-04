package main

import (
	"github.com/astaxie/beego"
	"github.com/qgweb/new/lib/timestamp"
	_ "github.com/qgweb/statistics/routers"
	"github.com/qgweb/new/lib/convert"
)

func TimeParse(val interface{}) string {
	return timestamp.GetUnixFormat(convert.ToString(val))
}

func main() {
	//beego.BConfig.WebConfig.AutoRender = false
	beego.AddFuncMap("unix", TimeParse)
	beego.Run()
}
