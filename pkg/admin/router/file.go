package router

import (
	"affiliate/pkg/admin/handler"

	echocustom "affiliate/internal/echo"

	"github.com/labstack/echo/v4"
)

func file(e *echo.Group) {
	var (
		g = e.Group("/file")

		h = handler.File()
	)

	g.POST("/photo", h.UploadImage, echocustom.UploadSingle())
}
