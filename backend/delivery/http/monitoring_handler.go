package http

import (
	"smart-aquaculture-backend/usecase"

	"github.com/gofiber/fiber/v2"
)

type MonitoringHandler struct {
	monitoringUsecase usecase.MonitoringUsecase
}

func NewMonitoringHandler(router fiber.Router, usecase usecase.MonitoringUsecase) {
	handler := &MonitoringHandler{
		monitoringUsecase: usecase,
	}

	router.Get("/latest", handler.GetLatest)
	router.Get("/history", handler.GetHistory)
}

func (h *MonitoringHandler) GetLatest(c *fiber.Ctx) error {
	data, err := h.monitoringUsecase.GetLatestMonitoring()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch latest monitoring data",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func (h *MonitoringHandler) GetHistory(c *fiber.Ctx) error {
	// Batasi 20 data terakhir
	data, err := h.monitoringUsecase.GetMonitoringHistory(20)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch monitoring history",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}
