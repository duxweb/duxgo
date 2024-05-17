package admin

import (
	"context"
	model "dux-project/app/system/models"
	"github.com/duxweb/go-fast/action"
	"github.com/duxweb/go-fast/database"
	"github.com/duxweb/go-fast/helper"
	"github.com/duxweb/go-fast/response"
	"github.com/duxweb/go-fast/validator"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
)

// UserRes @Resource(app="admin", name = "system.user", route = "/system/user")
func UserRes() action.Result {
	res := action.New[model.SystemUser](model.SystemUser{})

	res.Query(func(tx *gorm.DB, e echo.Context) *gorm.DB {
		tx = tx.Preload("Roles")
		return tx
	})

	res.QueryMany(func(tx *gorm.DB, params *gjson.Result, e echo.Context) *gorm.DB {
		keyword := params.Get("keyword").String()
		if keyword != "" {
			keyword = "%" + keyword + "%"
			tx = tx.Where("username like ? OR nickname like ?", keyword, keyword)
		}

		switch params.Get("tab").String() {
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
			"avatar":   item.Avatar,
			"status":   cast.ToBool(item.Status),
			"roles": lo.Map[model.SystemRole, uint](item.Roles, func(item model.SystemRole, index int) uint {
				return item.ID
			}),
		}
	})

	res.Validator(func(data *gjson.Result, e echo.Context) (validator.ValidatorRule, error) {
		return validator.ValidatorRule{
			"username": {Rule: "required", LangMessage: "system.user.validator.username"},
			"nickname": {Rule: "required", LangMessage: "system.user.validator.nickname"},
		}, nil
	})

	res.Format(func(model *model.SystemUser, data *gjson.Result, e echo.Context) error {
		status := data.Get("status").Bool()
		model.Username = data.Get("username").String()
		model.Nickname = data.Get("nickname").String()
		model.Avatar = data.Get("avatar").String()
		model.Status = &status

		password := data.Get("password").String()
		if password != "" {
			model.Password = helper.HashEncode(password)
		}
		return nil
	})

	res.CreateAfter(func(ctx context.Context, data *model.SystemUser, params *gjson.Result) error {
		if !params.Get("password").Exists() {
			return response.BusinessLangError("system.user.validator.password")
		}
		return nil
	})

	res.SaveBefore(func(ctx context.Context, data *model.SystemUser, params *gjson.Result) error {
		roleIds := cast.ToIntSlice(params.Get("roles").Value())
		roles := make([]model.SystemRole, 0)
		err := database.GormCtx(ctx).Model(model.SystemRole{}).Find(&roles, roleIds).Error
		if err != nil {
			return err
		}
		err = database.GormCtx(ctx).Model(data).Association("Roles").Replace(roles)
		if err != nil {
			return err
		}
		return nil
	})

	return res.Result()
}
