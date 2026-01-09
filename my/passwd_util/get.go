package passwd_util

func NewPasswordUtil() (passwordUtil *PasswordUtilImpl, err error) {
	passwordUtil = &PasswordUtilImpl{}
	return
}
