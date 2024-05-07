package admin

import (
	model "dux-project/app/tools/models"
	"github.com/duxweb/go-fast/action"
	"github.com/duxweb/go-fast/response"
	"github.com/duxweb/go-fast/validator"
	"github.com/gookit/goutil/jsonutil"
	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
	"gorm.io/datatypes"
)

// MagicSourceRes @Resource(app="admin", name = "tools.magicSource", route = "/tools/magicSource")
func MagicSourceRes() action.Result {
	res := action.New[model.ToolsMagicSource](model.ToolsMagicSource{})

	res.Transform(func(item *model.ToolsMagicSource, index int) map[string]any {
		jsonData, _ := jsonutil.EncodeString(item.Data)
		return map[string]any{
			"id":   item.ID,
			"name": item.Name,
			"type": item.Type,
			"url":  item.Url,
			"data": jsonData,
		}
	})

	res.Validator(func(data *gjson.Result, e echo.Context) (validator.ValidatorRule, error) {
		validatorData := validator.ValidatorRule{
			"name": {Rule: "required", Message: "请填写数据源名称"},
		}
		return validatorData, nil
	})

	res.Format(func(model *model.ToolsMagicSource, data *gjson.Result, e echo.Context) error {

		jsonData := data.Get("data").String()
		if data.Get("data").Exists() && !jsonutil.IsJSON(jsonData) {
			return response.BusinessError("数据格式不正确")
		}

		model.Name = data.Get("name").String()
		model.Type = uint(data.Get("type").Uint())
		model.Url = data.Get("url").String()
		model.Data = datatypes.JSON(jsonData)
		return nil
	})

	return res.Result()
}
