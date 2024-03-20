// This file is responsible for the logic of sending a random article to all users in the exising db.
package sendarticles

import (
	"math/rand"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sssyrbu/read-later-telegram-bot/storage"
)

func SendArticles(bot *tgbotapi.BotAPI, dbClient *sql.DB) {
	allUserIDs, _ := storage.GetUserIDs(dbClient)
	for _, userID := range allUserIDs {
		articlesArr, _ := storage.GetUserArticles(dbClient, userID)
		if len(articlesArr) > 0 {
			randomArticle := articlesArr[rand.Intn(len(articlesArr))]
			msg := tgbotapi.NewMessage(userID, "Check out this article today:\n"+randomArticle)
			if _, err := bot.Send(msg); err != nil {
				continue
			}
		}
	}
}
