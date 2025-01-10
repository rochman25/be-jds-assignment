package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

func LoadEnv(path string) {
	// Load .env file
	if err := godotenv.Load(path); err != nil {
		log.Println("No .env file found, using os environment variables")
	}

	// Set Viper to read from environment variables
	viper.AutomaticEnv()
}

func AppVersion() string {
	return viper.GetString("APP_AUTH_VERSION")
}

func AppPort() int {
	return viper.GetInt("APP_AUTH_PORT")
}

func AppEnv() string {
	return viper.GetString("APP_AUTH_ENV")
}

func AppHost() string {
	return viper.GetString("APP_AUTH_HOST")
}

func AppName() string {
	return viper.GetString("APP_AUTH_NAME")
}
