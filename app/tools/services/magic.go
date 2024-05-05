package services

import (
	"encoding/json"
	"github.com/duxweb/go-fast/validator"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"strings"
)

func MagicValidator(fields []byte) validator.ValidatorRule {
	data := gjson.ParseBytes(fields)

	ruleList := validator.ValidatorRule{}
	for _, field := range data.Array() {

		rules := []map[string]any{}
		if field.Get("setting.rules").Exists() {
			_ = json.Unmarshal([]byte(field.Get("setting.rules").String()), &rules)
		}
		if field.Get("required").Bool() {
			rules = append(rules, map[string]any{
				"required": true,
				"message":  "请填写" + field.Get("label").String(),
			})
		}

		tags := []string{}
		message := ""
		for _, rule := range rules {
			for key, vo := range rule {
				if key == "message" {
					message = cast.ToString(vo)
					continue
				}
				switch key {
				case "required":
					tags = append(tags, "required")
					break
				case "date":
					tags = append(tags, "date")
					break
				case "enum":
					tags = append(tags, "enum="+strings.Join(cast.ToStringSlice(vo), "|"))
					break
				case "idcard":
					tags = append(tags, "cnIdcard")
					break
				case "length":
					tags = append(tags, "length="+cast.ToString(vo))
					break
				case "min":
					tags = append(tags, "min="+cast.ToString(vo))
					break
				case "max":
					tags = append(tags, "max="+cast.ToString(vo))
					break
				case "number":
					tags = append(tags, "numeric")
					break
				case "url":
					tags = append(tags, "url")
					break
				}
			}
		}

		if len(tags) > 0 {
			ruleList[field.Get("name").String()] = validator.ValidatorWarp{
				Message: message,
				Rule:    strings.Join(tags, ","),
			}
		}
	}

	return ruleList
}

func MagicConfig(fields []byte) []map[string]any {

	data := gjson.ParseBytes(fields)

	result := []map[string]any{}

	for _, item := range data.Array() {

		nItem := map[string]any{}
		for k, v := range item.Map() {
			nItem[k] = v.Value()
		}

		setting := map[string]any{}
		for k, v := range item.Get("setting").Map() {
			setting[k] = v.Value()
		}
		if item.Get("setting.options").Exists() && item.Get("setting.options").Type == gjson.String {
			setting["options"] = gjson.Parse(item.Get("setting.options").String()).Value()
		}
		if item.Get("setting.rules").Exists() {
			setting["rules"] = gjson.Parse(item.Get("setting.rules").String()).Value()
		}
		nItem["setting"] = setting
		if item.Get("child").Exists() && item.Get("child").IsArray() {
			nItem["child"] = ConfigFormat([]byte(item.Get("child").String()))
		}
		result = append(result, nItem)
	}

	return result
}
