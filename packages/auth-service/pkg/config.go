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
	return viper.GetString("APP_VERSION")
}

func ApiSecretKey() string {
	return viper.GetString("API_KEY_SECRET")
}

func AppPort() int {
	return viper.GetInt("APP_PORT")
}

func AppEnv() string {
	return viper.GetString("APP_ENV")
}

func AppHost() string {
	return viper.GetString("APP_HOST")
}
