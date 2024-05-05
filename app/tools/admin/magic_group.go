package admin

import (
	model "dux-project/app/tools/models"
	"github.com/duxweb/go-fast/action"
	"github.com/duxweb/go-fast/validator"
	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
)

// MagicGroupRes @Resource(app="admin", name = "tools.magicGroup", route = "/tools/magicGroup")
func MagicGroupRes() action.Result {
	res := action.New[model.ToolsMagicGroup](model.ToolsMagicGroup{})

	res.Transform(func(item *model.ToolsMagicGroup, index int) map[string]any {
		return map[string]any{
			"id":    item.ID,
			"name":  item.Name,
			"label": item.Label,
			"icon":  item.Icon,
		}
	})

	res.Validator(func(data *gjson.Result, e echo.Context) (validator.ValidatorRule, error) {

		validatorData := validator.ValidatorRule{
			"name":  {Rule: "required", Message: "请填写组名"},
			"label": {Rule: "required", Message: "请填写组描述"},
			"icon":  {Rule: "required", Message: "请填写图标名"},
		}

		return validatorData, nil
	})

	res.Format(func(model *model.ToolsMagicGroup, data *gjson.Result, e echo.Context) error {
		model.Name = data.Get("name").String()
		model.Label = data.Get("label").String()
		model.Icon = data.Get("icon").String()
		return nil
	})

	return res.Result()
}
