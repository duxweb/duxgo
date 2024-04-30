package admin

import (
	model "dux-project/app/system/models"
	"github.com/duxweb/go-fast/action"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// UserRes @Resource(app="admin", name = "system.user", route = "/system/user")
func UserRes() action.Result {
	res := action.New[*model.SystemUser](&model.SystemUser{})
	res.QueryMany(func(tx *gorm.DB, params map[string]any, e echo.Context) *gorm.DB {

		keyword := cast.ToString(params["keyword"])
		if keyword != "" {
			keyword = "%" + keyword + "%"
			tx = tx.Where("username like ? OR nickname like ?", keyword, keyword)
		}

		switch params["tab"] {
		case "1":
			tx = tx.Where("status = ?", "1")
			break
		case "2":
			tx = tx.Where("status = ?", "0")
			break
		}
		return tx
	})
	res.Transform(func(item *model.SystemUser, index int) map[string]any {
		return map[string]any{
			"id":       item.ID,
			"username": item.Username,
			"nickname": item.Nickname,
			"status":   cast.ToBool(item.Status),
		}
	})
	return res.Result()
}
