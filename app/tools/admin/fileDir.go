package admin

import (
	model "dux-project/app/tools/models"
	"github.com/duxweb/go-fast/action"
	"github.com/duxweb/go-fast/validator"
	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
)

// FileDirRes @Resource(app="admin", name = "tools.fileDir", route = "/tools/fileDir")
func FileDirRes() action.Result {
	res := action.New[model.ToolsFileDir](model.ToolsFileDir{})

	res.QueryMany(func(tx *gorm.DB, params *gjson.Result, e echo.Context) *gorm.DB {
		tx = tx.Where("has_type", "admin")
		return tx
	})

	res.Transform(func(item *model.ToolsFileDir, index int) map[string]any {
		return map[string]any{
			"id":   item.ID,
			"name": item.Name,
		}
	})

	res.Validator(func(data *gjson.Result, e echo.Context) (validator.ValidatorRule, error) {
		return validator.ValidatorRule{
			"name": {Rule: "required", LangMessage: "tools.fileDir.validator.name"},
		}, nil
	})

	res.Format(func(model *model.ToolsFileDir, data *gjson.Result, e echo.Context) error {
		model.Name = data.Get("name").String()
		model.HasType = "admin"
		return nil
	})

	return res.Result()
}
