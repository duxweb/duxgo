package admin

import (
	model "dux-project/app/tools/models"
	"fmt"
	"github.com/duxweb/go-fast/action"
	"github.com/duxweb/go-fast/database"
	"github.com/duxweb/go-fast/helper"
	"github.com/duxweb/go-fast/models"
	"github.com/duxweb/go-fast/response"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
)

// AreaRes @Resource(app="admin", name = "tools.area", route = "/tools/area")
func AreaRes() action.Result {
	res := action.New[model.ToolsArea](model.ToolsArea{})

	res.ActionStore = false
	res.ActionEdit = false
	res.ActionCreate = false

	res.QueryMany(func(tx *gorm.DB, params *gjson.Result, e echo.Context) *gorm.DB {
		if params.Get("level").Exists() {
			tx = tx.Where("level = ?", params.Get("level").Int())
		}
		if params.Get("name").Exists() {
			name := params.Get("name").String()
			area := model.ToolsArea{}
			database.Gorm().Model(model.ToolsArea{}).Where("name = ?", name).First(&area)
			tx = tx.Where("parent_code = ?", area.Code)
		}
		return tx
	})

	res.Transform(func(item *model.ToolsArea, index int) map[string]any {
		return map[string]any{
			"id":    item.ID,
			"code":  item.Code,
			"name":  item.Name,
			"level": item.Level,
		}
	})

	return res.Result()
}

// AreaImport @Action(method = "POST", name = "import", route = "")
func AreaImport(c echo.Context) error {

	dataJson, err := helper.Body(c)
	if err != nil {
		return err
	}

	if !dataJson.Get("file.0.url").Exists() {
		return response.BusinessLangError("tools.area.validator.file")
	}

	url := dataJson.Get("file.0.url").String()

	rows, err := helper.ExcelImport(url)
	if err != nil {
		return err
	}

	data := map[string]map[string]any{}
	for _, row := range rows {

		if _, ok := data[row[1]+":1"]; !ok {
			data[row[1]+":1"] = map[string]any{
				"parent_code": 0,
				"code":        row[1],
				"name":        row[0],
				"level":       1,
				"leaf":        true,
			}
		}

		if _, ok := data[row[3]+":2"]; !ok {
			data[row[3]+":2"] = map[string]any{
				"parent_code": row[1],
				"code":        row[3],
				"name":        row[2],
				"level":       2,
				"leaf":        true,
			}
			data[row[1]+":1"]["leaf"] = false
		}

		if _, ok := data[row[5]+":3"]; !ok {
			data[row[5]+":3"] = map[string]any{
				"parent_code": row[3],
				"code":        row[5],
				"name":        row[4],
				"level":       3,
				"leaf":        true,
			}
			data[row[3]+":2"]["leaf"] = false
		}

		if _, ok := data[row[7]+":4"]; !ok {
			data[row[7]+":4"] = map[string]any{
				"parent_code": row[5],
				"code":        row[7],
				"name":        row[6],
				"level":       4,
				"leaf":        true,
			}
			data[row[5]+":3"]["leaf"] = false
		}
	}

	list := lo.Values[string, map[string]any](data)

	err = database.Gorm().Exec(fmt.Sprintf("TRUNCATE TABLE %s", models.GetTableName(database.Gorm(), model.ToolsArea{}))).Error
	if err != nil {
		return err
	}

	err = database.Gorm().Transaction(func(tx *gorm.DB) error {
		err = tx.Model(model.ToolsArea{}).CreateInBatches(list, 1000).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return response.Send(c, response.Data{
		Message: "导入成功",
	}, 200)
}
