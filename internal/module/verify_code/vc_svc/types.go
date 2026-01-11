package vc_svc

import (
	"dxxproject/toolkit"
	"fmt"
	"time"
)

type RegisterCode struct {
	Code              string // 实际 code
	Key               string //存放key
	MinuteLimitKey    string
	DayLimitKey       string
	TryVerifyCountKey string
	Expire            time.Duration
}

func NewRegisterCode(
	registerFeature string,
	bizCode string,
	code string,
) (vc *RegisterCode) {
	registerFeature = toolkit.StringToBase64RemoveEqual(registerFeature) //注册特征转base64

	codeKey := fmt.Sprintf("verifyCode:%s:%s", bizCode, registerFeature)
	codeMinuteLimitKey := fmt.Sprintf("verifyCode:%s:%s:minuteLimit", bizCode, registerFeature)
	codeDayLimitKey := fmt.Sprintf("verifyCode:%s:%s:dayLimit", bizCode, registerFeature)
	tryVerifyCountKey := fmt.Sprintf("verifyCode:%s:%s:tryCount", bizCode, registerFeature)
	vc = &RegisterCode{
		Code:              code,
		Key:               codeKey,
		MinuteLimitKey:    codeMinuteLimitKey,
		DayLimitKey:       codeDayLimitKey,
		TryVerifyCountKey: tryVerifyCountKey,
		Expire:            10 * time.Minute,
	}
	return
}
