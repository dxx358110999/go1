package user_svc

import "github.com/samber/do/v2"

func Provide(injector do.Injector) {

	do.Provide(injector, NewUserCache)
	do.Provide(injector, NewUserDao)
	do.Provide(injector, NewUserRepo)
	do.Provide(injector, NewUserSvc)
}
