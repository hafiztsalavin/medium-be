package main

import (
	"fmt"
	"news-be/internal/config"
	"news-be/internal/middlewares"
	"news-be/internal/routes"

	"github.com/labstack/echo"
)

func main() {

	// Initialize news config
	cfg := config.NewConfig()

	// Initialize DB repositories
	// db, err := postgres.NewPostgresRepo(&cfg.DatabaseConfig)
	// checkErr(err)

	e := echo.New()

	// repository
	// tagRepo := postgres.NewTagRepository(db)
	// newsRepo := postgres.NewNewsRepository(db)

	// controller
	// nc := news.NewNewsController(newsRepo)

	userGroup := e.Group("/auth")
	middlewares.CheckJWT(userGroup)
	routes.UserPath(userGroup)

	address := fmt.Sprintf(":%d", cfg.Port)
	e.Logger.Fatal(e.Start(address))
}

// func checkErr(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }
