package routes

import (
	"medium-be/internal/constants"
	"medium-be/internal/controller/tags"
	"medium-be/internal/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func TagPath(e *echo.Echo, tagsController *tags.TagController) {
	tag := e.Group("/tags", middleware.JWT([]byte(constants.JWT_ACCESS_KEY)))
	tag.POST("", tagsController.Create, middlewares.AdminRole)
	tag.GET("", tagsController.Read, middlewares.AdminRole)
	tag.POST("/delete", tagsController.Delete, middlewares.AdminRole)
}
