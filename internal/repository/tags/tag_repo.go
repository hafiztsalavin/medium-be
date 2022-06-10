package tags

import (
	"errors"
	"medium-be/internal/entity"
	"strings"

	"gorm.io/gorm"
)

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *tagRepository {
	return &tagRepository{db: db}
}

func (tr *tagRepository) CreateTag(newTag entity.Tag) error {
	existedTag, err := tr.tagGetByUsername(newTag.Name)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
	}

	if existedTag != (entity.Tag{}) {
		return errors.New("duplicate data")
	}

	err = tr.saveTag(newTag)
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

func (tr *tagRepository) tagGetByUsername(tagName string) (entity.Tag, error) {
	rec := entity.Tag{}

	err := tr.db.Where("id = ?", tagName).First(&rec).Error
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
