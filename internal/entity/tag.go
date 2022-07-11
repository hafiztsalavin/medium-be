package entity

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name string
}

type TagRequest struct {
	Name string `json:"name" validate:"required"`
}

type TagResponse struct {
	ID  uint   `json:"id"`
	Tag string `json:"tag"`
}
