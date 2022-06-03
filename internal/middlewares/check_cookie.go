package middlewares

import (
	"net/http"
	"news-be/internal/auth"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func CheckAccess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// if c.Get("access_token") == nil {
		// 	return next(c)
		// }

		claims := &auth.Claims{}
		authorizationCookie, err := c.Cookie("access_token")
		if err == nil && authorizationCookie != nil {
			jwtToken, err := jwt.ParseWithClaims(authorizationCookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_ACCESS_KEY")), nil
			})

			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					c.Response().Writer.WriteHeader(http.StatusUnauthorized)
				}
				c.Response().Writer.WriteHeader(http.StatusBadRequest)
			}

			if !jwtToken.Valid {
				c.Response().Writer.WriteHeader(http.StatusUnauthorized)
			}

		}

		// jwtTokenString := authorizationCookie
		c.Set("id", claims.Id)
		c.Set("email", claims.Email)

		return next(c)
	}
}
