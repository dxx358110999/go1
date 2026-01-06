package ioc

import (
	"dxxproject/config_prepare/app_config"
	"dxxproject/config_prepare/start_config"
	"dxxproject/main2/basic_info"
	"dxxproject/my/jwt_utils/jwt_user"
	"dxxproject/my/my_logger"
	"dxxproject/my/passwd_util"
	"dxxproject/pkg/gin/gin_server"
	"dxxproject/pkg/gorm_db"
	"dxxproject/pkg/nacos_ok"
	"dxxproject/pkg/redis_client"
	"dxxproject/pkg/sf_utils"
	"dxxproject/pkg/zap_ok"
	"github.com/samber/do/v2"
)

/*
使用 samber do v2
默认情况下，Provide 创建的是单例
*/

func Inject() (injector do.Injector, err error) {

	injector = do.New()

	/*
		基础信息
		包含:
		本机IP
	*/
	do.Provide(injector, basic_info.NewBasicInfo) //基本信息

	do.Provide(injector, start_config.NewStartConfig) //读取本地启动配置

	do.Provide(injector, nacos_ok.NewNacosInstance) //nacos instance

	do.Provide(injector, app_config.NewAppConfig) //app config

	do.Provide(injector, zap_ok.NewZapLogger) //zap

	do.Provide(injector, jwt_user.NewJwtUserImpl) //user jwt

	do.Provide(injector, sf_utils.NewSnowFlake) //雪花算法
	err = do.As[*sf_utils.SnowflakeIMPL, sf_utils.SnowflakeIF](injector)

	do.Provide(injector, passwd_util.NewPasswordUtil) //密码加密

	do.Provide(injector, my_logger.NewMyLogger) //自定义的logger
	err = do.As[*my_logger.MyLoggerZapImpl, my_logger.MyLoggerIF](injector)

	/*
		sms
	*/
	//err, aliSms := ali_sms.NewAliSms(appConfig.AliSms) //没有处理错误

	//初始化sqlx
	//if err = dao2.SqlxInit(AppConfig.MysqlConfig); err != nil {
	//	fmt.Printf("init mysql failed, err:%v\n", err)
	//	return
	//}
	//defer dao2.DbClose() // 程序退出关闭数据库连接

	do.Provide(injector, gorm_db.NewGormDb)           //初始化gorm
	do.Provide(injector, redis_client.NewRedisClient) //初始化redis连接

	err = injectMod(injector)
	if err != nil {
		return nil, err
	}
	do.Provide(injector, gin_server.NewGinServer)

	return
}

func injectMod(injector do.Injector) (err error) {
	err = smsProviderIoc(injector)
	if err != nil {
		return
	}

	err = cacheIoc(injector)
	if err != nil {
		return
	}

	err = daoIoc(injector)
	if err != nil {
		return
	}

	err = repoIoc(injector)
	if err != nil {
		return
	}

	err = svcIoc(injector)
	if err != nil {
		return
	}

	err = handlerIoc(injector)
	if err != nil {
		return
	}

	return err
}
