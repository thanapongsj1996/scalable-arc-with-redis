package redis

import (
	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,
	})

	return rdb
}
