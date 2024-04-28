package admin

import (
	model "dux-project/app/system/models"
	"github.com/duxweb/go-fast/action"
)

// ApiRes @Resource(app="web", name = "system.api", route = "/system")
func ApiRes() action.Result {
	res := action.New[*model.LogApi](&model.LogApi{})
	return res.Result()
}
