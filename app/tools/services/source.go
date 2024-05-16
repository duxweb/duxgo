package services

import (
	model "dux-project/app/tools/models"
	"encoding/json"
	"github.com/duxweb/go-fast/database"
	"github.com/duxweb/go-fast/logger"
	"github.com/go-resty/resty/v2"
	"github.com/golang-module/carbon/v2"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"reflect"
	"time"
)

// StructSourceData 映射结构体
func StructSourceData(list []map[string]any) []model.ToolsMagicData {

	result := []model.ToolsMagicData{}
	for _, item := range list {

		itemModel := model.ToolsMagicData{
			ID:        cast.ToUint(item["id"]),
			ParentID:  cast.ToUint(item["parent_id"]),
			CreatedAt: carbon.Parse(cast.ToString(item["created_at"])).ToDateTimeStruct(),
			UpdatedAt: carbon.Parse(cast.ToString(item["updated_at"])).ToDateTimeStruct(),
		}

		exclude := []string{"id", "parent_id", "created_at", "updated_at"}

		data := map[string]any{}
		for k, v := range cast.ToStringMap(item) {
			if lo.IndexOf[string](exclude, k) != -1 {
				continue
			}
			data[k] = v
		}
		dataJson, _ := json.Marshal(data)
		itemModel.Data = dataJson
		if children, ok := item["children"].([]map[string]any); ok {
			itemModel.Children = StructSourceData(children)
		}
		result = append(result, itemModel)
	}
	return result
}

// MergeSourceData 合并数据与数据源
func MergeSourceData(list []map[string]any, fields []model.ToolsMagicFields, sourceMaps map[string][]map[string]any) []map[string]any {

	for i, datum := range list {
		for key, value := range datum {
			// 获取字段信息
			filed, ok := lo.Find[model.ToolsMagicFields](fields, func(item model.ToolsMagicFields) bool {
				return item.Name == key
			})

			// 字段不存在跳过
			if !ok {
				continue
			}

			if filed.Setting["source"] != nil {

				// 数据源无数据跳过
				sourceId := cast.ToString(filed.Setting["source"])
				if len(sourceMaps[sourceId]) == 0 {
					continue
				}

				// 获取id字段
				idField := cast.ToString(filed.Setting["keys_value"])
				if idField == "" {
					idField = "value"
				}

				// 合并当前id
				ids := []string{}
				if reflect.TypeOf(value).Kind() == reflect.Slice {
					ids = cast.ToStringSlice(value)
				} else {
					ids = append(ids, cast.ToString(value))
				}

				// 获取当前源数据
				filterSourceData := lo.Filter[map[string]any](sourceMaps[sourceId], func(item map[string]any, index int) bool {
					// 数据源值
					val := cast.ToString(item[idField])
					return lo.IndexOf[string](ids, val) != -1
				})

				// 重设值数据
				if len(filterSourceData) > 0 {
					if reflect.TypeOf(value).Kind() == reflect.Slice {
						value = filterSourceData
					} else {
						value = filterSourceData[0]
					}
				}

			}

			// 处理子节点
			if len(filed.Child) > 0 {
				if valueData, ok := value.([]any); ok {
					child := make([]map[string]any, 0)
					for _, m := range valueData {
						child = append(child, cast.ToStringMap(m))
					}
					value = MergeSourceData(child, filed.Child, sourceMaps)
				}
			}

			datum[key] = value
		}

		// 处理树形
		if children, ok := datum["children"].([]map[string]any); ok {
			datum["children"] = MergeSourceData(children, fields, sourceMaps)
		}

		list[i] = datum
	}

	return list
}

// GetSourceMapsData 获取映射数据
func GetSourceMapsData(data []map[string]any, fields []model.ToolsMagicFields, c echo.Context) map[string][]map[string]any {
	sources := map[string][]string{}
	getSourceMapsIds(data, fields, sources)

	result := map[string][]map[string]any{}

	for sourceId, ids := range sources {
		sourceData, err := GetSourceData(sourceId, ids, c)
		if err != nil {
			logger.Log().Error(err.Error())
		}
		result[sourceId] = sourceData
	}
	return result
}

// GetModelData 解压模型数据
func GetModelData(models []model.ToolsMagicData) []map[string]any {
	result := []map[string]any{}
	for _, item := range models {
		data := map[string]any{}
		_ = json.Unmarshal(item.Data, &data)
		data["id"] = item.ID
		data["parent_id"] = item.ParentID
		data["created_at"] = item.CreatedAt.Format("2006-01-02 15:04:05")
		data["updated_at"] = item.UpdatedAt.Format("2006-01-02 15:04:05")
		if len(item.Children) > 0 {
			data["children"] = GetModelData(item.Children)
		}
		result = append(result, data)
	}
	return result
}

// 获取数据源->数据id映射
func getSourceMapsIds(list []map[string]any, fields []model.ToolsMagicFields, source map[string][]string) {

	for _, datum := range list {
		for key, value := range datum {
			// 获取字段信息
			filed, ok := lo.Find[model.ToolsMagicFields](fields, func(item model.ToolsMagicFields) bool {
				return item.Name == key
			})
			if !ok {
				continue
			}
			// 处理源数据
			if filed.Setting["source"] != nil {
				sourceId := cast.ToString(filed.Setting["source"])
				if source[sourceId] == nil {
					source[sourceId] = make([]string, 0)
				}
				if reflect.TypeOf(value).Kind() == reflect.Slice {
					source[sourceId] = append(source[sourceId], cast.ToStringSlice(value)...)
				} else {
					source[sourceId] = append(source[sourceId], cast.ToString(value))
				}
			}
			// 处理子节点
			if len(filed.Child) > 0 {
				if valueData, ok := value.([]any); ok {
					child := make([]map[string]any, 0)
					for _, m := range valueData {
						child = append(child, cast.ToStringMap(m))
					}
					getSourceMapsIds(child, filed.Child, source)
				}
			}
		}

		// 处理树形
		if children, ok := datum["children"].([]map[string]any); ok {
			getSourceMapsIds(children, fields, source)
		}
	}
}

// GetSourceData 获取源数据
func GetSourceData(id string, ids []string, c echo.Context) ([]map[string]any, error) {

	info := model.ToolsMagicSource{}
	err := database.Gorm().Model(model.ToolsMagicSource{}).First(&info, id).Error
	if err != nil {
		return nil, err
	}

	data := []map[string]any{}

	if info.Type == 0 {
		_ = json.Unmarshal(info.Data, &data)

	}

	if info.Type == 1 {
		client := resty.New()
		defer client.Clone()
		get, err := client.
			SetTimeout(10*time.Second).
			R().
			SetHeader("Authorization", c.Request().Header.Get("Authorization")).
			Get("http://" + c.Request().Host + "/admin/" + info.Url)
		if err != nil {
			return nil, err
		}

		resData := gjson.ParseBytes(get.Body()).Get("data")

		if !resData.Exists() {
			return data, nil
		}

		err = json.Unmarshal([]byte(resData.String()), &data)
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}
