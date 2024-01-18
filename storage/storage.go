// Database manipulations logic
package storage

import (
	"github.com/go-redis/redis"
	"github.com/sssyrbu/save-links-telegram-bot/config"
)

func LoadRedisClient() *redis.Client {
	config_addr := config.LoadConfig().Addr

	redisClient := redis.NewClient(&redis.Options{
		Addr:     config_addr,
		Password: "",
		DB:       0,
	})

	return redisClient
}

func InsertArticle(client *redis.Client, key, value string) (int, error) {
	result, err := client.SAdd(key, value).Result()
	if err != nil {
		return 0, err
	}
	return int(result), nil

}
