package ioc

import (
	"github.com/samber/do/v2"
)

func daoIoc(injector do.Injector) (err error) {
	do.Provide(injector, internal.NewUserDao)

	return err
}
