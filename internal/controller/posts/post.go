package posts

import (
	"medium-be/internal/entity"
	"medium-be/internal/repository/posts"
	"medium-be/internal/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PostController struct {
	Repository posts.PostInterface
}

// init
func NewPostController(postInterface posts.PostInterface) *PostController {
	return &PostController{Repository: postInterface}
}

func (pc PostController) CreatePost(c echo.Context) error {
	idUser := c.Get("id").(uint)

	var postRequest PostRequest

	c.Bind(&postRequest)
	if err := c.Validate(&postRequest); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestResponse())
	}

	newPost := entity.Posts{
		UserID: idUser,
		Title:  postRequest.Title,
		Body:   postRequest.Body,
	}

	err := pc.Repository.CreatePost(newPost, postRequest.Tags)
	if err != nil {
		return c.JSON(http.StatusConflict, utils.ErrorResponse(409, err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewSuccessOperationResponse())
}

func (pc PostController) UpdatePost(c echo.Context) error {
	idUser := c.Get("id").(uint)
	idPost, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestResponse())
	}

	var postRequest PostRequest
	c.Bind(&postRequest)
	if err := c.Validate(&postRequest); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestResponse())
	}

	editPost := entity.Posts{
		UserID: idUser,
		Title:  postRequest.Title,
		Body:   postRequest.Body,
	}

	err = pc.Repository.EditPost(idPost, editPost, postRequest.Tags)
	if err != nil {
		return c.JSON(http.StatusConflict, utils.ErrorResponse(409, err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewSuccessOperationResponse())
}

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
