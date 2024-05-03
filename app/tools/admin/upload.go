package admin

import (
	"dux-project/app/tools/handlers"
	"github.com/labstack/echo/v4"
)

// @RouteGroup(app="admin", name="upload", route="/upload")

// Upload @Route(method="POST", name="upload", route="")
func Upload(ctx echo.Context) error {
	return handlers.UploadHandler("admin", ctx)
}
