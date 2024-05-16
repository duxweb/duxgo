package admin

import (
	"bufio"
	"github.com/duxweb/go-fast/config"
	"github.com/duxweb/go-fast/global"
	"github.com/duxweb/go-fast/helper"
	"github.com/duxweb/go-fast/monitor"
	"github.com/duxweb/go-fast/response"
	"github.com/go-resty/resty/v2"
	"github.com/golang-module/carbon/v2"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/tidwall/gjson"
	"os"
	"runtime"
	"strings"
	"time"
)

// @RouteGroup(app="admin", name="system.total", route="/system/total")

// Total @Route(method="GET", name="total", route="")
func Total(ctx echo.Context) error {

	info := monitor.GetMonitorInfo()
	return response.Send(ctx, response.Data{
		Message: "ok",
		Data: map[string]any{
			"go":        runtime.Version(),
			"os":        info.OsName,
			"boot_time": carbon.Now().DiffAbsInMinutes(carbon.Parse(info.BootTime)),
			"log_size":  info.LogSizeF,
			"event":     0,
			"schedule":  0,
			"queue":     0,
			"time":      carbon.Now().ToDateTimeString(),
		},
	})
}

// Hardware @Route(method="GET", name="hardware", route="/hardware")
func Hardware(ctx echo.Context) error {

	filePath := global.DataDir + "logs/monitor.log"

	file, err := os.Open(filePath)
	if err != nil {
		return response.Send(ctx, response.Data{
			Message: "ok",
		})
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lastLine := ""
	for scanner.Scan() {
		lastLine = scanner.Text()
	}
	curData := gjson.Parse(lastLine)

	return response.Send(ctx, response.Data{
		Message: "ok",
		Data: map[string]any{
			"cpu":       curData.Get("CpuPercent").Float(),
			"mem":       curData.Get("MemPercent").Float(),
			"thread":    curData.Get("ThreadCount").Int(),
			"goroutine": curData.Get("GoroutineCount").Int(),
		},
	})
}

// Speed @Route(method="GET", name="speed", route="/speed")
func Speed(ctx echo.Context) error {

	client := resty.New()
	defer client.Clone()
	get, err := client.
		SetTimeout(10 * time.Second).
		R().
		Post("https://www.cesu.net/user/getPingList")
	if err != nil {
		return err
	}

	resData := gjson.ParseBytes(get.Body())

	if resData.Get("code").Int() > 0 {
		return response.BusinessError(resData.Get("msg").String())
	}

	if !resData.Get("data").Exists() {
		return response.BusinessError("无法获取测速节点")
	}

	values := map[string]bool{}
	for _, result := range resData.Get("data").Array() {
		values[result.Get("node_type_txt").String()] = true
	}

	token := config.Load("use").GetString("extend.speed")
	return response.Send(ctx, response.Data{
		Message: "ok",
		Data:    lo.Keys[string, bool](values),
		Meta: map[string]any{
			"token": token,
		},
	})
}

// SpeedSubmit @Route(method="POST", name="speed", route="/speed")
func SpeedSubmit(ctx echo.Context) error {

	data, err := helper.Body(ctx)
	if err != nil {
		return response.BusinessError(err.Error())
	}

	if !data.Get("net").Exists() {
		return response.BusinessError("请选择线路")
	}

	token := config.Load("use").GetString("extend.speed")

	if data.Get("token").String() == "" && token == "" {
		return response.BusinessError("请输入 Token")
	} else {
		token = data.Get("token").String()
	}

	net := []string{}
	for _, result := range data.Get("net").Array() {
		net = append(net, result.String())
	}

	client := resty.New()
	defer client.Clone()
	get, err := client.
		SetTimeout(10 * time.Second).
		R().
		Post("https://www.cesu.net/user/getPingList")
	if err != nil {
		return err
	}

	resData := gjson.ParseBytes(get.Body()).Get("data")

	if resData.Get("code").Int() > 0 {
		return response.BusinessError(resData.Get("msg").String())
	}

	if !resData.Exists() {
		return response.BusinessError("无法获取测速节点")
	}

	chinaMaps := SpeedChinaMaps()

	nids := []string{}
	nodes := []map[string]any{}
	for _, result := range resData.Array() {
		nodeType := result.Get("node_type_txt").String()
		if lo.IndexOf[string](net, nodeType) == -1 {
			continue
		}

		provincial := result.Get("provincial").String()
		chinaProvince, ok := chinaMaps[provincial]
		nodes = append(nodes, map[string]any{
			"nid":            result.Get("id").Int(),
			"china_province": chinaProvince,
			"province":       lo.Ternary[string](ok, chinaProvince, provincial),
			"city":           result.Get("city").String(),
			"type":           nodeType,
		})

		nids = append(nids, result.Get("id").String())
	}

	if len(nodes) == 0 {
		return response.BusinessError("无法获取测速节点")
	}

	get, err = client.
		SetTimeout(10 * time.Second).
		R().
		SetFormData(map[string]string{
			"token": token,
			"type":  "3",
			"url":   data.Get("url").String(),
			"nid":   strings.Join(nids, ","),
		}).
		Post("https://www.cesu.net/user/GetTaskId")

	if err != nil {
		return err
	}

	resData = gjson.ParseBytes(get.Body())

	if resData.Get("code").Int() > 0 {
		return response.BusinessError(resData.Get("msg").String())
	}

	return response.Send(ctx, response.Data{
		Message: "ok",
		Data: map[string]any{
			"taskId": resData.Get("data.taskId").String(),
			"nodes":  nodes,
		},
	})
}

func SpeedChinaMaps() map[string]string {
	return map[string]string{
		"山西":  "山西省",
		"辽宁":  "辽宁省",
		"吉林":  "吉林省",
		"黑龙江": "黑龙江省",
		"江苏":  "江苏省",
		"浙江":  "浙江省",
		"安徽":  "安徽省",
		"福建":  "福建省",
		"江西":  "江西省",
		"山东":  "山东省",
		"河南":  "河南省",
		"湖北":  "湖北省",
		"湖南":  "湖南省",
		"广东":  "广东省",
		"海南":  "海南省",
		"四川":  "四川省",
		"贵州":  "贵州省",
		"云南":  "云南省",
		"陕西":  "陕西省",
		"甘肃":  "甘肃省",
		"青海":  "青海省",
		"台湾":  "台湾省",
		"北京":  "北京市",
		"天津":  "天津市",
		"上海":  "上海市",
		"重庆":  "重庆市",
		"香港":  "香港特别行政区",
		"澳门":  "澳门特别行政区",
		"内蒙":  "内蒙古自治区",
		"广西":  "广西壮族自治区",
		"西藏":  "西藏自治区",
		"宁夏":  "宁夏回族自治区",
		"新疆":  "新疆维吾尔族自治区",
	}
}
