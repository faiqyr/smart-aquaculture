package http

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gofiber/fiber/v2"
)

type ControlHandler struct {
	mqttClient mqtt.Client
}

func NewControlHandler(router fiber.Router, mqttClient mqtt.Client) {
	handler := &ControlHandler{
		mqttClient: mqttClient,
	}

	router.Post("/feed", handler.FeedControl)
}

func (h *ControlHandler) FeedControl(c *fiber.Ctx) error {
	topic := "kolam/aquarium/control"
	payload := `{"command": "feed"}`

	token := h.mqttClient.Publish(topic, 0, false, payload)
	token.Wait()
	
	if token.Error() != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to publish mqtt message: %v", token.Error()),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Feed command sent to MQTT",
	})
}
