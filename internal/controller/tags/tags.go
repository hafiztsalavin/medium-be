package tags

import (
	"medium-be/internal/entity"
	"medium-be/internal/repository/tags"
	"medium-be/internal/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TagController struct {
	Repository tags.TagInterface
}

// init
func NewTagsController(tagsInterface tags.TagInterface) *TagController {
	return &TagController{Repository: tagsInterface}
}

func (tc *TagController) CreateTag(c echo.Context) error {
	var tagRequest TagRequest

	c.Bind(&tagRequest)
	if err := c.Validate(&tagRequest); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestResponse())
	}

	tag := entity.Tag{
		Name: tagRequest.Name,
	}

	err := tc.Repository.CreateTag(tag)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(406, "Tag already exist"))
	}

	return c.JSON(http.StatusOK, utils.NewSuccessOperationResponse())
}

func (tc *TagController) DeleteTag(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse(404, err.Error()))
	}

	err = tc.Repository.DeleteTag(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse(404, err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewSuccessOperationResponse())
}

func (tc *TagController) ReadTag(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestResponse())
	}

	existedTag, err := tc.Repository.GetTagId(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse(404, err.Error()))
	}

	response := TagResponse{
		ID:  existedTag.ID,
		Tag: existedTag.Name,
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(response))
}

func (tc *TagController) ReadAllTag(c echo.Context) error {
	responseTag := []TagResponse{}

	allTag, err := tc.Repository.GetAllTag()
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse(404, err.Error()))
	}

	for _, tag := range allTag {
		responseTag = append(responseTag, TagResponse{
			ID:  tag.ID,
			Tag: tag.Name,
		})
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(responseTag))
}

func (tc *TagController) UpdateTag(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse(404, err.Error()))
	}
	var tagRequest TagRequest

	c.Bind(&tagRequest)
	if err := c.Validate(&tagRequest); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestResponse())
	}

	tag := entity.Tag{
		Name: tagRequest.Name,
	}

	err = tc.Repository.EditTag(id, tag)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(404, err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewSuccessOperationResponse())
}
