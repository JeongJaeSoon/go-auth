package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JeongJaeSoon/go-auth/config"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Server struct {
	app    *fiber.App
	cfg    *config.Config
	logger *zap.Logger
}

func NewServer(cfg *config.Config, logger *zap.Logger) *Server {
	return &Server{
		app:    fiber.New(),
		cfg:    cfg,
		logger: logger,
	}
}

func (s *Server) Start() {
	app := s.app
	cfg := s.cfg
	logger := s.logger

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	port := fmt.Sprintf(":%d", cfg.Server.Port)
	logger.Info("Server starting on port", zap.String("port", port))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
			logger.Error("Failed to start server", zap.Error(err))
			quit <- syscall.SIGTERM
		}
	}()

	// SIGINT or SIGTERM will trigger this graceful shutdown
	<-quit
	logger.Info("Shutting down server gracefully...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Shutdown(); err != nil {
		s.logger.Error("Error shutting down server", zap.Error(err))
	} else {
		logger.Info("Server stopped gracefully")
	}
}
