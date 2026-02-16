package routes

import (
	"lppm/src/controllers"
	"lppm/src/middleware" // Pastikan import middleware sudah benar

	"github.com/labstack/echo/v4"
)

func PelangganRoutes(e *echo.Echo) {

	pelangganGroup := e.Group("/pelanggan")
	pelangganGroup.Use(middleware.JWTMiddleware())
	pelangganGroup.Use(middleware.IsRole("pelanggan"))

	pelangganGroup.GET("/test", controllers.TestPelanggan)
}
