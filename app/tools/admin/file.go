package admin

import (
	model "dux-project/app/tools/models"
	"github.com/duxweb/go-fast/action"
	"github.com/duxweb/go-fast/helper"
	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
)

// FileRes @Resource(app="admin", name = "tools.file", route = "/tools/file")
func FileRes() action.Result {
	res := action.New[model.ToolsFile](model.ToolsFile{})

	res.ActionStore = false
	res.ActionEdit = false
	res.ActionCreate = false

	res.QueryMany(func(tx *gorm.DB, params *gjson.Result, e echo.Context) *gorm.DB {
		tx = tx.Preload("Dir").Where("has_type", "admin").Order("id desc")

		if params.Get("dir_id").Exists() {
			tx = tx.Where("dir_id = ?", params.Get("dir_id").Int())
		}
		return tx
	})

	res.Transform(func(item *model.ToolsFile, index int) map[string]any {
		return map[string]any{
			"id":       item.ID,
			"dir_name": item.Dir.Name,
			"name":     item.Name,
			"ext":      item.Ext,
			"url":      item.Url,
			"size":     helper.HumanFileSize(int64(item.Size)),
			"mime":     item.Mime,
			"driver":   item.Driver,
			"time":     item.CreatedAt,
		}
	})

	return res.Result()
}
