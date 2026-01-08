package redis_client

import (
	"context"
	"dxxproject/config_prepare/app_config"
	"fmt"
	"github.com/samber/do/v2"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(injector do.Injector) (redisClient *redis.Client, err error) {
	cfg := do.MustInvoke[*app_config.AppConfig](injector).RedisConfig
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

func Provide(injector do.Injector) {
	do.Provide(injector, NewRedisClient)
}
