package main

import (
	"fmt"
	"medium-be/internal/config"
	"medium-be/internal/utils"

	dbPostgres "medium-be/internal/database/postgres"
	dbRedis "medium-be/internal/database/redis"

	_postRepo "medium-be/internal/repository/postgres/posts"
	_redisRepo "medium-be/internal/repository/redis"

	"medium-be/internal/routes"
	// "medium-be/internal/utils"

	_postController "medium-be/internal/controller/posts"
	_postService "medium-be/internal/service/postgres/post"

	"github.com/labstack/echo/v4"
	// tc "medium-be/internal/controller/tags"
	// _tagRepo "medium-be/internal/repository/postgres/tags"
	"github.com/go-playground/validator/v10"
	// "github.com/labstack/echo/v4"
)

func main() {

	// Setup Configuration
	cfg := config.NewConfig()

	// Setup Postgres & redis
	db, err := dbPostgres.NewPostgresRepo(&cfg.DatabaseConfig)
	checkErr(err)

	// Initialize redis
	dbR, err := dbRedis.NewRedisClientFromConfig(&cfg.RedisConfig)
	checkErr(err)

	// Setup repository
	postRepo := _postRepo.NewPostRepository(db)
	cacheRepo := _redisRepo.NewRedisRepository(dbR)

	// Setup service
	servicePost := _postService.NewPostService(postRepo, cacheRepo)

	// Setup controller
	postController := _postController.NewPostController(servicePost)

	e := echo.New()

	e.Validator = &utils.Validator{Validator: validator.New()}

	// // Setup Repository
	// postController := pc.NewPostController(postRepo)

	// tagRepo := _tagRepo.NewTagRepository(db)
	// tagController := tc.NewTagsController(tagRepo)

	// // tagService :=
	routes.PostPath(e, postController)
	// routes.TagPath(e, tagController)

	address := fmt.Sprintf(":%d", cfg.Port)
	e.Logger.Fatal(e.Start(address))
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
