//go:build wireinject
// +build wireinject

package main

import (
	"github.com/JeongJaeSoon/go-auth/cmd/server"
	"github.com/JeongJaeSoon/go-auth/config"
	"github.com/JeongJaeSoon/go-auth/internal/logging"
	"github.com/google/wire"
)

func InitializeServer() (*server.Server, error) {
	wire.Build(
		config.LoadConfig,
		config.ProvideLoggingConfig,
		logging.InitLogger,
		server.NewServer,
	)
	return nil, nil
}
