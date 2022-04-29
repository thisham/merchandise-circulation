package routes

import (
	"merchandise-circulation-api/src/factories"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	route := echo.New()
	handlers := factories.Init()

	// main
	route.GET("/", func(ctx echo.Context) error {
		return ctx.HTML(200, "Server runs perfectly!")
	})

	// merchandises
	route.GET("/merchandises", handlers.MerchandiseHandler.GetAllMerchandisesHandler)
	route.GET("/merchandises/:upc", handlers.MerchandiseHandler.GetMerchandiseByUPCHandler)

	// merchandise types
	route.GET("/merchandise_types", handlers.MerchandiseTypeHandler.GetAllMerchandiseTypesHandler)

	return route
}
