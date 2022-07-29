package middlewares

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetLogger(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}: ${status} ${method} ${host}${uri} ${latency_human}\n",
		Output: os.Stdout,
	}))
	file, err := os.Create("log/request.log")
	if err == nil {
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "${time_rfc3339}: ${status} ${method} ${host}${uri} ${latency_human}\n",
			Output: file,
		}))
	}
}
