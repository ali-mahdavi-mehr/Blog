package database

import (
	"github.com/go-redis/redis/v8"
	"os"
)

var redisConnection func() *redis.Client

func createRedisConnection() func() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "",
		DB:       0,
	})
	return func() *redis.Client {
		return redisClient
	}
}

func GetRedisClient() *redis.Client {
	if redisConnection == nil {
		redisConnection = createRedisConnection()
	}
	return redisConnection()
}
