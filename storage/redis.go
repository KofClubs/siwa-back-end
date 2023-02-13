package storage

import (
	"github.com/KofClubs/siwa-back-end/config"
	"github.com/go-redis/redis/v8"
)

func initRedisClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.RedisClientAddr,
	})
	return redisClient
}
