package tools

import (
	"dux-project/app/tools/models"
	"embed"
	"fmt"
	"github.com/duxweb/go-fast/app"
	"github.com/duxweb/go-fast/menu"
)

var config = struct {
}{}

func App() {
	app.Register(&app.Config{
		Name:     "tools",
		Config:   &config,
		Init:     Init,
		Register: Register,
	})
}

//go:embed langs/*.yaml
var LangFs embed.FS

func Init(t *app.Dux) {
	t.RegisterLangFS(LangFs)
}

func Register(t *app.Dux) {
	fmt.Println("xxx", "xxx")
	menus := menu.Get("admin")
	models.GetMagicMenu(menus)

}
