package tags

import (
	"errors"
	"medium-be/internal/entity"

	"gorm.io/gorm"
)

type TagRepository interface {
	CreateTag(newTag entity.Tag) error
	DeleteTag(tagId int) error
	GetTagId(tagId int) (entity.Tag, error)
	GetAllTag() ([]entity.Tag, error)
	EditTag(tagId int, newTag entity.Tag) error
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *tagRepository {
	return &tagRepository{db: db}
}

func (tr *tagRepository) CreateTag(newTag entity.Tag) error {
	existedTag, _ := tr.tagGetByTag(newTag.Name)
	if existedTag != (entity.Tag{}) {
		return errors.New("duplicate data")
	}

	err := tr.saveTag(newTag)
	if err != nil {
		return err
	}

	return nil
}

func (tr *tagRepository) DeleteTag(tagId int) error {
	existedTag, err := tr.tagGetById(tagId)
	if err != nil {
		return err
	}

	err = tr.deleteTag(existedTag)
	if err != nil {
		return err
	}

	return nil
}

func (tr *tagRepository) GetTagId(tagId int) (entity.Tag, error) {
	existedTag, err := tr.tagGetById(tagId)
	if err != nil {
		return existedTag, err
	}

	return existedTag, nil
}

func (tr *tagRepository) GetAllTag() ([]entity.Tag, error) {
	var tags []entity.Tag

	tr.db.Find(&tags)

	return tags, nil
}

func (tr *tagRepository) EditTag(tagId int, newTag entity.Tag) error {
	existedTag, err := tr.tagGetById(tagId)
	if err != nil {
		return err
	}

	if err := tr.db.Model(&existedTag).Updates(newTag).Error; err != nil {
		return err
	}

	return nil
}

func (tr *tagRepository) saveTag(newTag entity.Tag) error {
	err := tr.db.Save(&newTag).Error

	if err != nil {
		return err
	}

	return nil
}

func (tr *tagRepository) deleteTag(tag entity.Tag) error {
	err := tr.db.Delete(&tag).Error

	if err != nil {
		return err
	}

	return nil
}

func (tr *tagRepository) tagGetByTag(tagName string) (entity.Tag, error) {
	rec := entity.Tag{}

	err := tr.db.Where("name = ?", tagName).First(&rec).Error
	if err != nil {
		return entity.Tag{}, err
	}

	return rec, nil
}

func (tr *tagRepository) tagGetById(tagId int) (entity.Tag, error) {
	rec := entity.Tag{}

	err := tr.db.Where("id = ?", tagId).First(&rec).Error
	if err != nil {
		return rec, err
	}

	return rec, nil
}
