package routes

import (
	"fiber-postgres/internal/handlers"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app fiber.Router, h *handlers.Handler) {
	api := app.Group("/api")
	api.Get("/message", h.Get)
}
