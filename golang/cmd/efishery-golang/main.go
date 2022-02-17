package main

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rindangramadhan/efishery-be-task/internal/helper"
	"github.com/rindangramadhan/efishery-be-task/internal/routes"

	db "github.com/rindangramadhan/efishery-be-task/internal/db/sqlite"
)

func main() {
	// Load environment file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api := echo.New()

	// Custom Binder
	api.Binder = &helper.CustomBinder{}

	// Custom Validator
	api.Validator = &helper.CustomValidator{Validator: validator.New()}

	api.Use(middleware.Logger())
	api.Use(middleware.Recover())

	// Cors Middleware for API endpoint.
	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.POST},
	}))

	// Database
	db.SQLite.Init()

	// Route
	routes.RouteApply(api)
	routes.Swagger.Init(api)

	server := echo.New()
	server.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()

		api.ServeHTTP(res, req)

		return
	})

	server.Logger.Fatal(server.Start(":" + os.Getenv("SERVICE_PORT")))
}
