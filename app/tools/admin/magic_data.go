package admin

import (
	model "dux-project/app/tools/models"
	"dux-project/app/tools/services"
	"encoding/json"
	"github.com/duxweb/go-fast/action"
	"github.com/duxweb/go-fast/database"
	"github.com/duxweb/go-fast/validator"
	"github.com/gookit/goutil/jsonutil"
	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// MagicDataRes @Resource(app="admin", name = "tools.data", route = "/tools/data")
func MagicDataRes() action.Result {
	res := action.New[model.ToolsMagicData](model.ToolsMagicData{})

	res.Init(func(t *action.Resources[model.ToolsMagicData], e echo.Context) error {
		e.Set("action", e.QueryParam("action"))
		magic := e.QueryParam("magic")
		info := model.ToolsMagic{}
		database.Gorm().Model(model.ToolsMagic{}).Where("name = ?", magic).First(&info)
		e.Set("magic_id", info.ID)
		e.Set("info", &info)

		if info.Type == "common" {
			t.Pagination.Status = false
		}
		if info.Type == "pages" {
			t.Pagination.Status = true
		}
		if info.Type == "tree" {
			t.Tree = true
			t.Pagination.Status = false
		}
		return nil
	})

	res.Query(func(tx *gorm.DB, e echo.Context) *gorm.DB {
		tx = tx.Where("magic_id = ?", e.Get("magic_id"))
		return tx
	})

	res.Transform(func(item *model.ToolsMagicData, index int) map[string]any {
		data := map[string]any{}
		_ = json.Unmarshal(item.Data, &data)
		data["id"] = item.ID

		results := []map[string]any{}
		for i, vo := range item.Children {
			results = append(results, res.TransformFun(vo, i))
		}
		data["children"] = results
		return data
	})

	res.Validator(func(data *gjson.Result, e echo.Context) (validator.ValidatorRule, error) {
		info := e.Get("info").(*model.ToolsMagic)
		return services.MagicValidator(info.Fields), nil
	})

	res.Format(func(tx *model.ToolsMagicData, data *gjson.Result, e echo.Context) error {
		info := e.Get("info").(*model.ToolsMagic)

		tx.MagicID = info.ID
		tx.ParentID = uint(data.Get("parent_id").Uint())
		tx.Data = datatypes.JSON(jsonutil.MustString(data.Value()))

		return nil
	})

	return res.Result()
}
