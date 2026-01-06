package cache

import (
	"context"
	"dxxproject/my/my_err"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/samber/do/v2"
	"time"
)

type VerifyCode struct {
	redisClient *redis.Client
}

func (r *VerifyCode) VerifyCodeGet(ctx context.Context, codeKey string) (result string, err error) {
	result, err = r.redisClient.Get(ctx, codeKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			err = my_err.ErrorDataNotExist
			return
		}
		return
	}
	return
}

func (r *VerifyCode) VerifyCodeSet(ctx context.Context, codeKey string, code string, expire time.Duration) (err error) {
	err = r.redisClient.Set(ctx, codeKey, code, expire).Err()
	if err != nil {
		return
	}
	return
}

func NewVerifyCodeCache(injector do.Injector) (*VerifyCode, error) {
	redisClient := do.MustInvoke[*redis.Client](injector)
	vc := &VerifyCode{
		redisClient: redisClient,
	}
	return vc, nil
}
