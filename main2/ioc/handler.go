package ioc

import (
	"dxxproject/internal/handler"
	"dxxproject/internal/user/user_handler"
	"github.com/samber/do/v2"
)

func handlerIoc(injector do.Injector) (err error) {
	do.Provide(injector, handler.NewVerifyCodeHandlers)
	do.Provide(injector, user_handler.NewUserHandlers)
	return err
}
