package web

import (
	"fmt"
	"github.com/duxweb/go-fast/config"
	"github.com/duxweb/go-fast/global"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/tidwall/gjson"
	"io/fs"
	"log"
	"net"
	"time"
)

// @RouteGroup(app="web", name = "manage", route = "/manage")

// Location @Route(method = "GET", name = "manage.location", route = "")
func Location(ctx echo.Context) error {
	return ctx.Redirect(302, "/manage/")
}

// Index @Route(method = "GET", name = "manage.index", route = "/")
func Index(ctx echo.Context) error {

	err := fs.WalkDir(global.StaticFs, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			// 处理目录
			log.Println("dir:", path)
		} else {
			// 处理文件
			log.Println("file:", path)
		}
		return nil
	})

	file, err := global.StaticFs.ReadFile("static/web/.vite/manifest.json")
	isManifest := true
	if err != nil {
		isManifest = false
	}
	fmt.Println("file", isManifest)

	json := gjson.ParseBytes(file)

	vite := config.Load("use").GetStringMap("vite")
	port := lo.Ternary[any](vite["port"] != nil, vite["port"], 5173)
	dev := PortCheck(5173)

	if !isManifest && !dev {
		return echo.NewHTTPError(500, "未编译前端或者未开启前端服务")
	}
	vite["port"] = port
	vite["dev"] = dev

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

func PortCheck(port int) bool {
	address := fmt.Sprintf(":%d", port)
	conn, err := net.DialTimeout("tcp", address, 50*time.Millisecond)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
