package ioc

import (
	"dxxproject/pkg/gin"
	"github.com/samber/do/v2"
)

/*
使用 samber do v2
默认情况下，Article 创建的是单例
*/

func Inject() (injector do.Injector, err error) {

	injector = do.New()

	injectPkg(injector) //注入基础Pkg

	injectModule(injector) //注入服务

	gin.Provide(injector) //注入gin

	return
}

func injectModule(injector do.Injector) {
	RateLimit(injector)
	Sms(injector)
	Email(injector)
	User(injector)
	Article(injector)
	VerifyCode(injector)

}
