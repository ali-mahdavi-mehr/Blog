package database

import (
	"github.com/go-redis/redis/v8"
	"os"
)

func createRedisConnection() func() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "",
		DB:       0,
	})
	return func() *redis.Client {
		return client
	}
}

func GetRedisClient() *redis.Client {
	client := createRedisConnection()
	return client()
}
