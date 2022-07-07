package posts

import (
	"medium-be/internal/entity"
	service "medium-be/internal/service/postgres/post"
	"medium-be/internal/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type PostController struct {
	Service service.PostService
}

func NewPostController(postService service.PostService) *PostController {
	return &PostController{Service: postService}
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
		Status: postRequest.Status,
		Body:   postRequest.Body,
	}

	err := pc.Service.CreatePost(newPost, postRequest.Tags)
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

	err = pc.Service.EditPost(idPost, editPost, postRequest.Tags)
	if err != nil {
		return c.JSON(http.StatusConflict, utils.ErrorResponse(409, err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewSuccessOperationResponse())
}

func (pc PostController) ReadPost(c echo.Context) error {
	idPost, err := strconv.Atoi(c.Param("id"))
	idUser := c.Get("id").(uint)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestResponse())
	}

	postDB, err := pc.Service.ReadPost(idPost)
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

	postDB, err := pc.Service.ReadAllPost(int(idUser), statusPost)
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

	err = pc.Service.PublishPost(idPost, int(idUser))
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

	err = pc.Service.DeletePost(idPost, int(idUser))
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

	postDB, err := pc.Service.AllPostPublish(postFilter)
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

	postDB, err := pc.Service.ReadPostByPostId(idPost)
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
