package admin

import (
	model "dux-project/app/system/models"
	"github.com/duxweb/go-fast/action"
	"github.com/duxweb/go-fast/database"
	"github.com/duxweb/go-fast/helper"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// UserRes @Resource(app="admin", name = "system.user", route = "/system/user")
func UserRes() action.Result {
	res := action.New[model.SystemUser](model.SystemUser{})

	res.Query(func(tx *gorm.DB, params map[string]any, e echo.Context) *gorm.DB {
		tx = tx.Preload("Roles")
		return tx
	})

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

	res.Transform(func(item model.SystemUser, index int) map[string]any {
		return map[string]any{
			"id":       item.ID,
			"username": item.Username,
			"nickname": item.Nickname,
			"avatar":   item.Avatar,
			"status":   cast.ToBool(item.Status),
			"roles": lo.Map[model.SystemRole, uint](item.Roles, func(item model.SystemRole, index int) uint {
				return item.ID
			}),
		}
	})

	res.Format(func(model *model.SystemUser, data map[string]any, e echo.Context) error {
		model.Username = cast.ToString(data["username"])
		model.Nickname = cast.ToString(data["nickname"])
		model.Avatar = cast.ToString(data["avatar"])
		model.Status = cast.ToBool(data["status"])

		password := cast.ToString(data["password"])
		if password != "" {
			model.Password = helper.HashEncode(password)
		}
		return nil
	})

	res.SaveBefore(func(data *model.SystemUser, params map[string]any) error {
		roleIds := cast.ToIntSlice(params["roles"])
		roles := []model.SystemRole{}
		err := database.Gorm().Model(model.SystemRole{}).Find(&roles, roleIds).Error
		if err != nil {
			return err
		}
		err = database.Gorm().Model(data).Association("Roles").Replace(roles)
		if err != nil {
			return err
		}
		return nil
	})

	return res.Result()
}
