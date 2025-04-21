
package routes

import (
	"fiber-mysql/internal/handlers"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app fiber.Router, h *handlers.Handler) {
	api := app.Group("/api")
	api.Get("/message", h.Get)
}
