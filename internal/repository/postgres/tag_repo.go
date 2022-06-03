package postgres

import (
	"gorm.io/gorm"
)

type TagInterface interface {
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *tagRepository {
	return &tagRepository{db: db}
}
