package routes

import (
	"medium-be/internal/constants"
	"medium-be/internal/controller/posts"
	"medium-be/internal/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewsPath(e *echo.Echo, postController *posts.PostController) {
	// e.GET("/auth/me", news.ReadOne, middleware.CheckAccess)
	post := e.Group("/post", middleware.JWT([]byte(constants.JWT_ACCESS_KEY)), middlewares.UserRole)
	post.POST("", postController.CreatePost)
	post.PUT("/update/:id", postController.UpdatePost)
}
