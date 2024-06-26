package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ENVIRONMENT            string        `mapstructure:"ENVIRONMENT"`
	DB_SOURCE              string        `mapstructure:"DB_SOURCE"`
	MIGRATION_URL          string        `mapstructure:"MIGRATION_URL"`
	REDIS_ADDRESS          string        `mapstructure:"REDIS_ADDRESS"`
	HTTP_SERVER_ADDRESS    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPC_SERVER_ADDRESS    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TOKEN_SYMMETRIC_KEY    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	ACCESS_TOKEN_DURATION  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	REFRESH_TOKEN_DURATION time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	EMAIL_SENDER_NAME      string        `mapstructure:"EMAIL_SENDER_NAME"`
	EMAIL_SENDER_ADDRESS   string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EMAIL_SENDER_PASSWORD  string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
