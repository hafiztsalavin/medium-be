package posts

import (
	"errors"
	"medium-be/internal/entity"

	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{db: db}
}

func (pr *postRepository) CreatePost(newPost entity.Posts, tags []int) error {
	existedPost, _ := pr.postGetByTitle(newPost.Title)
	if existedPost.Title != "" {
		return errors.New("duplicate title")
	}

	if err := pr.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&newPost).Error; err != nil {
			return err
		}

		for _, tag := range tags {
			if err := tx.Create(&entity.PostTags{
				PostsID: newPost.ID,
				TagID:   uint(tag),
			}).Error; err != nil {
				return errors.New("one of tags not exist")
			}
		}

		return nil

	}); err != nil {
		return err
	}

	return nil
}

func (pr *postRepository) postGetByTitle(titlePost string) (entity.Posts, error) {
	rec := entity.Posts{}

	err := pr.db.Where("title = ?", titlePost).First(&rec).Error
	if err != nil {
		return entity.Posts{}, err
	}

	return rec, nil
}
