/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/JeongJaeSoon/go-auth/cmd/server"
	"github.com/JeongJaeSoon/go-auth/config"
	"github.com/JeongJaeSoon/go-auth/internal/logging"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	logging.InitLogger(cfg.Logging)
	defer func() {
		if r := recover(); r != nil {
			logging.Logger.Error("Application panicked: %v", zap.Any("error", r))
		}

		logging.Logger.Sync()
	}()

	logging.Logger.Info("Application starting")
	server.StartServer(cfg)
}
