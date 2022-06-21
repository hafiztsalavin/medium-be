package main

import (
	"fmt"
	"log"
	"medium-be/internal/config"
	"medium-be/internal/entity"
	"medium-be/internal/repository"

	"gorm.io/gorm"
)

func main() {
	cfg := config.NewConfig()

	db, err := repository.NewPostgresRepo(&cfg.DatabaseConfig)
	if err != nil {
		log.Fatal(err)
	}

	seedersFunc := []func(*gorm.DB) error{
		TagSeed}

	for _, function := range seedersFunc {
		if err := function(db); err != nil {
			log.Fatal(err)
		}
	}
}

func TagSeed(db *gorm.DB) error {

	var tags []entity.Tag = []entity.Tag{
		{Name: "otomotive"},
		{Name: "software engineering"},
		{Name: "furniture"},
		{Name: "live"},
		{Name: "health"},
	}

	for _, tag := range tags {
		entryTags := entity.Tag{
			Name: tag.Name,
		}
		db.Create(&entryTags)
	}

	fmt.Println("Tags(s) created")
	return nil
}
