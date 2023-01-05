package lib

import (
	"github.com/spf13/viper"
)

type Env struct {
	LogLevel    string `mapstructure:"LOG_LEVEL"`
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENVIRONMENT"`

	SentryDSN string `mapstructure:"SENTRY_DSN"`

	TimeZone string `mapstructure:"TIMEZONE"`
}

// add global envs here
var globalEnv = Env{}

func GetEnv() Env {
	return globalEnv
}

func NewEnv(logger Logger) *Env {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal("cannot read cofiguration", err)
	}

	viper.SetDefault("TIMEZONE", "UTC")

	err = viper.Unmarshal(&globalEnv)
	if err != nil {
		logger.Fatal("environment cant be loaded: ", err)
	}

	return &globalEnv
}
