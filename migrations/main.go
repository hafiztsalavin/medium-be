package main

import (
	"fmt"
	"log"

	"medium-be/internal/config"
	repository "medium-be/internal/database/postgres"
	"medium-be/internal/entity"

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

	if err := db.AutoMigrate(&entity.Posts{}); err != nil {
		log.Fatal(err)
		return err
	}

	// if err := db.SetupJoinTable(&entity.Posts{}, "Tags", &entity.PostTags{}); err != nil {
	// 	println(err.Error())
	// 	panic("Failed to setup join table")
	// }

	fmt.Println("Model(s) migrated")

	return nil
}

// func (PostTags) BeforeCreate(db *gorm.DB) error {
// 	err := db.SetupJoinTable(&Posts{}, "Tags", &PostTags{})

// 	if err != nil {
// 		return errors.New(err.Error())
// 	}

// 	return nil
// }
