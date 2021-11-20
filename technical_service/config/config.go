package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configure struct {
	App AppConfig
	Database DatabaseConfig
}

type AppConfig struct {
	HOST string `mapstructure:"HOST"`
	PORT int `mapstructure:"PORT"`
}

type DatabaseConfig struct {
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`
	DBHost string `mapstructure:"DB_HOST"`
	DBPort int `mapstructure:"DB_PORT"`
}

var (
	C *Configure
)

func SetupConfiguration() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading foundation file ", err)
		return
	}

	C = new(Configure)

	err := viper.Unmarshal(&C)

	if err != nil {
		log.Fatal("Unable to decode into struct ", err)
		return
	}
}