package verify_code

import (
	"context"
	"dxxproject/agreed/biz"
	"dxxproject/internal/cache"
	"dxxproject/internal/svc/email"
	"dxxproject/internal/svc/rate_limit"
	"dxxproject/internal/svc/sms"
	"dxxproject/my/my_err"
	"dxxproject/pkg/sms_provider"
	"fmt"
	"github.com/pkg/errors"
	"github.com/samber/do/v2"
	"time"
)

type VerifyCodeSvc struct {
	limitSvc *rate_limit.RateLimitSvc
	emailSvc *email.EmailSvc
	smsSvc   sms.SvcSmsIface
	vcCache  *cache.VerifyCode
}

func (r *VerifyCodeSvc) preCheckAndStore(ctx context.Context,
	codeFat *VerifyCodeFat,
) (err error) {
	/*
		每分钟发送频率
		每天发送频率
	*/

	err = r.limitSvc.LimitGetAllow(ctx, codeFat.codeMinuteLimitKey, 1, 1*time.Minute) //每分钟发送1条
	if err != nil {
		return
	}

	err = r.limitSvc.LimitGetAllow(ctx, codeFat.codeDayLimitKey, 5, 1*24*time.Hour) //每天发送5条
	if err != nil {
		return
	}

	err = r.vcCache.VerifyCodeSet(ctx, codeFat.codeKey, codeFat.code, codeFat.expire)
	if err != nil {
		err = errors.Wrap(err, "cache写入验证码失败")
		return
	}

	//发送新的验证,清空尝试验证次数
	_ = r.limitSvc.LimitClear(ctx, codeFat.tryVerifyCountKey)

	return
}

func (r *VerifyCodeSvc) SendCodeByEmail(ctx context.Context, biz biz.Biz, toEmail string) (err error) {
	code := generateCode()
	codeFat := NewVerifyCodeFat(biz.Code, toEmail, code)
	err = r.preCheckAndStore(ctx, codeFat)
	if err != nil {
		return err
	}

	// 2. 构建邮件内容
	subject := "【验证码】您的验证请求"
	body := fmt.Sprintf("您的%s验证码是：%s，有效期10分钟，请妥善保管，切勿泄露给他人。", biz.Name, code)

	err = r.emailSvc.SendEmail(toEmail, subject, body)
	if err != nil {
		return err
	}
	fmt.Printf("验证码 %s 已发送至 %s\n", code, toEmail)
	return
}

func (r *VerifyCodeSvc) SendCodeBySms(ctx context.Context,
	biz biz.Biz,
	phoneNumber string,
	params map[string]string,
) (err error) {
	code := generateCode()
	codeFat := NewVerifyCodeFat(phoneNumber, biz.Code, code)
	err = r.preCheckAndStore(ctx, codeFat)
	if err != nil {
		return err
	}

	// 2. 构建邮件内容

	phoneNumbers := []string{phoneNumber}
	params["code"] = code

	smsInfo := sms_provider.SmsSendInfo{
		PhoneNumbers: phoneNumbers,
		SignedName:   "[星星公司]",
		TemplateId:   "tpl123",
		Params:       params,
	}

	err = r.smsSvc.SendSms(ctx, smsInfo)
	if err != nil {
		return err
	}
	//fmt.Printf("验证码 %s 已发送至 %s\n", code, toEmail)
	return
}

func (r *VerifyCodeSvc) Verify(ctx context.Context,
	bizCode string,
	registerFeature string,
	userInputCode string,
) (err error) {
	codeFat := NewVerifyCodeFat(bizCode, registerFeature, userInputCode)

	//限制验证次数
	err = r.limitSvc.LimitGetAllow(ctx, codeFat.tryVerifyCountKey, 3, 10*time.Minute) //有效期内,允许验证3次
	if err != nil {
		return
	}

	signedCode, err := r.vcCache.VerifyCodeGet(ctx, codeFat.codeKey)
	if err != nil {
		return err
	}

	if signedCode != userInputCode {
		return my_err.ErrVerifyCodeNotMatch
	}

	return
}

func NewVerifyCodeSvc(injector do.Injector) (*VerifyCodeSvc, error) {
	rateLimiter := do.MustInvoke[*rate_limit.RateLimitSvc](injector)
	vcCache := do.MustInvoke[*cache.VerifyCode](injector)
	emailSvc := do.MustInvoke[*email.EmailSvc](injector)
	smsSvc := do.MustInvoke[sms.SvcSmsIface](injector)

	vc := &VerifyCodeSvc{
		limitSvc: rateLimiter,
		emailSvc: emailSvc,
		smsSvc:   smsSvc,
		vcCache:  vcCache,
	}
	return vc, nil
}
