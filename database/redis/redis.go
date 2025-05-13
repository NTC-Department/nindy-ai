package redis

import (
	"context"
	"fmt"
	"log"
	"nindychat/utils"

	"github.com/go-redis/redis/v8"
)

var redisInstance *redis.Client
var ctx = context.Background()

func InitializeRedis() {
	log.Println("===== Initialize Redis =====")
	client, err := connectRedis()
	if err != nil {
		panic(err)
	}

	redisInstance = client
	log.Printf("✓ connected to Redis: %s:%s\n", utils.GetEnv("REDIS_HOST"), utils.GetEnv("REDIS_PORT"))
}

func GetRedisInstance() *redis.Client {
	return redisInstance
}

func connectRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", utils.GetEnv("REDIS_HOST"), utils.GetEnv("REDIS_PORT")),
		Password: utils.GetEnv("REDIS_PASSWORD"),
		DB:       utils.GetIntEnv("REDIS_DB"),
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Println("Error while connecting to Redis")
		return nil, err
	}

	return client, nil
}

func CloseRedis() {
	if redisInstance != nil {
		if err := redisInstance.Close(); err != nil {
			log.Println("Error while closing Redis connection")
			return
		}
	}

	redisInstance = nil
	log.Println("✓ Redis connection closed")
}
