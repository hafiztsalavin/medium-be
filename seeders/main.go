package main

import (
	"fmt"
	"log"
	"math/rand"
	"medium-be/internal/config"
	repository "medium-be/internal/database/postgres"
	"medium-be/internal/entity"

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
	size := len(posts)

	for i := 0; i < size; i++ {
		entryPost := entity.Posts{
			UserID: posts[i].UserID,
			Title:  posts[i].Title,
			Body:   posts[i].Body,
			Status: status[rand.Intn(2-0)],
		}

		postTagEntry := entity.PostTags{
			PostsID: uint(i + 1),
			TagID:   uint(rand.Intn(6-1) + 1),
		}
		db.Create(&entryPost)
		db.Create(&postTagEntry)

	}

	fmt.Println("Post(s) created")
	return nil
}
