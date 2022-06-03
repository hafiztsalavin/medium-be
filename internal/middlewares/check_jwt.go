package middlewares

import (
	"news-be/internal/auth"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//SetJwtMiddlewares create middleware that will be used in router
func CheckJWT(g *echo.Group) {
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:      &auth.Claims{},
		SigningKey:  []byte(os.Getenv("JWT_ACCESS_KEY")),
		TokenLookup: "cookie:access_token",
	}))
}
