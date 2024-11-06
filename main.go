/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/JeongJaeSoon/go-auth/cmd/server"
	"github.com/JeongJaeSoon/go-auth/config"
	"github.com/JeongJaeSoon/go-auth/internal/logging"
	"github.com/JeongJaeSoon/go-auth/internal/otel"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/baggage"
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
	otel.InitTracer()

	// <<-- Sample code for testing for otel auto instrumentation -->>
	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}

	ctx := baggage.ContextWithoutBaggage(context.Background())

	req, _ := http.NewRequestWithContext(ctx, "GET", "http://sumologic.com", nil)

	fmt.Printf("Sending request...\n")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response status code: %v\n", res.Status)
	fmt.Printf("Waiting for few seconds to export spans ...\n\n")
	time.Sleep(10 * time.Second)
	// <<-- Sample code for testing for otel auto instrumentation -->>

	server.StartServer(cfg)
}
