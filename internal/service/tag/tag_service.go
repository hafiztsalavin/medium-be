package tag

import (
	"context"
	"encoding/json"
	"fmt"
	"medium-be/internal/entity"
	_tagRepo "medium-be/internal/repository/postgres/tags"
	_redisRepo "medium-be/internal/repository/redis"
	"time"
)

type TagService interface {
	CreateTag(ctx context.Context, newTag entity.Tag) error
	DeleteTag(ctx context.Context, tagId int) error
	GetTagId(ctx context.Context, tagId int) (entity.Tag, error)
	GetAllTag(ctx context.Context) ([]entity.Tag, error)
	EditTag(ctx context.Context, tagId int, newTag entity.Tag) error
}

type tagService struct {
	tagRepo        _tagRepo.TagRepository
	redisRepo      _redisRepo.RedisRepository
	contextTimeout time.Duration
}

func NewTagService(tagRepo _tagRepo.TagRepository, redisRepo _redisRepo.RedisRepository, timeout time.Duration) TagService {
	return &tagService{
		tagRepo:        tagRepo,
		redisRepo:      redisRepo,
		contextTimeout: timeout,
	}
}

const (
	tagEntity = "tag-post"
)

func (ts *tagService) CreateTag(ctx context.Context, newTag entity.Tag) error {
	ctx, cancel := context.WithTimeout(ctx, ts.contextTimeout)
	defer cancel()

	err := ts.tagRepo.CreateTag(newTag)

	return err
}

func (ts *tagService) DeleteTag(ctx context.Context, tagId int) error {
	ctx, cancel := context.WithTimeout(ctx, ts.contextTimeout)
	defer cancel()

	key := ts.keyPrefix(tagId)
	err := ts.tagRepo.DeleteTag(tagId)
	ts.redisRepo.DeleteCache(key)

	return err
}

func (ts *tagService) GetTagId(ctx context.Context, tagId int) (entity.Tag, error) {
	ctx, cancel := context.WithTimeout(ctx, ts.contextTimeout)
	defer cancel()

	var tag entity.Tag
	key := ts.keyPrefix(tagId)

	tagCached, _ := ts.redisRepo.GetCache(key)
	if err := json.Unmarshal([]byte(tagCached), &tag); err == nil {
		return tag, nil
	}

	tag, err := ts.tagRepo.GetTagId(tagId)
	if err != nil {
		return tag, err
	}

	tagString, _ := json.Marshal(&tag)
	ts.redisRepo.CreateCache(key, tagString, 0)

	return tag, nil
}

func (ts *tagService) GetAllTag(ctx context.Context) ([]entity.Tag, error) {
	ctx, cancel := context.WithTimeout(ctx, ts.contextTimeout)
	defer cancel()

	tags, err := ts.tagRepo.GetAllTag()
	return tags, err
}

func (ts *tagService) EditTag(ctx context.Context, tagId int, newTag entity.Tag) error {
	ctx, cancel := context.WithTimeout(ctx, ts.contextTimeout)
	defer cancel()

	err := ts.tagRepo.EditTag(tagId, newTag)
	key := ts.keyPrefix(tagId)
	ts.redisRepo.DeleteCache(key)

	return err
}

func (ts *tagService) keyPrefix(tagID int) string {
	return fmt.Sprintf("%s:%d", tagEntity, tagID)
}
