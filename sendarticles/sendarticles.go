// This file is responsible for the logic of sending a random article to all users in the exising db.
package sendarticles

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/go-redis/redis"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sssyrbu/read-later-telegram-bot/storage"
)

func SendArticles(bot *tgbotapi.BotAPI, redisClient *redis.Client) {
	allUserIDs, err := storage.LoadUserIDs(redisClient)
	if err != nil {
		log.Printf("Error getting user IDs from Redis: %v", err)
	}
	for _, userID := range allUserIDs {
		chatID, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			log.Printf("Error parsing chat ID for user %s: %v", userID, err)
			continue
		}
		articlesArr, err := storage.UserArticles(redisClient, userID)
		if err != nil {
			log.Printf("Error retrieving user's articles %s: %v", userID, err)
			continue
		}
		if len(articlesArr) > 0 {
			randomArticle := articlesArr[rand.Intn(len(articlesArr))]
			msg := tgbotapi.NewMessage(chatID, "Check out this article today:\n"+randomArticle)
			if _, err := bot.Send(msg); err != nil {
				log.Printf("Error sending message to user %s: %v", userID, err)
			}
		}
	}
}
