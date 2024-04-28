package web

import (
	"github.com/duxweb/go-fast/config"
	"github.com/duxweb/go-fast/global"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/tidwall/gjson"
)

// @RouteGroup(app="web", name = "manage", route = "/manage")

// Location @Route(method = "GET", name = "manage.location", route = "")
func Location(ctx echo.Context) error {
	return ctx.Redirect(302, "/manage/")
}

// Index @Route(method = "GET", name = "manage.index", route = "/")
func Index(ctx echo.Context) error {

	file, err := global.StaticFs.ReadFile("web/.vite/manifest.json")
	if err != nil {
		//return err
	}
	json := gjson.ParseBytes(file)

	vite := config.Load("use").GetStringMap("vite")
	vite["port"] = lo.Ternary[any](vite["port"] != nil, vite["port"], 5173)
	vite["dev"] = lo.Ternary[any](vite["dev"] != nil, vite["dev"], false)

	manage := config.Load("use").GetStringMap("manage")

	manage["indexName"] = lo.Ternary[any](manage["indexName"] != nil, manage["indexName"], "system")
	manage["sideType"] = lo.Ternary[any](manage["sideType"] != nil, manage["sideType"], "app")

	data := map[string]any{
		"title": config.Load("use").GetString("app.name"),
		"lang":  config.Load("use").GetString("app.lang"),
		"vite":  vite,
		"manifest": map[string]any{
			"js":  json.Get("src\\/index\\.tsx").Get("file").String(),
			"css": json.Get("style\\.css").Get("file").String(),
		},
		"manage": manage,
	}

	ctx.Set("tpl", "manage")

	return ctx.Render(200, "manage.gohtml", data)
}
