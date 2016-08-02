package controllers

import (
	"strings"
	"github.com/astaxie/beego"
	"github.com/juju/errors"
	"github.com/qgweb/new/lib/convert"
	"github.com/qgweb/statistics/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Return(v interface{}, err error) {
	if err != nil {
		c.Ctx.Output.JSON(map[string]string{
			"ret":  "1",
			"msg":  err.Error(),
			"data": "",
		}, true, false)
	} else {
		c.Ctx.Output.JSON(map[string]interface{}{

			"ret":  "0",
			"msg":  "",
			"data": v,
		}, true, false)
	}
}

func (c *MainController) Get() {
	var (
		db  = c.GetString("db")
		key = c.GetString("key")
		sdb = models.SDB{}
	)

	if strings.TrimSpace(db) == "" {
		c.Return("", errors.New("db不能为空"))
		return
	}
	if strings.TrimSpace(key) == "" {
		c.Return("", errors.New("key不能为空"))
		return
	}

	v, err := sdb.Get(db, key)
	c.Return(v.String(), err)
}

func (c *MainController) List() {
	var (
		db       = c.GetString("db")
		bkey     = c.GetString("bkey")
		ekey     = c.GetString("ekey")
		limit, _ = c.GetInt64("limit", 10)
		sdb      = models.SDB{}
	)

	if strings.TrimSpace(db) == "" {
		c.Return("", errors.New("db不能为空"))
		return
	}

	v, err := sdb.Scan(db, bkey, ekey, limit)
	c.Return(v, err)
}

func (c *MainController) Post() {
	var (
		db    = c.GetString("db")
		key   = c.GetString("key")
		value = c.GetString("value")
		opt   = c.GetString("opt")
		sdb   = models.SDB{}
	)

	if strings.TrimSpace(db) == "" {
		c.Return("", errors.New("db不能为空"))
		return
	}
	if strings.TrimSpace(key) == "" {
		c.Return("", errors.New("key不能为空"))
		return
	}
	if strings.TrimSpace(value) == "" {
		c.Return("", errors.New("value不能为空"))
		return
	}

	switch opt {
	case "incr":
		v, err := sdb.Incr(db, key, convert.ToInt64(value))
		c.Return(v, err)
	default:
		err := sdb.Set(db, key, value)
		c.Return("", err)
		return
	}
}

func (c *MainController) Delete() {
	var (
		db  = c.GetString("db")
		key = c.GetString("key")
		sdb = models.SDB{}
	)

	if strings.TrimSpace(db) == "" {
		c.Return("", errors.New("db不能为空"))
		return
	}

	if key == "" {
		err := sdb.DBDel(db)
		c.Return("", err)
		return
	}

	err := sdb.Del(db, key)
	c.Return("", err)
	return
}
