package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func (s *Server) Start() {
	// Setup routes
	s.setupRoutes()

	app := s.app
	cfg := s.cfg
	logger := s.logger

	port := fmt.Sprintf(":%d", cfg.Server.Port)
	logger.Info("Server starting on port", zap.String("port", port))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Listen(port); err != nil {
			logger.Error("Failed to start server", zap.Error(err))
			quit <- syscall.SIGTERM
		}
	}()

	// SIGINT or SIGTERM will trigger this graceful shutdown
	<-quit
	logger.Info("Shutting down server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		logger.Error("Error shutting down server", zap.Error(err))
	} else {
		logger.Info("Server stopped gracefully")
	}
}
