package admin

import (
	model "dux-project/app/system/models"
	"github.com/duxweb/go-fast/action"
	"github.com/duxweb/go-fast/permission"
	"github.com/duxweb/go-fast/response"
	"github.com/duxweb/go-fast/validator"
	"github.com/gookit/goutil/jsonutil"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/tidwall/gjson"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"strings"
)

// RoleRes @Resource(app="admin", name = "system.role", route = "/system/role")
func RoleRes() action.Result {
	res := action.New[model.SystemRole](model.SystemRole{})

	res.QueryMany(func(tx *gorm.DB, params *gjson.Result, e echo.Context) *gorm.DB {
		keyword := params.Get("keyword").String()
		if keyword != "" {
			keyword = "%" + keyword + "%"
			tx = tx.Where("name like ?", keyword)
		}
		return tx
	})

	res.Transform(func(item *model.SystemRole, index int) map[string]any {
		permissionData := map[string]bool{}
		_ = jsonutil.DecodeString(item.Permission.String(), &permissionData)

		permissions := []string{}
		for k, v := range permissionData {
			if v {
				permissions = append(permissions, k)
			}
		}

		return map[string]any{
			"id":         item.ID,
			"name":       item.Name,
			"permission": permissions,
		}
	})

	res.Validator(func(data *gjson.Result, e echo.Context) (validator.ValidatorRule, error) {
		return validator.ValidatorRule{
			"name": {Rule: "required", Message: "请填写名称"},
		}, nil
	})

	res.Format(func(model *model.SystemRole, data *gjson.Result, e echo.Context) error {
		permissionData := map[string]bool{}

		permissionsParams := data.Get("permission").Array()

		permissions := permission.Get("admin").GetData()
		permissionRest := []string{}
		for _, vo := range permissionsParams {
			if !strings.Contains(vo.String(), "group:") {
				permissionRest = append(permissionRest, vo.String())
			}
		}
		if len(permissionsParams) > 0 {
			for _, item := range permissions {
				permissionData[item] = lo.Ternary[bool](lo.IndexOf[string](permissionRest, item) != -1, true, false)
			}
		}

		model.Name = data.Get("name").String()
		model.Permission = datatypes.JSON(jsonutil.MustString(permissionData))

		return nil
	})

	return res.Result()
}

// RolePermission @Action(method = "GET", name = "permission", route = "/permission")
func RolePermission(c echo.Context) error {

	data := permission.Get("admin").Get()

	return response.Send(c, response.Data{
		Data: data,
	}, 200)
}
