package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func TestAdmin(c echo.Context) error {
	fmt.Println("Ini Adalah Admin")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Restriksi Berhasil Admin",
	})
}
