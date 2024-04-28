package web

import "github.com/labstack/echo/v4"

type Annotation struct {
	Name  string
	Route string
	Func  string
}

func Index(ctx echo.Context) error {
	return ctx.JSON(200, "dsadsad")
}
