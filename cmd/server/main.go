package server

import (
	"fmt"
	"log"

	"github.com/JeongJaeSoon/go-auth/config"
	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	port := fmt.Sprintf(":%d", config.Server.Port)
	log.Printf("Server is running on port %s", port)

	if err := app.Listen(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
