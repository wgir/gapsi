package config

import (
	"strings"

	"github.com/spf13/viper"
)

// Config holds all the configuration for the application
type Config struct {
	AppPort             string `mapstructure:"APP_PORT"`
	ProjectID           string `mapstructure:"PROJECT_ID"`
	FirestoreEmulatorHost string `mapstructure:"FIRESTORE_EMULATOR_HOST"`
	LogLevel            string `mapstructure:"LOG_LEVEL"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	// Replace dot with underscore for nested configs if needed, though not used here
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Default values
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("PROJECT_ID", "todo-project")

	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return
		}
	}

	err = viper.Unmarshal(&config)
	return
}
