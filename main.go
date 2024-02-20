package main

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sssyrbu/save-links-telegram-bot/config"
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

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		// Checking if message is a start command
		switch update.Message.Command() {
		case "start":
			msg.Text = "first message output"
			msg.ReplyToMessageID = update.Message.MessageID
		// Command to view articles that user previously saved
		case "view_articles":
			fmt.Println("we are in the view articles section")
			userArticles, err := storage.UserArticles(redisClient, strconv.FormatInt(update.Message.Chat.ID, 10))
			if err != nil {
				fmt.Println("an error occured:", err)
				return
			}
			msg.Text = "Your saved articles: \n"
			for index, article := range userArticles {
				formattedArticle := fmt.Sprintf("%d. %s \n", index+1, article)
				msg.Text += formattedArticle
			}
			// msg.ReplyToMessageID = update.Message.MessageID
		case "dojka":
			msg.Text = "Я люблю Аню 🩷"
		default:
			// Checking if url that was sent by user is valid
			sentUrl := verify.VerifyLink(update.Message.Text)
			if sentUrl {
				result, err := storage.InsertArticle(redisClient, strconv.FormatInt(update.Message.Chat.ID, 10), update.Message.Text)
				if err != nil {
					fmt.Println("an error occured:", err)
				} else {
					fmt.Println("success result:", result)
				}
				msg.Text = "valid link"
			} else {
				msg.Text = "invalid link"
			}
		}

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}
