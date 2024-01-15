package main

import (
    "github.com/sssyrbu/save-links-telegram-bot/config"
    "log"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
    botApiToken := config.LoadConfig().Token
    bot, err := tgbotapi.NewBotAPI(botApiToken)
    if err != nil {
        log.Fatalf("Failed to create bot: %v", err)    
    }

    log.Printf("Authorized on account %s", bot.Self.UserName)

    // bot.Debug = true
    updateConfig := tgbotapi.NewUpdate(0)
    updateConfig.Timeout = 30
    updates := bot.GetUpdatesChan(updateConfig)

    for update := range updates {
        if update.Message == nil {
            continue
        }
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
        // Checking if message is a start command
        switch update.Message.Command() {
        case "start":
            msg.Text = "hi there! send me an article and I will save it to my local db. I will be randomly sending you an article every day."
        case "sayhi":
            msg.Text = "Hi :)"
        }
        msg.ReplyToMessageID = update.Message.MessageID

        if _, err := bot.Send(msg); err != nil {
            panic(err)
        }
    }

}

