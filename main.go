package main

import (
	"dux-project/app/home"
	"dux-project/app/system"
	"dux-project/runtime"
	"embed"
	dux "github.com/duxweb/go-fast"
)

//go:embed views/*.gohtml
var ViewsFs embed.FS

//go:embed static/*
var StaticFs embed.FS

func main() {

	app := dux.New()
	app.RegisterTplFS("manage", ViewsFs)
	app.RegisterStaticFs(StaticFs)
	app.RegisterAnnotations(runtime.Annotations)
	app.RegisterApp(home.App, system.App)
	app.Run()
}
