package main

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sssyrbu/read-later-telegram-bot/config"
	"github.com/sssyrbu/read-later-telegram-bot/sendarticles"
	"github.com/sssyrbu/read-later-telegram-bot/storage"
	"github.com/sssyrbu/read-later-telegram-bot/verify"
)

func main() {
	log.Printf("hi world")
	configs := config.LoadConfig()
	botApiToken := configs.Token
	sqliteClient, db_err := storage.InitializeSQLiteDB()
	if db_err != nil {
		log.Fatalf("Failed to create bot: %v", db_err)
	}
	bot, err := tgbotapi.NewBotAPI(botApiToken)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	// Sending an article once in 24 hours
	dailyTicker := time.NewTicker(24 * time.Hour)
	go func() {
		for range dailyTicker.C {
			sendarticles.SendArticles(bot, sqliteClient)
		}
	}()

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		// Checking if message is a start command
		switch update.Message.Command() {
		case "start":
			msg.Text = "Hi! Send me an article (url) and I will remind you to read it later."
			msg.ReplyToMessageID = update.Message.MessageID
		// Command to view articles that user previously saved
		case "view_articles":
			userArticles, err := storage.GetUserArticles(sqliteClient, update.Message.Chat.ID)
			if err != nil {
				msg.Text = "Internal error."
				fmt.Println("An error occured while executing 'view_articles' command:", err)
				return
			}
			msg.Text = "Your saved articles: \n"
			for index, article := range userArticles {
				formattedArticle := fmt.Sprintf("%d. %s \n", index+1, article)
				msg.Text += formattedArticle
			}
		default:
			// Checking if url that was sent by user is valid
			sentUrl := verify.VerifyLink(update.Message.Text)
			if sentUrl {
				db_err := storage.AddArticle(sqliteClient, update.Message.Chat.ID, update.Message.Text)
				if db_err != nil {
					msg.Text = "Internal error."
					log.Printf("An error occured while executing defaul case (db related): %v", err)
				}
				msg.Text = "Article was saved!"
			} else {
				msg.Text = "The link you provided is invalid."
			}
		}

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}
