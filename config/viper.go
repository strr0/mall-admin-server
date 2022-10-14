package config

import (
	"github.com/spf13/viper"
	"log"
)

var settings *viper.Viper

func init() {
	settings = viper.New()
	settings.SetConfigName("application")
	settings.SetConfigType("yml")
	settings.AddConfigPath("./")
	err := settings.ReadInConfig()
	if err != nil {
		log.Fatalf("load config failed: %v", err)
	}
}
