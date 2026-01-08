package passwd_util

import "github.com/samber/do/v2"

func NewPasswordUtil(injector do.Injector) (passwordUtil PasswordUtilIface, err error) {
	passwordUtil = &PasswordUtilImpl{}
	return
}

func Provide(injector do.Injector) {
	do.Provide(injector, NewPasswordUtil)
}
