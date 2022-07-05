package routes

import (
	"medium-be/internal/constants"
	"medium-be/internal/controller/tags"
	"medium-be/internal/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func TagPath(e *echo.Echo, tagsController *tags.TagController) {
	e.GET("/tags", tagsController.ReadAllTag)
	e.GET("/tags/:id", tagsController.ReadTag) // params id

	tag := e.Group("/tag", middleware.JWT([]byte(constants.JWT_ACCESS_KEY)), middlewares.AdminRole)
	tag.POST("/create", tagsController.CreateTag)
	tag.POST("/delete", tagsController.DeleteTag) // params id
	tag.POST("/edit", tagsController.UpdateTag)   // params id
}
