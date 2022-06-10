package tags

import "medium-be/internal/entity"

type TagInterface interface {
	CreateTag(newTag entity.Tag) error
	DeleteTag(tagId int) error
	GetTagId(tagId int) (entity.Tag, error)
}
