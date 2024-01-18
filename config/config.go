// Read config from the .env file in the main directory
package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Token string
	Addr  string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &Config{
		Token: os.Getenv("BOT_API_KEY"),
		Addr:  os.Getenv("REDIS_ADDR"),
	}
}
