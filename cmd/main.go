package main

import (
    "os"
    "log"
    "path/filepath"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "github.com/joho/godotenv"
)

func main() {
    // Get the current working directory
    dir, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }

    envPath := filepath.Join(dir, "..", ".env")

    err = godotenv.Load(envPath)
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    botApiToken := os.Getenv("BOT_API_TOKEN")
    bot, err := tgbotapi.NewBotAPI(botApiToken)
    if err != nil {
        panic(err)
    }

    bot.Debug = true
}

