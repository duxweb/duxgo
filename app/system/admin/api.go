package admin

import (
	model "dux-project/app/system/models"
	"github.com/duxweb/go-fast/action"
	"github.com/duxweb/go-fast/validator"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

// ApiRes @Resource(app="admin", name = "system.api", route = "/system/api")
func ApiRes() action.Result {
	type data struct {
		Name string `json:"name"`
	}
	res := action.New[*model.SystemApi](&model.SystemApi{})
	res.Transform(func(item *model.SystemApi, index int) map[string]any {
		return map[string]any{
			"id":         item.ID,
			"name":       item.Name,
			"secret_id":  item.SecretID,
			"secret_key": item.SecretKey,
			"status":     cast.ToBool(item.Status),
		}
	})

	res.Validator(func(data map[string]any, e echo.Context) (validator.ValidatorRule, error) {
		return validator.ValidatorRule{
			"name":      {Rule: "required", Message: "请填写名称"},
			"secret_id": {Rule: "required", Message: "请填写secret_id"},
		}, nil
	})

	return res.Result()
}
