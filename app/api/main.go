package main

import (
	"fmt"
	"medium-be/internal/config"
	"medium-be/internal/middlewares"
	"medium-be/internal/utils"
	"time"

	dbPostgres "medium-be/internal/database/postgres"
	dbRedis "medium-be/internal/database/redis"

	_redisRepo "medium-be/internal/repository/redis"

	"medium-be/internal/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	_postController "medium-be/internal/controller/posts"
	_postRepo "medium-be/internal/repository/postgres/posts"
	_postService "medium-be/internal/service/post"

	_tagController "medium-be/internal/controller/tags"
	_tagRepo "medium-be/internal/repository/postgres/tags"
	_tagService "medium-be/internal/service/tag"
)

func main() {

	// Setup Configuration
	cfg := config.NewConfig()

	timeoutContext := time.Duration(cfg.DatabaseConfig.Timeout) * time.Second

	// Setup Postgres & redis
	db, err := dbPostgres.NewPostgresRepo(&cfg.DatabaseConfig)
	checkErr(err)

	// Initialize redis
	dbR, err := dbRedis.NewRedisClientFromConfig(&cfg.RedisConfig)
	checkErr(err)

	// Setup repository
	postRepo := _postRepo.NewPostRepository(db)
	tagRepo := _tagRepo.NewTagRepository(db)

	cacheRepo := _redisRepo.NewRedisRepository(dbR)

	// Setup service
	servicePost := _postService.NewPostService(postRepo, cacheRepo)
	serviceTag := _tagService.NewTagService(tagRepo, cacheRepo, timeoutContext)

	// Setup controller
	postController := _postController.NewPostController(servicePost)
	tagController := _tagController.NewTagsController(serviceTag)

	e := echo.New()

	e.Validator = &utils.Validator{Validator: validator.New()}

	middlewares.SetLogger(e)

	routes.PostPath(e, postController)
	routes.TagPath(e, tagController)

	address := fmt.Sprintf(":%d", cfg.Port)
	e.Logger.Fatal(e.Start(address))
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
