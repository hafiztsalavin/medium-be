package routes

import (
	"medium-be/internal/constants"
	"medium-be/internal/controller/tags"
	"medium-be/internal/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func TagPath(e *echo.Echo, tagsController *tags.TagController) {
	tag := e.Group("/tags", middleware.JWT([]byte(constants.JWT_ACCESS_KEY)), middlewares.AdminRole)
	tag.POST("", tagsController.CreateTag)
	tag.POST("/delete", tagsController.DeleteTag) // params id

	tag.GET("/list", tagsController.ReadAllTag)
	tag.GET("", tagsController.ReadTag) // params id

	tag.POST("", tagsController.UpdateTag) // params id
}
