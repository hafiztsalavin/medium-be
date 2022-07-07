package tag

import (
	"medium-be/internal/entity"
	_tagRepo "medium-be/internal/repository/postgres/tags"
	_redisRepo "medium-be/internal/repository/redis"
)

type TagService interface {
	CreateTag(newTag entity.Tag) error
	DeleteTag(tagId int) error
	GetTagId(tagId int) (entity.Tag, error)
	GetAllTag() ([]entity.Tag, error)
	EditTag(tagId int, newTag entity.Tag) error
}

type tagService struct {
	tagService _tagRepo.TagRepository
	redisRepo  _redisRepo.RedisRepository
}

func NewTagService(tagRepo _tagRepo.TagRepository, redisRepo _redisRepo.RedisRepository) TagService {
	return &tagService{
		tagService: tagRepo,
		redisRepo:  redisRepo,
	}
}

func (ts *tagService) CreateTag(newTag entity.Tag) error {
	err := ts.tagService.CreateTag(newTag)
	return err
}

func (ts *tagService) DeleteTag(tagId int) error {
	err := ts.tagService.DeleteTag(tagId)
	return err
}

func (ts *tagService) GetTagId(tagId int) (entity.Tag, error) {
	tag, err := ts.tagService.GetTagId(tagId)
	return tag, err
}

func (ts *tagService) GetAllTag() ([]entity.Tag, error) {
	tags, err := ts.tagService.GetAllTag()
	return tags, err
}

func (ts *tagService) EditTag(tagId int, newTag entity.Tag) error {
	err := ts.tagService.EditTag(tagId, newTag)
	return err
}
