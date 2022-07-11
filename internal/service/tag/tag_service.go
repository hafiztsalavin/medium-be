package tag

import (
	"encoding/json"
	"fmt"
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

const (
	tagEntity = "tag-post"
)

func (ts *tagService) CreateTag(newTag entity.Tag) error {
	err := ts.tagService.CreateTag(newTag)

	return err
}

func (ts *tagService) DeleteTag(tagId int) error {
	key := ts.keyPrefix(tagId)

	ts.redisRepo.DeleteCache(key)
	err := ts.tagService.DeleteTag(tagId)

	return err
}

func (ts *tagService) GetTagId(tagId int) (entity.Tag, error) {
	var tag entity.Tag
	key := ts.keyPrefix(tagId)

	tagCached, _ := ts.redisRepo.GetCache(key)
	if err := json.Unmarshal([]byte(tagCached), &tag); err == nil {
		return tag, nil
	}

	tag, err := ts.tagService.GetTagId(tagId)
	if err != nil {
		return tag, err
	}

	tagString, _ := json.Marshal(&tag)
	ts.redisRepo.CreateCache(key, tagString, 0)

	return tag, nil
}

func (ts *tagService) GetAllTag() ([]entity.Tag, error) {
	tags, err := ts.tagService.GetAllTag()
	return tags, err
}

func (ts *tagService) EditTag(tagId int, newTag entity.Tag) error {
	err := ts.tagService.EditTag(tagId, newTag)

	key := ts.keyPrefix(tagId)
	ts.redisRepo.DeleteCache(key)

	return err
}

func (ts *tagService) keyPrefix(tagID int) string {
	return fmt.Sprintf("%s:%d", tagEntity, tagID)
}
