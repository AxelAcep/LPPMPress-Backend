package routes

import (
	"lppm/src/controllers"
	"lppm/src/middleware"

	"github.com/labstack/echo/v4"
)

func AdminRoutes(e *echo.Echo) {

	adminGroup := e.Group("/admin")
	adminGroup.Use(middleware.JWTMiddleware()) //
	adminGroup.Use(middleware.IsRole("admin"))

	adminGroup.GET("/test", controllers.TestAdmin)
}
