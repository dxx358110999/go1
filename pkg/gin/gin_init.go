package gin

import (
	"dxxproject/config_prepare/app_config"
	"dxxproject/pkg/gin/gin_middleware"
	"dxxproject/pkg/gin/route"
	"github.com/gin-gonic/gin"
	"github.com/samber/do/v2"
)

func NewGinServer(injector do.Injector) (engine *gin.Engine, err error) {
	appConfig := do.MustInvoke[*app_config.Config](injector)

	mode := appConfig.Mode
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	engine = gin.New()

	//注册中间件
	err = useMiddleware(injector, engine)
	if err != nil {
		return
	}

	//注册基本路由
	route.AddBasicRoute(engine)

	// 添加业务路由
	apiGroup := engine.Group("/api")          // 前缀 /api
	route.AddWebRoute(injector, apiGroup)     // /api/user
	route.VerifyCodeRoute(injector, apiGroup) // /api/code

	return
}

func useMiddleware(injector do.Injector, engine *gin.Engine) (err error) {

	errorHandler := gin_middleware.NewErrorHandler() //较为复杂,封装为对象

	//添加中间件
	//Engine.Use(middlewares.RateLimit(2*time.Second, 1) )
	engine.Use(gin_middleware.CorsHandler())         //跨域处理
	engine.Use(errorHandler.ErrorHandler())          //全局错误
	engine.Use(gin_middleware.GinRecovery(true))     //recover
	engine.Use(gin_middleware.RateLimiter(injector)) //限速
	engine.Use(gin_middleware.GinLogger())           //全局日志处理

	//store := cookie.NewStore([]byte("secret"))
	//Engine.Use(sessions.Sessions("mysession", store)) //使用session
	return
}

func Provide(injector do.Injector) {
	do.Provide(injector, NewGinServer)
}
