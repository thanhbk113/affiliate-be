package main

import (
	"affiliate/docs/admin"
	"affiliate/internal/config"
	"affiliate/pkg/admin/server"
	"fmt"
	"os"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.elastic.co/apm/module/apmechov4"
)

// @title affiliate - Admin API
// @version 1.0
// @description All APIs for affiliate Manage admin.
// @description
// @description ******************************
// @description - Add description
// @description ******************************
// @description
// @termsOfService https://bag-manage.vn
// @contact.name Dev team
// @contact.url https://bag-manage.vn
// @contact.email dev@reshare.vn
// @basePath /admin

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Echo instance
	e := echo.New()

	e.Use(apmechov4.Middleware())

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	}))
	e.Use(middleware.Gzip())
	if os.Getenv("ENV") == "release" {
		e.Use(middleware.Recover())
	}

	// Bootstrap things
	server.Bootstrap(e)

	//Swagger
	if config.IsEnvDevelop() {
		domain := config.GetENV().DOMAIN_ADMIN
		admin.SwaggerInfo.Host = domain
		e.GET(admin.SwaggerInfo.BasePath+"/swagger/*", echoSwagger.WrapHandler)
		fmt.Println("Swagger url : ", domain+admin.SwaggerInfo.BasePath+"/swagger/index.html")
	}

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
