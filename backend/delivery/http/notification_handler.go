package http

import (
	"smart-aquaculture-backend/repository"

	"github.com/gofiber/fiber/v2"
)

type NotificationHandler struct {
	tokenRepo repository.TokenRepository
}

func NewNotificationHandler(repo repository.TokenRepository) *NotificationHandler {
	return &NotificationHandler{repo}
}

func (h *NotificationHandler) SaveToken(c *fiber.Ctx) error {
	var payload struct {
		Token string `json:"token"`
	}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if payload.Token != "" {
		h.tokenRepo.SaveToken(payload.Token)
	}
	return c.JSON(fiber.Map{"success": true})
}
