package system

import (
	model "dux-project/app/system/models"
	"embed"
	"github.com/duxweb/go-fast/app"
	"github.com/duxweb/go-fast/database"
	"github.com/duxweb/go-fast/middleware"
	"github.com/duxweb/go-fast/resources"
	"github.com/duxweb/go-fast/route"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
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

//go:embed langs/*.yaml
var LangFs embed.FS

func Init(t *app.Dux) {
	t.RegisterLangFS(LangFs)
	t.RegisterTplFS("manage", ViewsFs)

	resources.Set("admin", resources.
		New("admin", "/admin").
		SetPermission(func(id string) (map[string]bool, error) {
			user := model.SystemUser{}
			err := database.Gorm().Model(model.SystemUser{}).Preload(clause.Associations).Where("id = ?", id).First(&user).Error
			if err != nil {
				return nil, err
			}
			return user.GetPermission(), nil
		}).
		SetOperate(true),
	)
}

func Register(t *app.Dux) {
	group := route.Get("web").Group("", "web.test", middleware.VisitorMiddleware)
	group.Get("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}, "web.home")

}
