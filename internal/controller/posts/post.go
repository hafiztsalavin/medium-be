package posts

import (
	"medium-be/internal/entity"
	"medium-be/internal/repository/posts"
	"medium-be/internal/utils"
	"net/http"
	"strconv"
	"strings"

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

func (pc PostController) ReadPost(c echo.Context) error {
	idPost, err := strconv.Atoi(c.Param("id"))
	idUser := c.Get("id").(uint)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestResponse())
	}

	postDB, err := pc.Repository.ReadPost(idPost)
	if err != nil || postDB.UserID != idUser {
		return c.JSON(http.StatusNotFound, utils.NewNotFoundResponse())
	}

	tags := []string{}
	for _, tag := range postDB.Tags {
		tags = append(tags, tag.Name)
	}

	response := PostResponse{}
	response = PostResponse{
		ID:     int(postDB.ID),
		Title:  postDB.Title,
		Body:   postDB.Body,
		Status: postDB.Status,
		Tags:   tags,
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse(response))
}

func (pc PostController) ReadAllPost(c echo.Context) error {
	idUser := c.Get("id").(uint)
	statusPost := c.QueryParam("status")

	postDB, err := pc.Repository.ReadAllPost(int(idUser), statusPost)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewNotFoundResponse())
	}

	response := []PostResponse{}
	for _, post := range postDB {
		if len(post.Tags) >= 0 {

			tags := []string{}
			for _, tag := range post.Tags {
				tags = append(tags, tag.Name)
			}
			response = append(response, PostResponse{
				ID:     int(post.ID),
				Title:  post.Title,
				Body:   post.Body,
				Status: post.Status,
				Tags:   tags,
			})
		}
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(response))
}

func (pc PostController) PublishPost(c echo.Context) error {
	idUser := c.Get("id").(uint)
	idPost, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestResponse())
	}

	err = pc.Repository.PublishPost(idPost, int(idUser))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewNotFoundResponse())
	}

	return c.JSON(http.StatusOK, utils.NewSuccessOperationResponse())
}

func (pc PostController) DeletePost(c echo.Context) error {
	idUser := c.Get("id").(uint)
	idPost, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestResponse())
	}

	err = pc.Repository.DeletePost(idPost, int(idUser))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewNotFoundResponse())
	}

	return c.JSON(http.StatusOK, utils.NewSuccessOperationResponse())
}

func (pc PostController) AllPostPublish(c echo.Context) error {

	pageNum, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pagesize"))
	userID, _ := strconv.Atoi(c.QueryParam("userid"))
	tags := c.QueryParam("tags")
	postFilter := entity.PostsFilter{
		PageNum:  pageNum,
		PageSize: pageSize,
		UserID:   userID,
		Tags:     strings.Split(tags, ","),
	}

	postDB, err := pc.Repository.AllPostPublish(postFilter)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewNotFoundResponse())
	}

	response := []PostResponse{}
	for _, post := range postDB {
		if len(post.Tags) > 0 {

			tags := []string{}
			for _, tag := range post.Tags {
				tags = append(tags, tag.Name)
			}
			response = append(response, PostResponse{
				ID:     int(post.ID),
				Title:  post.Title,
				Body:   post.Body,
				Status: post.Status,
				Tags:   tags,
			})
		}
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(response))
}

func (pc PostController) PostByIDPost(c echo.Context) error {

	idPost, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestResponse())
	}

	postDB, err := pc.Repository.ReadPostByPostId(idPost)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewNotFoundResponse())
	}

	tags := []string{}
	for _, tag := range postDB.Tags {
		tags = append(tags, tag.Name)
	}

	response := PostResponse{}
	response = PostResponse{
		ID:     int(postDB.ID),
		Title:  postDB.Title,
		Body:   postDB.Body,
		Status: postDB.Status,
		Tags:   tags,
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(response))
}
