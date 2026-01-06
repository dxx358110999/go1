package ioc

import (
	"github.com/samber/do/v2"
)

func repoIoc(injector do.Injector) (err error) {
	do.Provide(injector, internal.NewUserRepo)

	return err
}
