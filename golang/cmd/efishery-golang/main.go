package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	db "github.com/rindangramadhan/efishery-be-task/internal/helper/db/sqlite"
)

func main() {
	// Load environment file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := echo.New()

	db.SQLite.Init()

	server.GET("/", func(c echo.Context) error {
		var response = struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{
			Status:  "ok",
			Message: "Hello eFishery!",
		}

		return c.JSON(http.StatusOK, response)
	})

	server.Logger.Fatal(server.Start(":" + os.Getenv("SERVICE_PORT")))
}
