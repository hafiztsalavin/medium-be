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

const statusPublish = "publish"

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

func (pr *postRepository) EditPost(idPost int, editPost entity.Posts, tags []int) error {
	existedPost, _ := pr.postGetByTitle(editPost.Title)
	if existedPost.Title != "" && existedPost.ID != uint(idPost) {
		return errors.New("you can't use this title, because title already used in another article")
	}

	var post entity.Posts
	if err := pr.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&post, idPost).Error; err != nil {
			return err
		}

		if err := pr.db.Model(&post).Updates(editPost).Error; err != nil {
			return err
		}

		if err := tx.Delete(&entity.PostTags{}, "posts_id = ?", idPost).Error; err != nil {
			return err
		}

		for _, tag := range tags {
			if err := tx.Create(&entity.PostTags{
				PostsID: uint(idPost),
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

func (pr *postRepository) ReadPost(id int) (entity.Posts, error) {
	var post entity.Posts

	if err := pr.db.Preload("Tags").First(&post, id).Error; err != nil {
		return post, err
	}

	return post, nil
}

func (pr *postRepository) ReadAllPost(idUser int, statusPost string) ([]entity.Posts, error) {
	var post []entity.Posts

	if statusPost != "" {
		if err := pr.db.Preload("Tags").Where("user_id = ? AND status = ?", idUser, statusPost).Find(&post).Error; err != nil {
			return post, err
		}
	} else {
		if err := pr.db.Preload("Tags").Where("user_id = ?", idUser).Find(&post).Error; err != nil {
			return post, err
		}
	}

	return post, nil
}

func (pr *postRepository) PublishPost(idPost, idUser int) error {
	var post entity.Posts

	if err := pr.db.Where("id = ? AND user_id = ?", idPost, idUser).First(&post).Error; err != nil {

		return err
	}

	if err := pr.db.Model(&post).Update("status", "publish").Error; err != nil {
		return err

	}
	return nil
}

func (pr *postRepository) DeletePost(idPost, idUser int) error {
	var post entity.Posts

	if err := pr.db.Where("id = ? AND user_id = ?", idPost, idUser).First(&post).Error; err != nil {

		return err
	}

	if err := pr.db.Model(&post).Update("status", "deleted").Error; err != nil {
		return err
	}

	pr.db.Delete(&post)
	return nil
}

func (pr *postRepository) AllPostPublish(filter entity.PostsFilter) ([]entity.Posts, error) {
	var post []entity.Posts

	offset := filter.PageSize * (filter.PageNum - 1)

	if filter.UserID != 0 && filter.Tags[0] != "" {
		pr.db.Preload("Tags", "name IN (?)", filter.Tags).Where("status = ? AND user_id = ? ", statusPublish, filter.UserID).Offset(offset).Limit(filter.PageSize).Find(&post)
	} else if filter.Tags[0] != "" {
		pr.db.Preload("Tags", "name IN (?)", filter.Tags).Where("status = ?", statusPublish).Offset(offset).Limit(filter.PageSize).Find(&post)
	} else if filter.UserID != 0 {
		pr.db.Preload("Tags").Where("status = ? AND user_id = ? ", statusPublish, filter.UserID).Offset(offset).Limit(filter.PageSize).Find(&post)
	} else {
		pr.db.Preload("Tags").Where("status = ?", statusPublish).Offset(offset).Limit(filter.PageSize).Find(&post)
	}
	return post, nil
}

func (pr *postRepository) ReadPostByPostId(idPost int) (entity.Posts, error) {
	var post entity.Posts

	if err := pr.db.Preload("Tags").Where("id = ? AND status = ?", idPost, statusPublish).Find(&post).Error; err != nil {
		return post, err
	}

	return post, nil
}

func (pr *postRepository) postGetByTitle(titlePost string) (entity.Posts, error) {
	rec := entity.Posts{}

	err := pr.db.Where("title = ?", titlePost).First(&rec).Error
	if err != nil {
		return entity.Posts{}, err
	}

	return rec, nil
}
