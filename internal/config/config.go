package config

import (
	"github.com/AlikhanIT/hotel-api/internal/logger"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct{ Port int }
	DB     struct {
		Host, User, Password, Name string
		Port                       int
	}
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	logger.Info("config loaded")
	return &cfg
}
