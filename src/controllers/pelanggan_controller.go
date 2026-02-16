package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func TestPelanggan(c echo.Context) error {
	fmt.Println("Ini Adalah pelanggan")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Restriksi Berhasil Pelanggan",
	})
}
