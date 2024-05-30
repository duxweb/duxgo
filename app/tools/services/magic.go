package services

import (
	model "dux-project/app/tools/models"
	"encoding/json"
	"github.com/duxweb/go-fast/i18n"
	"github.com/duxweb/go-fast/validator"
	"github.com/spf13/cast"
	"strings"
)

func MagicValidator(fields []model.ToolsMagicFields) validator.ValidatorRule {

	ruleList := validator.ValidatorRule{}
	for _, field := range fields {

		rules := []map[string]any{}
		if field.Setting["rules"] != nil {
			_ = json.Unmarshal([]byte(field.Setting["rules"].(string)), &rules)
		}
		if field.Required {
			rules = append(rules, map[string]any{
				"required": true,
				"message": i18n.Trans.GetData("tools.magic.validator.required", map[string]any{
					"label": field.Label,
				}),
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
			ruleList[field.Name] = validator.ValidatorWarp{
				Message: message,
				Rule:    strings.Join(tags, ","),
			}
		}
	}

	return ruleList
}

func MagicConfig(fields []model.ToolsMagicFields) []model.ToolsMagicFields {
	for i, item := range fields {
		if item.Setting["rules"] != nil {
			rule := []map[string]any{}
			_ = json.Unmarshal([]byte(item.Setting["rules"].(string)), &rule)
			item.Setting["rules"] = rule
		}
		if len(item.Child) > 0 {
			item.Child = MagicConfig(item.Child)
		}
		fields[i] = item
	}
	return fields
}
