package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Logging  LoggingConfig
}

type ServerConfig struct {
	Name string
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type LoggingConfig struct {
	Level  string
	Format string
}

func LoadConfig() (*Config, error) {
	configFile := "config"
	configEnv := os.Getenv("CONFIG_ENV")
	configPath := os.Getenv("CONFIG_PATH")

	if configEnv == "production" {
		configFile = "config.production"
	}

	if configPath == "" {
		configPath = "./config"
	}

	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()

	viper.SetEnvPrefix("ENV")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %v", err)
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
