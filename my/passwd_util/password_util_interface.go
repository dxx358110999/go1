package passwd_util

type PasswordUtilIF interface {
	Encrypt(pass string) (err error, enPass string)
	Compare(enPass, pass string) (err error)
}
