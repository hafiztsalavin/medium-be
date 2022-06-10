package tags

import (
	"medium-be/internal/repository/tags"
)

type TagController struct {
	Repository tags.TagInterface
}

// init
func NewTagsController(tagsInterface tags.TagInterface) *TagController {
	return &TagController{Repository: tagsInterface}
}
