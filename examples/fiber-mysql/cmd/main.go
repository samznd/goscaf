package main

import (
	"fiber-mysql/config"
	"fiber-mysql/internal/handlers"
	"fiber-mysql/internal/routes"
	"fiber-mysql/pkg/utils"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	utils.InitialEnv()
	config.Connect()
	app := fiber.New()

	// Define routes
	api := app.Group("/api/v1")
	routes.SetupRoutes(api, &handlers.Handler{})

	log.Println("🚀 Fiber server is running on http://localhost:3000")
	app.Listen(":3000")
}
