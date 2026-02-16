// src/middleware/role_middleware.go
package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func IsRole(expectedRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 1. Ambil user dari context (hasil dari JWTMiddleware)
			userToken, ok := c.Get("user").(*jwt.Token)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Token invalid"})
			}

			// 2. Ekstrak Claims
			claims, ok := userToken.Claims.(jwt.MapClaims)
			if !ok {
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal membaca data token"})
			}

			// 3. Cek Role
			role := claims["jenis"].(string)
			if role != expectedRole {
				return c.JSON(http.StatusForbidden, map[string]string{
					"message": "Akses ditolak! Anda bukan " + expectedRole,
				})
			}

			return next(c)
		}
	}
}
