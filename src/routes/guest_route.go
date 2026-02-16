// src/routes/user_route.go
package routes

import (
	"lppm/src/controllers"

	"github.com/labstack/echo/v4"
)

func GuestRoutes(e *echo.Echo) {
	e.GET("/barang", controllers.GetBook)
}
