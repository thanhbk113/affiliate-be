package router

import (
	"affiliate/pkg/admin/handler"
	"affiliate/pkg/admin/router/routervalidation"

	"github.com/labstack/echo/v4"
)

func subPost(e *echo.Group) {
	var (
		g = e.Group("/subPost")
		h = handler.SubPost()
		v = routervalidation.SubPost()
		c = routervalidation.Common()
	)
	// Check permission

	g.POST("", h.Create, v.Create)
	g.GET("/:id", h.GetList, v.GetList, c.ParamID)
	g.PUT("/:id", h.Update, v.Update, c.ParamID)
	g.GET("/detail/:id", h.GetDetail, c.ParamID)
}
