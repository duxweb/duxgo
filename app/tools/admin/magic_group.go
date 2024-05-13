package admin

import (
	model "dux-project/app/tools/models"
	"github.com/duxweb/go-fast/action"
	"github.com/duxweb/go-fast/cache"
	"github.com/duxweb/go-fast/validator"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
)

// MagicGroupRes @Resource(app="admin", name = "tools.magicGroup", route = "/tools/magicGroup")
func MagicGroupRes() action.Result {
	res := action.New[model.ToolsMagicGroup](model.ToolsMagicGroup{})

	res.Tree = true
	res.Pagination.Status = false

	res.Transform(func(item *model.ToolsMagicGroup, index int) map[string]any {
		data := map[string]any{
			"id":        item.ID,
			"parent_id": lo.Ternary[any](item.ParentID == 0, nil, item.ParentID),
			"name":      item.Name,
			"label":     item.Label,
			"icon":      item.Icon,
		}

		var results []map[string]any
		for i, vo := range item.Children {
			results = append(results, res.TransformFun(&vo, i))
		}
		data["children"] = results

		return data
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
		model.ParentID = cast.ToUint(data.Get("parent_id").Uint())

		return nil
	})

	res.SaveBefore(func(data *model.ToolsMagicGroup, params *gjson.Result) error {
		key := []byte("magic.menus")
		cache.Injector().Del(key)
		return nil
	})

	return res.Result()
}
