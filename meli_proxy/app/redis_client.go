package app

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func InitRedis() {
	config := LoadConfig()

	RedisClient = redis.NewClient(&redis.Options{
		Addr: config.RedisHost + ":" + config.RedisPort,
	})
	if err := RedisClient.Ping(RedisCtx).Err(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}

var RedisClient *redis.Client
var RedisCtx = context.Background()
