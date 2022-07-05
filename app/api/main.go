package main

import (
	"fmt"
	"medium-be/internal/config"
	"medium-be/internal/constants"

	"medium-be/internal/database/postgres"
	_redisConn "medium-be/internal/database/redis"

	"medium-be/internal/routes"
	"medium-be/internal/utils"

	tc "medium-be/internal/controller/tags"
	_tagRepo "medium-be/internal/repository/postgres/tags"

	pc "medium-be/internal/controller/posts"
	_postRepo "medium-be/internal/repository/postgres/posts"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {

	// Setup Configuration
	cfg := config.NewConfig()

	// Setup Postgres
	db, err := postgres.NewPostgresRepo(&cfg.DatabaseConfig)
	checkErr(err)

	// Initialize redis
	constants.Rdb = _redisConn.NewRedisClientFromConfig(&cfg.RedisConfig)

	e := echo.New()

	e.Validator = &utils.Validator{Validator: validator.New()}

	// Setup Repository
	postRepo := _postRepo.NewPostRepository(db)
	postController := pc.NewPostController(postRepo)

	tagRepo := _tagRepo.NewTagRepository(db)
	tagController := tc.NewTagsController(tagRepo)

	// tagService :=
	routes.NewsPath(e, postController)
	routes.TagPath(e, tagController)

	address := fmt.Sprintf(":%d", cfg.Port)
	e.Logger.Fatal(e.Start(address))
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
