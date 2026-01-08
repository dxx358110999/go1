package verify_code_svc

import "github.com/samber/do/v2"

func Provide(injector do.Injector) {
	/*
		验证码服务
	*/

	do.Provide(injector, NewVerifyCodeCache)
	do.Provide(injector, NewVerifyCodeSvc)
}
