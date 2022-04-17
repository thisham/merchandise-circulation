package routes

import (
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	route := echo.New()

	// main
	route.GET("/", func(ctx echo.Context) error {
		return ctx.HTML(200, "Server runs perfectly!")
	})

	return route
}
