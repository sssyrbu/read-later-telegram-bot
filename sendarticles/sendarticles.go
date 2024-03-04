// Every 24 hours a message with articles is being sent to each user.
// This file is responsible for that logic.
package sendarticles

import (
	"github.com/go-redis/redis"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendArticles(bot *tgbotapi.BotAPI, redisClient *redis.Client) {

}
