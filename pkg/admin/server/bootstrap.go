package server

import (
	"affiliate/internal/module/logger"
	"affiliate/pkg/admin/router"
	"affiliate/pkg/admin/server/initialize"

	"github.com/labstack/echo/v4"
)

// Bootstrap ...
func Bootstrap(e *echo.Echo) {
	logger.Init("Admin", "admin")

	// Init modules
	initialize.Init()

	// Route
	router.Init(e)
}
