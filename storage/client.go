package storage

import (
	"github.com/KofClubs/siwa-back-end/config"
	"github.com/MonteCarloClub/log"
)

func InitClient() {
	config.RedisClient = initRedisClient()
	log.Info("redis client inited")
}
