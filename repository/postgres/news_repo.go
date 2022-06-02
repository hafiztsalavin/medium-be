package postgres

import (
	"gorm.io/gorm"
)

type NewsInterface interface {
}

type newsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) *newsRepository {
	return &newsRepository{db: db}
}
