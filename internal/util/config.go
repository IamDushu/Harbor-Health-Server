package util

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	ServerAddress        string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	AuthTokenExpiry      time.Duration `mapstructure:"AUTH_TOKEN_EXPIRY"`
	TwillioAccountSID    string        `mapstructure:"TWILLIO_ACCOUNT_SID"`
	TwillioAuthToken     string        `mapstructure:"TWILLIO_AUTH_KEY"`
	BrevoAPIKey          string        `mapstructure:"BREVO_API_KEY"`
	TemplateID           string        `mapstructure:"TEMPLATE_ID"`
	StreamApiKey         string        `mapstructure:"STREAM_API_KEY"`
	StreamSecretKey      string        `mapstructure:"STREAM_SECRET_KEY"`
}

// LoadConfig reads configuration from file or env variables
func LoadConfig() (config Config, err error) {
	viper.AutomaticEnv()

	// Default PORT fallback
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	viper.SetDefault("SERVER_ADDRESS", fmt.Sprintf("0.0.0.0:%s", port))

	// Manually get all fields
	config.DBDriver = viper.GetString("DB_DRIVER")
	config.DBSource = viper.GetString("DB_SOURCE")
	config.ServerAddress = viper.GetString("SERVER_ADDRESS")
	config.TokenSymmetricKey = viper.GetString("TOKEN_SYMMETRIC_KEY")
	config.TwillioAccountSID = viper.GetString("TWILLIO_ACCOUNT_SID")
	config.TwillioAuthToken = viper.GetString("TWILLIO_AUTH_KEY")
	config.BrevoAPIKey = viper.GetString("BREVO_API_KEY")
	config.TemplateID = viper.GetString("TEMPLATE_ID")
	config.StreamApiKey = viper.GetString("STREAM_API_KEY")
	config.StreamSecretKey = viper.GetString("STREAM_SECRET_KEY")

	// Parse durations safely
	config.AccessTokenDuration, err = time.ParseDuration(viper.GetString("ACCESS_TOKEN_DURATION"))
	if err != nil {
		return
	}
	config.RefreshTokenDuration, err = time.ParseDuration(viper.GetString("REFRESH_TOKEN_DURATION"))
	if err != nil {
		return
	}
	config.AuthTokenExpiry, err = time.ParseDuration(viper.GetString("AUTH_TOKEN_EXPIRY"))
	if err != nil {
		return
	}

	return
}
