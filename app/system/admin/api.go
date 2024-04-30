package admin

import (
	model "dux-project/app/system/models"
	"encoding/hex"
	"github.com/duxweb/go-fast/action"
	"github.com/duxweb/go-fast/validator"
	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/strutil"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

// ApiRes @Resource(app="admin", name = "system.api", route = "/system/api")
func ApiRes() action.Result {
	res := action.New[model.SystemApi](model.SystemApi{})
	res.Transform(func(item model.SystemApi, index int) map[string]any {
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
			"name": {Rule: "required", Message: "请填写名称"},
		}, nil
	})

	res.Format(func(model *model.SystemApi, data map[string]any, e echo.Context) error {
		model.Name = cast.ToString(data["name"])
		model.Status = cast.ToBool(data["status"])

		if model.ID == 0 {
			str, err := strutil.RandomBytes(16)
			if err != nil {
				return err
			}
			model.SecretID = cast.ToString(mathutil.RandomInt(10000000, 99999999))
			model.SecretKey = hex.EncodeToString(str)
		}
		return nil
	})

	return res.Result()
}
