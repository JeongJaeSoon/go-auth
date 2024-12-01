package handler

import (
	"time"

	"github.com/JeongJaeSoon/go-auth/internal/generated/health"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type HealthHandler struct {
	logger *zap.Logger
}

func NewHealthHandler(logger *zap.Logger) *HealthHandler {
	return &HealthHandler{
		logger: logger,
	}
}

// Check implements the health check endpoint
func (h *HealthHandler) Check(c *fiber.Ctx) error {
	h.logger.Info("Health check requested")

	status := "OK"
	now := time.Now()
	response := health.HealthResponse{
		Status:    &status,
		Timestamp: &now,
	}

	return c.JSON(response)
}
