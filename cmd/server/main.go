package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JeongJaeSoon/go-auth/config"
	"github.com/JeongJaeSoon/go-auth/internal/logging"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func StartServer(cfg *config.Config) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	port := fmt.Sprintf(":%d", cfg.Server.Port)
	logging.Logger.Info("Server starting on port %s", zap.String("port", port))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
			logging.Logger.Error("Failed to start server", zap.Error(err))
			quit <- syscall.SIGTERM
		}
	}()

	// SIGINT or SIGTERM will trigger this graceful shutdown
	<-quit
	logging.Logger.Info("Shutting down server gracefully...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Shutdown(); err != nil {
		logging.Logger.Error("Error shutting down server", zap.Error(err))
	} else {
		logging.Logger.Info("Server stopped gracefully")
	}
}
