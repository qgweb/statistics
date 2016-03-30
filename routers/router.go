package routers

import (
	"github.com/qgweb/statistics/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
