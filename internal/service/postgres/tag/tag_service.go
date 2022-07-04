package tag

import "medium-be/internal/entity"

type TagService interface {
	CreateTag(newTag entity.Tag) error
	DeleteTag(tagId int) error
	GetTagId(tagId int) (entity.Tag, error)
	GetAllTag() ([]entity.Tag, error)
	EditTag(tagId int, newTag entity.Tag) error
}
