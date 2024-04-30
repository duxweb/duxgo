package system

import (
	"embed"
	"github.com/duxweb/go-fast/app"
	"github.com/duxweb/go-fast/menu"
	"github.com/duxweb/go-fast/route"
	"github.com/labstack/echo/v4"
	"net/http"
)

var config = struct {
}{}

func App() {
	app.Register(&app.Config{
		Name:     "system",
		Config:   &config,
		Init:     Init,
		Register: Register,
	})
}

//go:embed views/*.gohtml
var ViewsFs embed.FS

func Init(t *app.Dux) {
	t.RegisterTplFS("manage", ViewsFs)

	route.Set("admin", route.New("/admin"))
	menu.Set("admin", menu.New())

}

func Register(t *app.Dux) {
	group := route.Get("web")
	group.Get("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}, "web.home")

}
