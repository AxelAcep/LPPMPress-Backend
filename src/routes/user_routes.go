// src/routes/user_route.go
package routes

import (
	"lppm/src/controllers"
	"lppm/src/middleware"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	pelangganGroup := e.Group("/user")

	pelangganGroup.POST("/registerPelanggan", controllers.RegisterUser, middleware.LoginRateLimiter())
	pelangganGroup.POST("/login", controllers.LoginUser, middleware.LoginRateLimiter())
	pelangganGroup.POST("/registerAdmin", controllers.RegisterAdmin, middleware.JWTMiddleware(), middleware.IsRole("admin"))

}
