package main

import (
	"fmt"
	"medium-be/internal/config"
	"medium-be/internal/repository"
	"medium-be/internal/routes"

	pc "medium-be/internal/controller/posts"
	pr "medium-be/internal/repository/posts"

	tc "medium-be/internal/controller/tags"
	tr "medium-be/internal/repository/tags"

	"github.com/labstack/echo/v4"
)

func main() {

	// Setup Configuration
	cfg := config.NewConfig()

	// Setup Postgres
	db, err := repository.NewPostgresRepo(&cfg.DatabaseConfig)
	checkErr(err)

	e := echo.New()
	// Setup Repository
	newsRepo := pr.NewPostRepository(db)
	newsController := pc.NewPostController(newsRepo)

	tagRepo := tr.NewTagRepository(db)
	tagController := tc.NewTagsController(tagRepo)

	routes.NewsPath(e, newsController)
	routes.TagPath(e, tagController)

	address := fmt.Sprintf(":%d", cfg.Port)
	e.Logger.Fatal(e.Start(address))
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
