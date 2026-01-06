package password_utils

import "github.com/samber/do/v2"

func NewPasswordUtil(injector do.Injector) (passwordUtil PasswordUtilIF, err error) {
	passwordUtil = &PasswordUtilImpl{}
	return
}
