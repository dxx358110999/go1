package gin_middleware

import (
	"dxxproject/agreed/dto"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func CorsHandler() gin.HandlerFunc {
	return cors.New(cors.Config{
		//AllowOrigins: []string{"https://github.com"},
		//AllowMethods: []string{"PUT", "POST", "PATCH", "OPTIONS", "GET"},
		AllowHeaders: []string{"Origin",
			dto.HeaderAccessToken,
			dto.HeaderRefreshToken,
			"Content-Type", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", dto.HeaderAccessToken, dto.HeaderRefreshToken},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			//return origin == "https://github.com"
			return true
		},
		MaxAge: 12 * time.Hour,
	})
}
