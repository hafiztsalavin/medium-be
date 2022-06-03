package news

import (
	"net/http"

	"github.com/labstack/echo"
)

// var newsEntity string = "news"

// type NewsController struct {
// 	Repository postgres.NewsInterface
// }

// func NewNewsController(repository postgres.NewsInterface) *NewsController {
// 	return &NewsController{Repository: repository}
// }

// func(c echo.Context) error {
// 	return c.String(http.StatusOK, "OK")
// }

func ReadOne() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi, you have access!")
	}
}
