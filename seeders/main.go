package main

import (
	"fmt"
	"log"
	"math/rand"
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
		TagSeed, PostSeed}

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

func PostSeed(db *gorm.DB) error {

	var posts []entity.Posts = []entity.Posts{
		{UserID: 1, Title: "how to", Body: "lorem ipsum dolor amet"},
		{UserID: 1, Title: "how to rich", Body: "lorem ipsum dolor amet"},
		{UserID: 1, Title: "how to became software engineer", Body: "lorem ipsum dolor amet"},
		{UserID: 2, Title: "how to became data engineer", Body: "lorem ipsum dolor amet"},
		{UserID: 2, Title: "how to became backend engineer", Body: "lorem ipsum dolor amet"},
		{UserID: 2, Title: "how to became frontend engineer", Body: "lorem ipsum dolor amet"},
		{UserID: 3, Title: "road map software engineer", Body: "lorem ipsum dolor amet"},
		{UserID: 3, Title: "road map backend engineer", Body: "lorem ipsum dolor amet"},
		{UserID: 3, Title: "learn golang", Body: "lorem ipsum dolor amet"},
	}

	status := []string{"draft", "publish"}
	for _, post := range posts {
		entryPost := entity.Posts{
			UserID: post.UserID,
			Title:  post.Title,
			Body:   post.Body,
			Status: status[rand.Intn(2-0)],
		}
		postTagEntry := entity.PostTags{
			PostsID: uint(rand.Intn(10-1) + 1),
			TagID:   uint(rand.Intn(6-1) + 1),
		}

		db.Create(&entryPost)
		db.Create(&postTagEntry)
	}

	fmt.Println("Post(s) created")
	return nil
}
