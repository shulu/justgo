package main

import (
	"log"
	"os"
	"url-shortener/api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {

	app := fiber.New()
	app.Use(logger.New())

	setUpRoutes(app)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
