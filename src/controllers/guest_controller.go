package controllers

import (
	"lppm/src/database"
	"lppm/src/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetBook(c echo.Context) error {
	var books []models.Buku

	result := database.DB.Find(&books)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Gagal mengambil data buku",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil mengambil semua data buku",
		"data":    books,
	})
}
