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

func (tc *TagController) Create(c echo.Context) error {
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

func (tc *TagController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))

	err := tc.Repository.DeleteTag(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse(404, err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewSuccessOperationResponse())
}

func (tc *TagController) Read(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil || id <= 0 {
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
