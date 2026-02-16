package routes

import (
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	GuestRoutes(e)
	UserRoutes(e)
	AdminRoutes(e)
	PelangganRoutes(e)

	return e
}
