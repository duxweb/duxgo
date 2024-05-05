package main

import (
	"dux-project/app/home"
	"dux-project/app/system"
	"dux-project/app/tools"
	"dux-project/runtime"
	"embed"
	dux "github.com/duxweb/go-fast"
)

//go:embed all:static
var StaticFs embed.FS

func main() {

	app := dux.New()

	app.SetStaticFs(StaticFs)
	app.SetAnnotations(runtime.Annotations)
	app.RegisterApp(home.App, system.App, tools.App)
	app.Run()
}
