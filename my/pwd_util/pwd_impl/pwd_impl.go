package pwd_impl

import (
	"dxxproject/my/pwd_util"
	"golang.org/x/crypto/bcrypt"
)

type PwdImpl struct {
}

var _ pwd_util.PwdIface = (*PwdImpl)(nil)

func (rec *PwdImpl) Encrypt(pass string) (err error, enPass string) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return err, ""
	}
	return nil, string(bytes)
}

func (rec *PwdImpl) Compare(encryptedPwd, pwd string) (err error) {
	err = bcrypt.CompareHashAndPassword(
		[]byte(encryptedPwd),
		[]byte(pwd),
	)
	if err != nil {
		return err
	}
	return nil
}
func NewPasswordUtil() (passwordUtil *PwdImpl, err error) {
	passwordUtil = &PwdImpl{}
	return
}
