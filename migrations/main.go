package main

import (
	"fmt"
	"log"

	"medium-be/internal/config"
	"medium-be/internal/entity"
	"medium-be/internal/repository"

	"gorm.io/gorm"
)

// drop tables
func DropTables(db *gorm.DB) error {

	if err := db.Migrator().DropTable(&entity.Posts{}, &entity.Tag{}, &entity.PostTags{}); err != nil {
		return err
	}

	fmt.Println("Table(s) dropped")

	return nil
}

func main() {
	cfg := config.NewConfig()

	db, err := repository.NewPostgresRepo(&cfg.DatabaseConfig)
	if err != nil {
		log.Fatal(err)
	}

	if err := MigrateModels(db); err != nil {
		log.Fatalf("Error when migrate models, %v", err)
	}
}

// make a new table based on entity
func MigrateModels(db *gorm.DB) error {
	if err := DropTables(db); err != nil {
		return err
	}

	if err := db.AutoMigrate(&entity.Tag{}, &entity.Posts{}); err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Model(s) migrated")

	return nil
}
