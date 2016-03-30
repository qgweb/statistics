package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/qgweb/statistics/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	sdb := models.SDB{}
	sdb.Set("tes", "name_1", 1)
	sdb.Set("tes", "name_2", 2)
	sdb.Set("tes", "name_3", 3)
	v, err := sdb.Scan("tes", "name_", "", 100)
	fmt.Fprintln(c.Ctx.ResponseWriter, v, err)
	v1, err1 := sdb.Get("tes", "name_1")
	fmt.Fprintln(c.Ctx.ResponseWriter, v1, err1)
	sdb.Del("tes", "name_1")
	v, err = sdb.Scan("tes", "name_", "", 100)
	fmt.Fprintln(c.Ctx.ResponseWriter, v, err)
	sdb.Incr("tes", "name_1", 1)
	v, err = sdb.Scan("tes", "name_", "", 100)
	fmt.Fprintln(c.Ctx.ResponseWriter, v, err)
	beego.Error(sdb.Size("tes"))
}
