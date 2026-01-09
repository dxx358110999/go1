package passwd_util

func NewPasswordUtil() (passwordUtil *PasswdUtilImpl, err error) {
	passwordUtil = &PasswdUtilImpl{}
	return
}
