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

	groups := models.GetMagicMenu()

	fmt.Println("groups", groups)

	a := menus.Add(&menu.MenuData{
		Name:  "data",
		Label: "data",
		Icon:  "i-tabler:database",
		Meta: map[string]any{
			"sort": 100,
		},
	})
	for _, group := range groups {
		groupName := "tools.data." + group["name"].(string)
		g := a.Group(groupName, group["label"].(string), group["icon"].(string))
		for _, item := range group["children"].([]map[string]any) {
			name := item["name"].(string)
			g.Item(groupName+"."+name, item["label"].(string), "data/"+name, 0)
		}
	}

}
