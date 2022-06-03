package postgres

import (
	"news-be/internal/entity"

	"gorm.io/gorm"
)

type NewsInterface interface {
	ReadOne() ([]entity.News, error)
}

type newsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) *newsRepository {
	return &newsRepository{db: db}
}
