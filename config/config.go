// Read config from the .env file in the main directory
package config

import (
    "os"
    "log"
    "github.com/joho/godotenv"
)

type Config struct {
    Token string
}

func LoadConfig() *Config {
    err := godotenv.Load(".env")

    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    return &Config{Token: os.Getenv("BOT_API_KEY")}
}
