package ioc

import (
	"dxxproject/config_prepare/app_config"
	"dxxproject/config_prepare/start_config"
	"dxxproject/internal/handlers"
	"dxxproject/internal/svc/email_svc"
	"dxxproject/internal/svc/rate_limit_svc"
	"dxxproject/internal/svc/sms_svc"
	"dxxproject/internal/svc/user_svc"
	"dxxproject/internal/svc/verify_code_svc"
	"dxxproject/my/jwt_utils/jwt_user"
	"dxxproject/my/my_logger"
	"dxxproject/my/passwd_util"
	"dxxproject/pkg/gin/gin_server"
	"dxxproject/pkg/gorm_db"
	"dxxproject/pkg/nacos_ok"
	"dxxproject/pkg/redis_client"
	"dxxproject/pkg/snowflake_ok"
	"dxxproject/pkg/zap_ok"
	"github.com/samber/do/v2"
)

/*
使用 samber do v2
默认情况下，Provide 创建的是单例
*/

func Inject() (injector do.Injector, err error) {

	injector = do.New()

	injectPkg(injector) //注入基础Pkg

	injectSvc(injector) //注入服务

	handlers.Provide(injector) //注入handler

	gin_server.Provide(injector) //注入gin

	return
}

func injectPkg(injector do.Injector) {

	start_config.Provide(injector) //启动配置
	nacos_ok.Provide(injector)     //nacos instance
	app_config.Provide(injector)   //app config
	zap_ok.Provide(injector)       //zap

	jwt_user.Provide(injector)     //user jwt
	snowflake_ok.Provide(injector) //雪花算法
	passwd_util.Provide(injector)  //密码加密
	my_logger.Provide(injector)    //自定义的logger

	gorm_db.Provide(injector)      //初始化gorm
	redis_client.Provide(injector) //初始化redis连接

}

func injectSvc(injector do.Injector) {
	rate_limit_svc.Provide(injector)
	sms_svc.Provide(injector)
	email_svc.Provide(injector)
	verify_code_svc.Provide(injector)
	user_svc.Provide(injector)

}
