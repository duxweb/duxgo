package admin

import (
	model "dux-project/app/system/models"
	"fmt"
	"github.com/duxweb/go-fast/action"
	"github.com/duxweb/go-fast/database"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

// OperateRes @Resource(app="admin", name = "system.operate", route = "/system/operate")
func OperateRes() action.Result {
	res := action.New[model.LogOperate](model.LogOperate{})

	res.ActionDelete = false
	res.ActionStore = false
	res.ActionEdit = false
	res.ActionCreate = false

	res.ManyAfter(func(data []model.LogOperate, params map[string]any, ctx echo.Context) []model.LogOperate {
		hasIds := lo.Uniq(lo.Map[model.LogOperate, uint](data, func(item model.LogOperate, index int) uint {
			return item.UserID
		}))

		userList := []model.SystemUser{}
		err := database.Gorm().Model(model.SystemUser{}).Where("id IN ?", hasIds).Find(&userList).Error
		if err != nil {
			fmt.Println("err", err)
		}
		userData := lo.KeyBy[uint, model.SystemUser](userList, func(item model.SystemUser) uint {
			return item.ID
		})

		data = lo.Map[model.LogOperate, model.LogOperate](data, func(item model.LogOperate, index int) model.LogOperate {
			item.User = userData[item.UserID]
			return item
		})
		return data
	})

	res.Transform(func(item model.LogOperate, index int) map[string]any {
		user := item.User.(model.SystemUser)
		return map[string]any{
			"id":             item.ID,
			"username":       user.Username,
			"nickname":       user.Nickname,
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
