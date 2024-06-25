package router

import (
	"affiliate/internal/config"
	"affiliate/internal/middleware"

	"github.com/labstack/echo/v4"
)

// Init ...
func Init(e *echo.Echo) {
	// Middlewares ...
	var secretKey = config.GetENV().AuthSecret.Admin
	e.Use(middleware.JWT(secretKey))

	e.Use(middleware.CORSConfig())

	r := e.Group("/admin")
	file(r)
	parPost(r)
	subPost(r)
}
