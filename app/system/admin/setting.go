package admin

import (
	"github.com/duxweb/go-fast/auth"
	"github.com/duxweb/go-fast/menu"
	"github.com/duxweb/go-fast/response"
	"github.com/labstack/echo/v4"
)

// @RouteGroup(app="admin", name="auth", route="")

// Menu @Route(method="GET", name="menu", route="/menu")
func Menu(ctx echo.Context) error {
	id := auth.NewService("admin", ctx).ID()
	if id == "" {
		return response.BusinessError("Expired or incorrect token")
	}

	return response.Send(ctx, response.Data{
		Data: menu.Get("admin").Get(),
	})
}
