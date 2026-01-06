package verify_code_svc

import (
	"dxxproject/toolkit"
	"fmt"
	"time"
)

type VerifyCodeFat struct {
	code string

	codeKey            string
	codeMinuteLimitKey string
	codeDayLimitKey    string
	tryVerifyCountKey  string
	expire             time.Duration
}

func NewVerifyCodeFat(
	registerFeature string,
	bizCode string,
	code string,
) (vc *VerifyCodeFat) {
	registerFeature = toolkit.StringToBase64RemoveEqual(registerFeature) //注册特征转base64

	codeKey := fmt.Sprintf("verifyCode:%s:%s", bizCode, registerFeature)
	codeMinuteLimitKey := fmt.Sprintf("verifyCode:%s:%s:minuteLimit", bizCode, registerFeature)
	codeDayLimitKey := fmt.Sprintf("verifyCode:%s:%s:dayLimit", bizCode, registerFeature)
	tryVerifyCountKey := fmt.Sprintf("verifyCode:%s:%s:tryCount", bizCode, registerFeature)
	vc = &VerifyCodeFat{
		code:               code,
		codeKey:            codeKey,
		codeMinuteLimitKey: codeMinuteLimitKey,
		codeDayLimitKey:    codeDayLimitKey,
		tryVerifyCountKey:  tryVerifyCountKey,
		expire:             10 * time.Minute,
	}
	return
}
