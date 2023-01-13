package lib

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	LogLevel    string `mapstructure:"LOG_LEVEL"`
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENVIRONMENT"`

	SentryDSN string `mapstructure:"SENTRY_DSN"`

	TimeZone  string `mapstructure:"TIMEZONE"`
	JWTSecret string `mapstructure:"JWT_SECRET"`
}

// add global envs here
var globalEnv = Env{}

func GetEnv() Env {
	return globalEnv
}

func NewEnv() Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		// TODO: default provider only for dockerfile
		env.Environment = "dev"
		env.LogLevel = "info"
		env.ServerPort = "8000"
		return env
		// log.Fatal("☠️ cannot read configuration")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
	}

	return env
}
