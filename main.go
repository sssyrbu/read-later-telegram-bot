package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sssyrbu/save-links-telegram-bot/config"
	"github.com/sssyrbu/save-links-telegram-bot/sendarticles"
	"github.com/sssyrbu/save-links-telegram-bot/storage"
	"github.com/sssyrbu/save-links-telegram-bot/verify"
)

func main() {
	configs := config.LoadConfig()
	botApiToken := configs.Token
	redisClient := storage.LoadRedisClient()

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
			sendarticles.SendArticles(bot, redisClient)
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
			userArticles, err := storage.UserArticles(redisClient, strconv.FormatInt(update.Message.Chat.ID, 10))
			if err != nil {
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
				_, err := storage.InsertArticle(redisClient, strconv.FormatInt(update.Message.Chat.ID, 10), update.Message.Text)
				if err != nil {
					log.Printf("An error occured while executing defaul case: %v", err)
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
