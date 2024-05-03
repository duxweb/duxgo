package tools

import (
	"embed"
	"github.com/duxweb/go-fast/app"
)

var config = struct {
}{}

func App() {
	app.Register(&app.Config{
		Name:   "tools",
		Config: &config,
		Init:   Init,
	})
}

//go:embed langs/*.yaml
var LangFs embed.FS

func Init(t *app.Dux) {
	t.RegisterLangFS(LangFs)
}
