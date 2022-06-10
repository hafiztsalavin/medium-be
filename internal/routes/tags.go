package routes

import (
	"medium-be/internal/controller/tags"

	"github.com/labstack/echo/v4"
)

func TagPath(e *echo.Echo, tagsController *tags.TagController) {

	// e.GET("/news", , middleware.JWT([]byte(constants.JWT_ACCESS_KEY)), middlewares.UserRole)
}
