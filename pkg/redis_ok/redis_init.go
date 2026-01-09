package redis_ok

import (
	"context"
	"dxxproject/config_prepare/app_config"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(appCfg *app_config.AppConfig) (redisClient *redis.Client, err error) {
	cfg := appCfg.RedisConfig
	redisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	ctx := context.Background()
	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		return
	}
	return
}
