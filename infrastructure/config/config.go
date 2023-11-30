package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type SystemConfig struct {
	Port string `mapstructure:"PORT"`
}

type DatabaseConfig struct {
	DatabaseURI string `mapstructure:"MYSQL_URI"`
}

type AppConfigData struct {
	SystemConfig   SystemConfig
	DatabaseConfig DatabaseConfig
}

var AppConfig *AppConfigData

func init() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading .env file: %s\n", err)
	}

	systemConfig := SystemConfig{
		Port: viper.GetString("PORT"),
	}

	databaseConfig := DatabaseConfig{
		DatabaseURI: viper.GetString("MYSQL_URI"),
	}

	AppConfig = &AppConfigData{
		SystemConfig:   systemConfig,
		DatabaseConfig: databaseConfig,
	}
}
