package middlewares

import (
	"bank-sampah/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user") // Ambil user dari konteks
		if user == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
		}

		if u, ok := user.(*models.User); ok {
			if u.Role != "admin" {
				return c.JSON(http.StatusForbidden, map[string]string{"message": "Access denied"})
			}
		} else {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
		}

		return next(c)
	}
}
