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
	e.GET("/posts", postController.AllPostPublish)
	e.GET("/posts/:id", postController.PostByIDPost)

	post := e.Group("/post", middleware.JWT([]byte(constants.JWT_ACCESS_KEY)), middlewares.UserRole)
	post.POST("", postController.CreatePost)

	post.PUT("/:id", postController.UpdatePost)
	post.PUT("/:id/publish", postController.PublishPost)
	post.PUT("/:id/delete", postController.DeletePost)

	post.GET("/:id", postController.ReadPost)
	post.GET("/list", postController.ReadAllPost) // params status
}
