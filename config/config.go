package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Name string
		Port int
	}
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}
	Logging struct {
		Level string
	}
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