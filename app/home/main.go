package home

import (
	"github.com/duxweb/go-fast/app"
	"github.com/duxweb/go-fast/route"
	"github.com/labstack/echo/v4"
	"net/http"
)

var config = struct {
}{}

func App() {
	app.Register(&app.Config{
		Name:     "home",
		Config:   &config,
		Init:     Init,
		Register: Register,
	})
}

func Init() {
	route.Set("web", route.New(""))
}

func Register() {
	group := route.Get("web")
	group.Get("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}, "web.home")

}
