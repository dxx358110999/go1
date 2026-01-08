package handlers

import "github.com/samber/do/v2"

func Provide(injector do.Injector) {
	do.Provide(injector, NewVerifyCodeHdl)
	do.Provide(injector, NewUserHandlers)
}
