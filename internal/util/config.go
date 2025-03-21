package util

import (
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
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") // json, xml

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
