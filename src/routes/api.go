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
	route.GET("/merchandises", handlers.MerchandiseHandler.GetAllMerchandiseHandler)
	route.POST("/merchandises", handlers.MerchandiseHandler.CreateMerchandiseHandler)
	route.GET("/merchandises/items/:upc", handlers.MerchandiseHandler.GetMerchandiseByUPC)
	route.PUT("/merchandises/items/:upc", handlers.MerchandiseHandler.GetMerchandiseByUPC)
	route.DELETE("/merchandises/items/:upc", handlers.MerchandiseHandler.DeleteMerchandiseByID)

	return route
}
