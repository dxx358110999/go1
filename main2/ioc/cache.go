package ioc

import (
	"dxxproject/internal/cache"
	"github.com/samber/do/v2"
)

func cacheIoc(injector do.Injector) (err error) {
	do.Provide(injector, internal.NewUserCache)
	do.Provide(injector, cache.NewVerifyCodeCache)

	return err
}
