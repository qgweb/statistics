package routers

import (
	"github.com/qgweb/statistics/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/api/get", &controllers.MainController{}, "get:Get")
	beego.Router("/api/create", &controllers.MainController{},"post:Post")
	beego.Router("/api/update", &controllers.MainController{},"post:Post")
	beego.Router("/api/list", &controllers.MainController{},"get:List")
	beego.Router("/api/delete", &controllers.MainController{},"delete:Delete")
}
