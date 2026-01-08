package gin_server

import (
	"dxxproject/internal/handlers"
	"dxxproject/my/jwt_utils/jwt_gin_middlerware"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/samber/do/v2"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func AddBasicRoute(engine *gin.Engine) {
	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pprof.Register(engine) // 注册pprof相关路由
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

}

func addWebRoute(injector do.Injector, prefixGroup *gin.RouterGroup) {
	userHandlers := do.MustInvoke[*handlers.UserHandler](injector)

	userGroup := prefixGroup.Group("/user")
	userGroup.POST("/signup", userHandlers.Signup)
	userGroup.POST("/login", userHandlers.Login)

	//后续的接口需要认证
	userGroup.Use(jwt_gin_middlerware.JwtAuthMiddleware(injector))
	userGroup.POST("/profile", userHandlers.Profile)

}

func addVerifyCodeRoute(injector do.Injector, prefixGroup *gin.RouterGroup) {
	vcHandlers := do.MustInvoke[*handlers.VerifyCodeHandlers](injector)

	group := prefixGroup.Group("/code")
	group.POST("/registerCodeSendEmail", vcHandlers.SendUserRegisterCodeByEmail)
	group.POST("/registerCodeSendSms", vcHandlers.SendUserRegisterCodeBySms)

	//后续的接口需要认证
	//group.Use(gin_middleware.JwtAuthMiddleware(injector))
}
