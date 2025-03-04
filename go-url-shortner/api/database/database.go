package database

import (
	"context"
	"os"
	// Redis client for Go (v8). It allows your application to interact with a Redis database.
	"github.com/go-redis/redis/v8"
)
// REDIS IS A KEY VALUE PAIR DATABASE

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client{
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB: dbNo,
	})

	return rdb
}