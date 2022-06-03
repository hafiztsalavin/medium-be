package routes

import (
	"news-be/internal/controller/news"

	"github.com/labstack/echo"
)

func UserPath(e *echo.Group) {
	// e.GET("/auth/me", news.ReadOne, middleware.CheckAccess)

	e.GET("", news.ReadOne())
}
