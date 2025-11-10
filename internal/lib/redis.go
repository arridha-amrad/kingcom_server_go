package lib

import (
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	*redis.Client
}

func NewRedis(env *Env, logger *Logger) *Redis {
	opt, err := redis.ParseURL(env.RedisUrl)
	if err != nil {
		logger.Fatal("Failed to parse url redis")
	}
	rdb := redis.NewClient(opt)
	logger.Info("Redis connection pool established")
	return &Redis{
		Client: rdb,
	}
}
