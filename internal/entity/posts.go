package entity

import (
	"errors"

	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	UserID uint
	Title  string
	Body   string
	Status string `gorm:"default:draft"`
	Tags   []Tag  `gorm:"many2many:post_tags;"`
}

type PostsFilter struct {
	Status string
	Tags   []string
}

type PostTags struct {
	PostsID uint `gorm:"primaryKey"`
	TagID   uint `gorm:"primaryKey"`
}

func (PostTags) BeforeCreate(db *gorm.DB) error {
	err := db.SetupJoinTable(&Posts{}, "Tags", &PostTags{})

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
