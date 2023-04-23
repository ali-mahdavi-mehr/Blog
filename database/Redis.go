package database

import (
	"github.com/go-redis/redis/v8"
	"os"
)

func CreateConnection() func() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_host"),
		Password: "",
		DB:       0,
	})
	return func() *redis.Client {
		return client
	}
}

func GetRedisClient() *redis.Client {
	client := CreateConnection()
	return client()
}
