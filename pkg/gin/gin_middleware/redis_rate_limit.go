package gin_middleware

import (
	"dxxproject/my/my_err"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/samber/do/v2"
)

// 接口请求速率配置, 建议放入redis/数据库同步本地缓存.
var exampleLimiterMap = map[string]redis_rate.Limit{
	"/api/user/profile": redis_rate.PerMinute(10),
}

func RateLimiter(injector do.Injector) gin.HandlerFunc {
	redisClient := do.MustInvoke[*redis.Client](injector)
	limiter := redis_rate.NewLimiter(redisClient)

	return func(ctx *gin.Context) {
		var uri = ctx.Request.URL.Path
		uriLimit, ok := exampleLimiterMap[uri]

		if ok {
			result, err := limiter.Allow(ctx, uri, uriLimit)
			if err != nil {
				ctx.Error(errors.Wrap(err, "获取限速允许失败"))
				ctx.Next()
			} else {
				fmt.Println("uri", uri, "allowed", result.Allowed, "remaining", result.Remaining)

				if result.Allowed == 0 {
					ctx.Error(my_err.ErrRateLimited)
					ctx.Abort()
				}

				ctx.Next()
			}
		} else {
			ctx.Next()
		}
	}
}
