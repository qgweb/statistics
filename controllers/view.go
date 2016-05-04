package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/bitly/go-simplejson"
	"github.com/juju/errors"
	"github.com/qgweb/new/lib/convert"
	"github.com/qgweb/new/lib/timestamp"
	"net/url"

	"sort"
	"strings"
)

type Result struct {
	Timestamp  int
	SourceData map[string]string `数据源数据量`
	AdvertData map[string]string `广告数据量`
	DXData     string            `写入电信的量`
}

type Results []Result

func (this Results) Less(i, j int) bool {
	if this[i].Timestamp > this[j].Timestamp {
		return true
	}
	return false
}

func (this Results) Len() int {
	return len(this)
}

func (this Results) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type ViewController struct {
	beego.Controller
}

func (this *ViewController) request(param url.Values) (map[string]interface{}, error) {
	var host = "http://" + this.Ctx.Input.Context.Request.Host +
		"/api/list?" + param.Encode()
	body, err := httplib.Get(host).Bytes()
	if err != nil {
		return nil, err
	}
	sj, err := simplejson.NewJson(body)
	if err != nil {
		return nil, err
	}
	beego.Info(sj, param)
	if v, _ := sj.Get("ret").String(); v == "0" {
		return sj.Get("data").Map()
	}
	msg, _ := sj.Get("msg").String()
	return nil, errors.New(msg)

}

// 来源数据
func (this *ViewController) sourceData(bt string, et string) map[string]map[string]string {
	var param = url.Values{}
	param.Set("db", "dsource_stats")
	param.Set("bkey", bt)
	param.Set("ekey", et)
	param.Set("limit", "100000")
	info, err := this.request(param)
	if err != nil {
		return nil
	}

	res := make(map[string]map[string]string)
	for k, v := range info {
		ks := strings.Split(k, "_")
		if _, ok := res[ks[1]]; !ok {
			res[ks[1]] = make(map[string]string)
		}
		res[ks[1]][ks[2]] = convert.ToString(v.(map[string]interface{})["Value"])
	}

	return res
}

// 广告数据
func (this *ViewController) advertData(bt string, et string) map[string]map[string]string {
	var param = url.Values{}
	param.Set("db", "advert_stats")
	param.Set("bkey", bt)
	param.Set("ekey", et)
	param.Set("limit", "100000")
	info, err := this.request(param)
	if err != nil {
		return nil
	}
	res := make(map[string]map[string]string)
	for k, v := range info {
		ks := strings.Split(k, "_")
		if _, ok := res[ks[1]]; !ok {
			res[ks[1]] = make(map[string]string)
		}
		res[ks[1]][ks[2]] = convert.ToString(v.(map[string]interface{})["Value"])
	}

	return res
}

// 写入dpi数据
func (this *ViewController) dianxinData(bt string, et string) map[string]map[string]string {
	var param = url.Values{}
	param.Set("db", "dx_stats")
	param.Set("bkey", bt)
	param.Set("ekey", et)
	param.Set("limit", "100000")
	info, err := this.request(param)
	if err != nil {
		return nil
	}
	res := make(map[string]map[string]string)
	for k, v := range info {
		ks := strings.Split(k, "_")
		if _, ok := res[ks[1]]; !ok {
			res[ks[1]] = make(map[string]string)
		}
		count := convert.ToString(v.(map[string]interface{})["Value"])
		res[ks[1]][count] = count
	}

	return res
}

// 合并数据
func (this *ViewController) merageData(source, advert, dianx map[string]map[string]string) Results {
	res := make(Results, 0, 24)
	searchFun := func(t int) int {
		for k, v := range res {
			if v.Timestamp == t {
				return k
			}
		}
		return -1
	}

	for k, v := range source {
		ki := convert.ToInt(k)
		if i := searchFun(ki); i != -1 {
			res[i].SourceData = v
		} else {
			rt := Result{}
			rt.Timestamp = ki
			rt.SourceData = v
			res = append(res, rt)
		}
	}

	for k, v := range advert {
		ki := convert.ToInt(k)
		if i := searchFun(ki); i != -1 {
			res[i].AdvertData = v
		} else {
			rt := Result{}
			rt.Timestamp = ki
			rt.AdvertData = v
			res = append(res, rt)
		}
	}

	for k, v := range dianx {
		ki := convert.ToInt(k)
		vi := ""
		for kk, _ := range v {
			vi = kk
		}
		if i := searchFun(ki); i != -1 {
			res[i].DXData = vi
		} else {
			rt := Result{}
			rt.Timestamp = ki
			rt.DXData = vi
			res = append(res, rt)
		}
	}

	return res
}

func (this *ViewController) Index() {
	var (
		pre    = this.GetString("qy", "zj")
		dbtime = convert.ToString(convert.ToInt(timestamp.GetDayTimestamp(0)) - 86400)
		detime = timestamp.GetHourTimestamp(0)
		btime  = this.GetString("btime", dbtime)
		etime  = this.GetString("etime", detime)
	)

	pbtime := pre + "_" + btime
	petime := pre + "_" + etime

	source := this.sourceData(pbtime, petime)
	advert := this.advertData(pbtime, petime)
	dianx := this.dianxinData(pbtime, petime)

	beego.Error(source)
	beego.Error(advert)
	beego.Error(dianx)

	result := this.merageData(source, advert, dianx)
	sort.Sort(result)
	this.Data["info"] = result
	this.TplName = "view.tpl"
}
