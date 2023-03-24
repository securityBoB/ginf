package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type AppConfig struct {
	Port string `mapstructure:"port"`
}

var appConfig AppConfig

func LoadAppConfig() {
	if _, err := toml.DecodeFile("config/app.toml", &appConfig); err != nil {
		log.Fatal(err)
	}
}

func GetAppConfig() AppConfig {
	return appConfig
}
