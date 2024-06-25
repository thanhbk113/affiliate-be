package router

import (
	"affiliate/pkg/admin/handler"
	"affiliate/pkg/admin/router/routervalidation"

	"github.com/labstack/echo/v4"
)

func parPost(e *echo.Group) {
	var (
		g = e.Group("/parPost")
		h = handler.ParPost()
		v = routervalidation.ParPost()
	)
	// Check permission

	g.POST("", h.Create, v.Create)
	g.GET("", h.GetList, v.GetList)
}
