package rate_limit

import (
	"context"
	"dxxproject/my/my_err"
	"github.com/go-redis/redis_rate/v10"
	"github.com/pkg/errors"
	"time"
)

type RateLimitSvc struct {
	limiter *redis_rate.Limiter
}

func (r *RateLimitSvc) GetAllow(ctx context.Context, key string, rate int, period time.Duration) (err error) {
	limitRule := redis_rate.Limit{
		Rate:   rate,
		Burst:  rate,
		Period: period,
	} //使用结构体直接创建,更灵活的时间范围

	res, err := r.limiter.Allow(ctx, key, limitRule)
	if err != nil {
		return errors.Wrap(err, "获取允许失败")
	} else {
		if res.Allowed == 0 {
			return my_err.ErrRateLimited
		}
	}
	return
}

func (r *RateLimitSvc) LimitClear(ctx context.Context, key string) (err error) {
	err = r.limiter.Reset(ctx, key)
	if err != nil {
		return err
	}
	return
}

func NewRateLimit(limiter *redis_rate.Limiter) (*RateLimitSvc, error) {
	return &RateLimitSvc{
		limiter: limiter,
	}, nil
}
