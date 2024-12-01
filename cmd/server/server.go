package server

import (
	"github.com/JeongJaeSoon/go-auth/config"
	"github.com/JeongJaeSoon/go-auth/internal/generated"
	"github.com/JeongJaeSoon/go-auth/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fibermiddleware "github.com/oapi-codegen/fiber-middleware"
	"go.uber.org/zap"
)

type Server struct {
	app           *fiber.App
	cfg           *config.Config
	logger        *zap.Logger
	healthHandler *handler.HealthHandler
}

func NewServer(cfg *config.Config, zapLogger *zap.Logger) *Server {
	app := fiber.New(fiber.Config{
		AppName: cfg.Server.Name,
	})

	// Add middleware
	app.Use(logger.New())

	// Add OpenAPI validation middleware
	swagger, err := generated.GetSwagger()
	if err != nil {
		zapLogger.Fatal("Failed to load OpenAPI spec", zap.Error(err))
	}
	swagger.Servers = nil // Disable server validation
	app.Use(fibermiddleware.OapiRequestValidator(swagger))

	healthHandler := handler.NewHealthHandler(zapLogger)

	return &Server{
		app:           app,
		cfg:           cfg,
		logger:        zapLogger,
			healthHandler: healthHandler,
	}
}

func (s *Server) setupRoutes() {
	// Health check endpoint
	s.app.Get("/health", s.healthHandler.Check)

	// Future authentication endpoints will be added here
	// auth := s.app.Group("/auth")
	// auth.Post("/login", s.authHandler.Login)
	// auth.Post("/register", s.authHandler.Register)
	// etc...
}

func (s *Server) GetApp() *fiber.App {
	return s.app
}
