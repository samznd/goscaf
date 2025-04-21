package handlers

import (
	"fiber-mysql/internal/services"
	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	service services.Service
}

func NewHandler(s services.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Get(c fiber.Ctx) error {
	message, err := h.service.GetMessage()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": message})
}
