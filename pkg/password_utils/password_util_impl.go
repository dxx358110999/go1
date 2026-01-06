package password_utils

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordUtilImpl struct {
}

var _ PasswordUtilIF = (*PasswordUtilImpl)(nil)

func (rec *PasswordUtilImpl) Encrypt(pass string) (err error, enPass string) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return err, ""
	}
	return nil, string(bytes)
}

func (rec *PasswordUtilImpl) Compare(enPass, pass string) (err error) {
	err = bcrypt.CompareHashAndPassword(
		[]byte(enPass),
		[]byte(pass),
	)
	if err != nil {
		return err
	}
	return nil
}
