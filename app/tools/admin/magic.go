package admin

import (
	model "dux-project/app/tools/models"
	"dux-project/app/tools/services"
	"encoding/json"
	"github.com/duxweb/go-fast/action"
	"github.com/duxweb/go-fast/database"
	"github.com/duxweb/go-fast/response"
	"github.com/duxweb/go-fast/validator"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/tidwall/gjson"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"regexp"
)

// MagicRes @Resource(app="admin", name = "tools.magic", route = "/tools/magic")
func MagicRes() action.Result {
	res := action.New[model.ToolsMagic](model.ToolsMagic{})

	res.QueryMany(func(tx *gorm.DB, params *gjson.Result, e echo.Context) *gorm.DB {
		tx = tx.Preload("Group").Order("id desc")

		if params.Get("label").Exists() {
			tx = tx.Where("label = ?", params.Get("label").String())
		}
		if params.Get("group").Exists() {
			tx = tx.Where("group_id = ?", params.Get("group").String())
		}
		return tx
	})

	res.Transform(func(item *model.ToolsMagic, index int) map[string]any {
		return map[string]any{
			"id":          item.ID,
			"group_id":    item.GroupID,
			"name":        item.Name,
			"group_icon":  item.Group.Icon,
			"group_label": item.Group.Label,
			"label":       item.Label,
			"type":        item.Type,
			"page":        lo.Ternary[uint](*item.Page, 1, 0),
			"external":    item.External,
			"tree_label":  item.TreeLabel,
			"inline":      lo.Ternary[uint](*item.Inline, 1, 0),
			"fields":      item.Fields,
		}
	})

	res.Validator(func(data *gjson.Result, e echo.Context) (validator.ValidatorRule, error) {

		pattern := "^[a-zA-Z][\\w]*[a-zA-Z0-9]$"

		validatorData := validator.ValidatorRule{
			"fields":   {Rule: "required", Message: "请填写魔方描述"},
			"group_id": {Rule: "required", Message: "请选择魔方组"},
			"name":     {Rule: "fieldName", Message: "表名填写错误"},
		}

		fields := data.Get("fields").Array()

		for _, field := range fields {
			if !field.Get("name").Exists() && !field.Get("label").Exists() {
				return nil, response.BusinessError("字段信息不完整")
			}
			_, err := regexp.MatchString(pattern, field.Get("name").String())
			if err != nil {
				return nil, response.BusinessError("字段名不规范")
			}
		}

		return validatorData, nil
	})

	res.Format(func(model *model.ToolsMagic, data *gjson.Result, e echo.Context) error {
		page := data.Get("page").Bool()
		inline := data.Get("inline").Bool()
		model.GroupID = uint(data.Get("group_id").Uint())
		model.Name = data.Get("name").String()
		model.Label = data.Get("label").String()
		model.Type = data.Get("type").String()
		model.TreeLabel = data.Get("tree_label").String()
		model.Page = &page
		model.Inline = &inline
		model.External = datatypes.JSON(data.Get("external").Raw)
		model.Fields = datatypes.JSON(data.Get("fields").Raw)

		return nil
	})

	return res.Result()
}

// MagicConfig @Action(method = "GET", name = "config", route = "/config")
func MagicConfig(c echo.Context) error {
	info := model.ToolsMagic{}
	err := database.Gorm().Model(model.ToolsMagic{}).Where("name = ?", c.QueryParam("magic")).First(&info).Error
	if err != nil {
		return response.BusinessError(err.Error())
	}

	var fields []model.ToolsMagicFields
	_ = json.Unmarshal(info.Fields, &fields)

	fieldsJson, _ := json.Marshal(services.MagicConfig(fields))
	info.Fields = fieldsJson
	return response.Send(c, response.Data{
		Data: info,
	}, 200)
}

// MagicSource @Action(method = "GET", name = "source", route = "/source")
func MagicSource(c echo.Context) error {
	var list []model.ToolsMagicSource
	database.Gorm().Model(model.ToolsMagicSource{}).Find(&list)

	data := []map[string]any{}
	for _, source := range list {
		data = append(data, map[string]any{
			"label": source.Name,
			"value": source.ID,
		})
	}
	return response.Send(c, response.Data{
		Data: data,
	}, 200)
}

// MagicSourceData @Action(method = "GET", name = "sourceData", route = "/sourceData")
func MagicSourceData(c echo.Context) error {
	name := c.QueryParam("name")
	data, err := services.GetSourceData(name, []string{}, c)
	if err != nil {
		return response.BusinessError(err.Error())
	}
	return response.Send(c, response.Data{
		Data: data,
	}, 200)
}
