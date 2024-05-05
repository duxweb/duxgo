package admin

import (
	model "dux-project/app/system/models"
	"github.com/duxweb/go-fast/action"
	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
	"strings"
)

// OperateRes @Resource(app="admin", name = "system.operate", route = "/system/operate")
func OperateRes() action.Result {
	res := action.New[model.LogOperateUser](model.LogOperateUser{})

	res.ActionDelete = false
	res.ActionStore = false
	res.ActionEdit = false
	res.ActionCreate = false

	res.QueryMany(func(tx *gorm.DB, params *gjson.Result, e echo.Context) *gorm.DB {
		tx = tx.Preload("User").Where("user_type", "admin").Order("id desc")

		if params.Get("user").Exists() {
			tx = tx.Where("user_id = ?", params.Get("user").Int())
		}

		if params.Get("method").Exists() {
			tx = tx.Where("request_method = ?", strings.ToUpper(params.Get("method").String()))
		}

		dates := params.Get("date").Array()
		if len(dates) > 0 {
			tx = tx.Where("created_at BETWEEN ? AND ?", dates[0].String(), dates[1].String())
		}
		return tx
	})

	res.Transform(func(item *model.LogOperateUser, index int) map[string]any {
		return map[string]any{
			"id":             item.ID,
			"username":       item.User.Username,
			"nickname":       item.User.Nickname,
			"request_method": item.RequestMethod,
			"request_url":    item.RequestUrl,
			"request_time":   item.RequestTime,
			"request_params": item.RequestParams,
			"route_name":     item.RouteName,
			"route_title":    item.RouteTitle,
			"client_ua":      item.ClientUa,
			"client_ip":      item.ClientIp,
			"client_browser": item.ClientBrowser,
			"client_device":  item.ClientDevice,
			"time":           item.CreatedAt,
		}
	})

	return res.Result()
}
