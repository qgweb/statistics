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

type DataResult struct {
	Timestamp string
	Data      map[string]string
}

type ViewDataResultAry struct {
	Timestamp string
	Data      DataResultAry
}

type DataResultAry []DataResult

func (dr DataResultAry) Len() int {
	return len(dr)
}

func (dr DataResultAry) Less(i, j int) bool {
	return dr[i].Timestamp < dr[i].Timestamp
}

func (dr DataResultAry) Swap(i, j int) {
	dr[i], dr[j] = dr[j], dr[i]
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
	if v, _ := sj.Get("ret").String(); v == "0" {
		return sj.Get("data").Map()
	}
	msg, _ := sj.Get("msg").String()
	return nil, errors.New(msg)

}

// 来源数据
func (this *ViewController) sourceData(bt string, et string) (res DataResultAry) {
	var param = url.Values{}
	param.Set("db", "dsource_stats")
	param.Set("bkey", bt)
	param.Set("ekey", et)
	param.Set("limit", "100000")
	info, err := this.request(param)
	if err != nil {
		return nil
	}
	res = make(DataResultAry, 0, len(info))
	for k, v := range info {
		ks := strings.Split(k, "_")
		dr := DataResult{}
		dr.Timestamp = ks[1]
		dr.Data = map[string]string{
			ks[2]: convert.ToString(v.(map[string]interface{})["Value"]),
		}
		beego.Info(dr)
		res = append(res, dr)
	}
	sort.Sort(res)

	for _, v := range res {
		beego.Error(v.Timestamp)
	}
	return res
}

// 广告数据
func (this *ViewController) advertData(bt string, et string) (res DataResultAry) {
	var param = url.Values{}
	param.Set("db", "advert_stats")
	param.Set("bkey", bt)
	param.Set("ekey", et)
	param.Set("limit", "100000")
	info, err := this.request(param)
	if err != nil {
		return nil
	}
	res = make(DataResultAry, 0, len(info))
	for k, v := range info {
		ks := strings.Split(k, "_")
		dr := DataResult{}
		dr.Timestamp = ks[1]
		dr.Data = map[string]string{
			ks[2]: convert.ToString(v.(map[string]interface{})["Value"]),
		}
		beego.Info(dr)
		res = append(res, dr)
	}
	sort.Sort(res)

	for _, v := range res {
		beego.Error(v.Timestamp)
	}
	return res
}

// 写入dpi数据
func (this *ViewController) dianxinData(bt string, et string) (res DataResultAry) {
	var param = url.Values{}
	param.Set("db", "dx_stats")
	param.Set("bkey", bt)
	param.Set("ekey", et)
	param.Set("limit", "100000")
	info, err := this.request(param)
	if err != nil {
		return nil
	}
	res = make(DataResultAry, 0, len(info))
	for k, v := range info {
		ks := strings.Split(k, "_")
		dr := DataResult{}
		dr.Timestamp = ks[1]
		dr.Data = map[string]string{
			convert.ToString(v.(map[string]interface{})["Value"]): "0",
		}
		res = append(res, dr)
	}
	sort.Sort(res)
	return res
}

// 合并数据
func (this *ViewController) merageData(btime int, etime int, ds ...DataResultAry) []ViewDataResultAry {
	var res = make([]ViewDataResultAry, 0, 24)
	for t := etime; t > btime; t = t - 3600 {
		var drs = ViewDataResultAry{}
		drs.Timestamp = timestamp.GetUnixFormat(convert.ToString(t))
		drs.Data = make(DataResultAry, 0, len(ds))
		for _, v := range ds {
			for _, vv := range v {
				if convert.ToString(t) == vv.Timestamp {
					vv.Timestamp = timestamp.GetUnixFormat(vv.Timestamp)
					drs.Data = append(drs.Data, vv)
					break
				}
			}
		}
		res = append(res, drs)
	}
	return res
}

func (this *ViewController) Index() {
	var (
		pre   = this.GetString("qy", "zj")
		btime = this.GetString("btime", pre+"_"+
			convert.ToString(convert.ToInt(timestamp.GetDayTimestamp(0))-3600))
		etime = this.GetString("etime", pre+"_"+timestamp.GetHourTimestamp(-1))
	)

	source := this.sourceData(btime, etime)
	advert := this.advertData(btime, etime)
	dianx := this.dianxinData(btime, etime)
	beego.Error(source)
	beego.Error(advert)
	beego.Error(dianx)
	this.Data["info"] = this.merageData(convert.ToInt(btime), convert.ToInt(etime),
		source, advert, dianx)
	this.TplName = "view.tpl"
}
