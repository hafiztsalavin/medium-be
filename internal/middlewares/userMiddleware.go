package middlewares

import (
	"fmt"
	"medium-be/internal/auth"
	"medium-be/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		user, _ := auth.ExtractToken(c)

		if user.Role != "user" {
			return c.JSON(http.StatusUnauthorized, utils.NewUnauthorizeResponse())
		}
		fmt.Println(user.Id, user.Email, user.Role)

		c.Set("id", user.Id)
		c.Set("email", user.Email)
		c.Set("role", user.Role)
		return next(c)
	}
}
