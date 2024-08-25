package main

import (
	"fmt"
	"jayendrapawar/goSmartURL/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	setUpRoutes(app)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000" // Default port
	}

	log.Fatal(app.Listen(":" + port))
}
