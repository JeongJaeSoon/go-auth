package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Set environment variables for testing
	os.Setenv("CONFIG_PATH", "./")

	os.Setenv("ENV_SERVER_NAME", "test-auth-server")
	os.Setenv("ENV_SERVER_PORT", "9999")
	os.Setenv("ENV_DATABASE_HOST", "test-db-host")
	os.Setenv("ENV_DATABASE_PORT", "5432")
	os.Setenv("ENV_DATABASE_USER", "test-user")
	os.Setenv("ENV_DATABASE_PASSWORD", "test-password")
	os.Setenv("ENV_DATABASE_NAME", "test-db")
	os.Setenv("ENV_LOGGING_LEVEL", "debug")

	viper.SetConfigFile("config.yaml")

	// Load configuration
	config, err := LoadConfig()
	assert.NoError(t, err, "Error occurred while loading configuration.")

	// Test configuration values
	assert.Equal(t, "test-auth-server", config.Server.Name, "Server name should be test-auth-server")
	assert.Equal(t, 9999, config.Server.Port, "Server port should be 9999")
	assert.Equal(t, "test-db-host", config.Database.Host, "Database host should be test-db-host")
	assert.Equal(t, 5432, config.Database.Port, "Database port should be 5432")
	assert.Equal(t, "test-user", config.Database.User, "Database user should be test-user")
	assert.Equal(t, "test-password", config.Database.Password, "Database password should be test-password")
	assert.Equal(t, "test-db", config.Database.Name, "Database name should be test-db")
	assert.Equal(t, "debug", config.Logging.Level, "Logging level should be debug")
}
