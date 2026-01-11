package ioc

import (
	"dxxproject/config_prepare/app_config"
	"dxxproject/config_prepare/start_config"
	"dxxproject/my/id_generator/snow_impl"
	"dxxproject/my/jwt_utils/jwt_user"
	"dxxproject/my/my_logger"
	"dxxproject/my/pwd_util/pwd_impl"
	"dxxproject/pkg/gorm_ok"
	"dxxproject/pkg/nacos_ok"
	"dxxproject/pkg/redis_ok"
	"dxxproject/pkg/zap_ok"
	"github.com/redis/go-redis/v9"
	"github.com/samber/do/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func injectPkg(injector do.Injector) {
	//读取本地启动配置
	do.Provide(injector, func(injector do.Injector) (*start_config.Config, error) {
		return start_config.NewStartConfig()
	})

	//nacos instance
	do.Provide(injector, func(injector do.Injector) (*nacos_ok.NacosInstance, error) {
		startConfig := do.MustInvoke[*start_config.Config](injector)
		return nacos_ok.NewNacosInstance(startConfig)
	})

	//app config
	do.Provide(injector, func(injector do.Injector) (*app_config.Config, error) {
		instance := do.MustInvoke[*nacos_ok.NacosInstance](injector)
		return app_config.NewAppConfig(instance)
	})

	//zap logger
	do.Provide(injector, func(injector do.Injector) (*zap.Logger, error) {
		appCfg := do.MustInvoke[*app_config.Config](injector)
		return zap_ok.NewZapLogger(appCfg)
	})

	//user jwt
	do.Provide(injector, func(injector do.Injector) (*jwt_user.UserImpl, error) {
		cfg := do.MustInvoke[*app_config.Config](injector)
		return jwt_user.NewJwtUserImpl(cfg)
	})

	// 雪花算法
	do.Provide(injector, func(injector do.Injector) (*snow_impl.SnowflakeIMPL, error) {
		startConfig := do.MustInvoke[*start_config.Config](injector)
		appConfig := do.MustInvoke[*app_config.Config](injector)
		return snow_impl.NewSnowFlake(startConfig, appConfig)
	})

	//密码工具
	do.Provide(injector, func(injector do.Injector) (*pwd_impl.PwdImpl, error) {
		return pwd_impl.NewPasswordUtil()
	}) //密码加密

	// 自定义 logger
	do.Provide(injector, func(injector do.Injector) (*my_logger.MyLoggerZapImpl, error) {
		zapLogger := do.MustInvoke[*zap.Logger](injector)
		return my_logger.NewMyLogger(zapLogger)
	})

	// gorm db
	do.Provide(injector, func(injector do.Injector) (*gorm.DB, error) {
		appCfg := do.MustInvoke[*app_config.Config](injector)
		return gorm_ok.NewGormDb(appCfg)
	}) //初始化gorm

	// redis
	do.Provide(injector, func(injector do.Injector) (*redis.Client, error) {
		appCfg := do.MustInvoke[*app_config.Config](injector)
		return redis_ok.NewRedisClient(appCfg)
	})

}
