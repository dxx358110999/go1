package ioc

import (
	"dxxproject/internal/dao"
	"github.com/samber/do/v2"
)

func daoIoc(injector do.Injector) (err error) {
	do.Provide(injector, dao.NewUserDao)

	return err
}
