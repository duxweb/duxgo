package admin

import (
	model "dux-project/app/system/models"
	"github.com/duxweb/go-fast/action"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// OperateRes @Resource(app="admin", name = "system.operate", route = "/system/operate")
func OperateRes() action.Result {
	res := action.New[model.LogOperateUser](model.LogOperateUser{})

	res.ActionDelete = false
	res.ActionStore = false
	res.ActionEdit = false
	res.ActionCreate = false

	res.QueryMany(func(tx *gorm.DB, params map[string]any, e echo.Context) *gorm.DB {
		return tx.Preload("User").Where("user_type", "admin").Order("id desc")
	})

	res.Transform(func(item model.LogOperateUser, index int) map[string]any {
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
