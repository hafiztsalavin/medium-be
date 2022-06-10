package middlewares

import (
	"medium-be/internal/auth"
	"medium-be/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdminRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		user, _ := auth.ExtractToken(c)

		if user.Role != "admin" {
			return c.JSON(http.StatusUnauthorized, utils.NewUnauthorizeResponse())
		}

		c.Set("id", user.Id)
		c.Set("email", user.Email)
		c.Set("role", user.Role)
		return next(c)
	}
}
