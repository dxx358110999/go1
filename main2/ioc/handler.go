package ioc

import (
	"dxxproject/internal/handler"
	"github.com/samber/do/v2"
)

func handlerIoc(injector do.Injector) (err error) {
	do.Provide(injector, handler.NewVerifyCodeHandlers)
	do.Provide(injector, handler.NewUserHandlers)
	return err
}
