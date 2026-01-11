package pwd_util

type PwdIface interface {
	Encrypt(pass string) (err error, enPass string)
	Compare(enPass, pass string) (err error)
}
