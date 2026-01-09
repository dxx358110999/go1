package verify_code

import (
	"dxxproject/internal/verify_code/cache"
	"dxxproject/internal/verify_code/hdl"
	"dxxproject/internal/verify_code/svc"
	"github.com/samber/do/v2"
)

func Provide(injector do.Injector) {
	do.Provide(injector, cache.NewVerifyCodeCache)
	do.Provide(injector, svc.NewVerifyCodeSvc)
	do.Provide(injector, hdl.NewVerifyCodeHdl)
}
