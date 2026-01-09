package ioc

import (
	"dxxproject/config_prepare/app_config"
	"dxxproject/config_prepare/start_config"
	"dxxproject/my/jwt_utils/jwt_user"
	"dxxproject/my/my_logger"
	"dxxproject/my/passwd_util"
	"dxxproject/pkg/gin"
	"dxxproject/pkg/gorm_ok"
	"dxxproject/pkg/nacos_ok"
	"dxxproject/pkg/redis_ok"
	"dxxproject/pkg/snowflake_ok"
	"dxxproject/pkg/zap_ok"
	"github.com/redis/go-redis/v9"
	"github.com/samber/do/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

/*
使用 samber do v2
默认情况下，Article 创建的是单例
*/

func Inject() (injector do.Injector, err error) {

	injector = do.New()

	injectPkg(injector) //注入基础Pkg

	injectModule(injector) //注入服务

	VerifyCode(injector) //注入handler

	gin.Provide(injector) //注入gin

	return
}

func injectPkg(injector do.Injector) {
	//读取本地启动配置
	do.Provide(injector, func(injector do.Injector) (*start_config.StartConfig, error) {
		return start_config.NewStartConfig()
	})

	//nacos instance
	do.Provide(injector, func(injector do.Injector) (*nacos_ok.NacosInstance, error) {
		startConfig := do.MustInvoke[*start_config.StartConfig](injector)
		return nacos_ok.NewNacosInstance(startConfig)
	})

	//app config
	do.Provide(injector, func(injector do.Injector) (*app_config.AppConfig, error) {
		instance := do.MustInvoke[*nacos_ok.NacosInstance](injector)
		return app_config.NewAppConfig(instance)
	})

	//zap logger
	do.Provide(injector, func(injector do.Injector) (*zap.Logger, error) {
		appCfg := do.MustInvoke[*app_config.AppConfig](injector)
		return zap_ok.NewZapLogger(appCfg)
	})

	//user jwt
	do.Provide(injector, func(injector do.Injector) (*jwt_user.UserImpl, error) {
		cfg := do.MustInvoke[*app_config.AppConfig](injector)
		return jwt_user.NewJwtUserImpl(cfg)
	})

	// 雪花算法
	do.Provide(injector, func(injector do.Injector) (*snowflake_ok.SnowflakeIMPL, error) {
		startConfig := do.MustInvoke[*start_config.StartConfig](injector)
		appConfig := do.MustInvoke[*app_config.AppConfig](injector)
		return snowflake_ok.NewSnowFlake(startConfig, appConfig)
	})

	//密码工具
	do.Provide(injector, func(injector do.Injector) (*passwd_util.PasswdUtilImpl, error) {
		return passwd_util.NewPasswordUtil()
	}) //密码加密

	// 自定义 logger
	do.Provide(injector, func(injector do.Injector) (*my_logger.MyLoggerZapImpl, error) {
		zapLogger := do.MustInvoke[*zap.Logger](injector)
		return my_logger.NewMyLogger(zapLogger)
	})

	// gorm db
	do.Provide(injector, func(injector do.Injector) (*gorm.DB, error) {
		appCfg := do.MustInvoke[*app_config.AppConfig](injector)
		return gorm_ok.NewGormDb(appCfg)
	}) //初始化gorm

	// redis
	do.Provide(injector, func(injector do.Injector) (*redis.Client, error) {
		appCfg := do.MustInvoke[*app_config.AppConfig](injector)
		return redis_ok.NewRedisClient(appCfg)
	})

}

func injectModule(injector do.Injector) {
	RateLimit(injector)
	Sms(injector)
	Email(injector)
	User(injector)
	Article(injector)
}
