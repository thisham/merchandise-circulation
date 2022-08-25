package routes

import (
	"merchandise-circulation-api/src/factories"
	"merchandise-circulation-api/src/middlewares"

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
	route.GET("/merchandises", handlers.MerchandiseHandler.GetAllMerchandisesHandler,
		middlewares.VerifyAuthentication())
	route.GET("/merchandises/:upc", handlers.MerchandiseHandler.GetMerchandiseByUPCHandler,
		middlewares.VerifyAuthentication())

	// merchandise types
	route.GET("/merchandise_types", handlers.MerchandiseTypeHandler.GetAllMerchandiseTypesHandler,
		middlewares.VerifyAuthentication())

	// users
	route.GET("/users", handlers.UserHandler.GetAllUsersHandler,
		middlewares.VerifyAuthentication())
	route.GET("/users/:id", handlers.UserHandler.GetUserByIDHandler,
		middlewares.VerifyAuthentication())

	// auth
	route.POST("/login", handlers.UserHandler.LoginHandler)
	route.POST("/register", handlers.UserHandler.RegisterHandler)
	route.GET("/logout", handlers.UserHandler.LogoutHandler,
		middlewares.VerifyAuthentication())

	return route
}
