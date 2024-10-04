package redis

import (
	"github.com/redis/go-redis/v9"
	"ministry/config"
	"ministry/pkg/logger"
)

func NewRedisClient(log *logger.Logger, cfg *config.Config) *redis.Client {
	redisConfig := cfg.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	return client
}
