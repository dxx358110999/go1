package gin_server

import (
	"dxxproject/config_prepare/app_config"
	"dxxproject/pkg/gin/gin_middleware"
	"dxxproject/pkg/gin/gin_translator"
	"github.com/gin-gonic/gin"
	"github.com/samber/do/v2"
)

func NewGinServer(injector do.Injector) (engine *gin.Engine, err error) {
	appConfig := do.MustInvoke[*app_config.AppConfig](injector)

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
	AddBasicRoute(engine)

	// 添加业务路由
	apiGroup := engine.Group("/api")       // 前缀 /api
	addWebRoute(injector, apiGroup)        // /api/user
	addVerifyCodeRoute(injector, apiGroup) // /api/code

	return
}

func useMiddleware(injector do.Injector, engine *gin.Engine) (err error) {
	//翻译器
	myTrans := &gin_translator.MyTranslator{}
	err = myTrans.Init("zh")
	if err != nil {
		return
	}

	//gin中间件
	errorHandler := gin_middleware.NewErrorHandler(myTrans.Translator)

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
