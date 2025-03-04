package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/joho/godotenv"

	"github.com/ramalloc/go-url-shortner/routes"
)

func SeuptRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	// Loads environment variables from a .env file.
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	app := fiber.New()

	// Middleware to log HTTP requests.
	app.Use(logger.New())

	SeuptRoutes(app)


	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000" // Default to 3000 if APP_PORT is not set
	}

	fmt.Println("Server starting on port:", port)
	log.Fatal(app.Listen(":" + port))
}
