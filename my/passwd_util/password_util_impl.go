package passwd_util

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswdUtilImpl struct {
}

var _ PasswordUtilIface = (*PasswdUtilImpl)(nil)

func (rec *PasswdUtilImpl) Encrypt(pass string) (err error, enPass string) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return err, ""
	}
	return nil, string(bytes)
}

func (rec *PasswdUtilImpl) Compare(enPass, pass string) (err error) {
	err = bcrypt.CompareHashAndPassword(
		[]byte(enPass),
		[]byte(pass),
	)
	if err != nil {
		return err
	}
	return nil
}
