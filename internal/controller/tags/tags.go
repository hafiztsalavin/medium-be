package tags

import (
	"encoding/json"
	"medium-be/internal/entity"
	redisRepo "medium-be/internal/repository/redis"
	service "medium-be/internal/service/postgres/tag"
	"medium-be/internal/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var tagEntity string = "tag"

type TagController struct {
	Service service.TagService
}

// init
func NewTagsController(tagsService service.TagService) *TagController {
	return &TagController{Service: tagsService}
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

	err := tc.Service.CreateTag(tag)

	if err != nil {
		return c.JSON(http.StatusConflict, utils.ErrorResponse(409, err.Error()))
	}

	go redisRepo.DeleteCache(tagEntity)

	return c.JSON(http.StatusOK, utils.NewSuccessOperationResponse())
}

func (tc *TagController) DeleteTag(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse(404, err.Error()))
	}

	err = tc.Service.DeleteTag(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse(404, err.Error()))
	}

	go redisRepo.DeleteCache(tagEntity)

	return c.JSON(http.StatusOK, utils.NewSuccessOperationResponse())
}

func (tc *TagController) ReadTag(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestResponse())
	}

	existedTag, err := tc.Service.GetTagId(id)
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

	tagCache, err := redisRepo.GetCache(tagEntity, 0, "")
	if err == nil {
		_ = json.Unmarshal([]byte(tagCache), &responseTag)
		return c.JSON(http.StatusOK, utils.SuccessResponse(responseTag))
	}

	allTag, err := tc.Service.GetAllTag()
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse(404, err.Error()))
	}

	for _, tag := range allTag {
		responseTag = append(responseTag, TagResponse{
			ID:  tag.ID,
			Tag: tag.Name,
		})
	}

	resMarshal, _ := json.Marshal(responseTag)
	go redisRepo.CreateCache(tagEntity, 0, "", resMarshal)

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

	err = tc.Service.EditTag(id, tag)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(404, err.Error()))
	}

	go redisRepo.DeleteCache(tagEntity)

	return c.JSON(http.StatusOK, utils.NewSuccessOperationResponse())
}
