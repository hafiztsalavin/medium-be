package posts

import (
	"medium-be/internal/repository/posts"
	"medium-be/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PostController struct {
	Repository posts.PostInterface
}

// init
func NewPostController(postInterface posts.PostInterface) *PostController {
	return &PostController{Repository: postInterface}
}

// func (pc PostController) CreatePost(c echo.Context) error {

// }

func (pc PostController) UserDetails(c echo.Context) error {
	idUser := c.Get("id").(uint)
	emailUser := c.Get("email").(string)
	roleUser := c.Get("role").(string)

	response := UserResponse{
		ID:    idUser,
		Email: emailUser,
		Role:  roleUser,
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(response))
}
