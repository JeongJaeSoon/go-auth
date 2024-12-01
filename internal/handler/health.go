package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type HealthHandler struct {
	logger *zap.Logger
}

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

func NewHealthHandler(logger *zap.Logger) *HealthHandler {
	return &HealthHandler{
		logger: logger,
	}
}

func (h *HealthHandler) Check(c *fiber.Ctx) error {
	h.logger.Info("Health check requested")

	response := HealthResponse{
		Status:    "OK",
		Timestamp: time.Now(),
	}

	return c.JSON(response)
}
