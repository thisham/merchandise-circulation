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

	// users
	route.GET("/users", handlers.UserHandler.GetAllUsersHandler)
	route.GET("/users/:id", handlers.UserHandler.GetUserByIDHandler)

	// auth
	route.POST("/login", handlers.UserHandler.LoginHandler)
	route.POST("/register", handlers.UserHandler.RegisterHandler)
	route.GET("/logout", handlers.UserHandler.LogoutHandler)

	return route
}
