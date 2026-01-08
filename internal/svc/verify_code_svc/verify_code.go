package verify_code_svc

import (
	"context"
	"dxxproject/agreed/biz"
	"dxxproject/internal/svc/email_svc"
	"dxxproject/internal/svc/rate_limit_svc"
	"dxxproject/internal/svc/sms_svc"
	"dxxproject/internal/svc/sms_svc/sms_provider"
	"dxxproject/my/my_err"
	"fmt"
	"github.com/pkg/errors"
	"github.com/samber/do/v2"
	"time"
)

type VerifyCodeSvc struct {
	limitSvc *rate_limit_svc.RateLimitSvc
	emailSvc *email_svc.EmailSvc
	smsSvc   sms_svc.SvcSmsIface
	vcCache  *VerifyCodeCache
}

func (r *VerifyCodeSvc) preCheckAndStore(ctx context.Context,
	regCode *RegisterCode,
) (err error) {
	/*
		每分钟发送频率
		每天发送频率
	*/

	err = r.limitSvc.GetAllow(ctx, regCode.MinuteLimitKey, 1, 1*time.Minute) //每分钟发送1条
	if err != nil {
		return
	}

	err = r.limitSvc.GetAllow(ctx, regCode.DayLimitKey, 5, 1*24*time.Hour) //每天发送5条
	if err != nil {
		return
	}

	err = r.vcCache.VerifyCodeSet(ctx, regCode.Key, regCode.Code, regCode.Expire)
	if err != nil {
		err = errors.Wrap(err, "cache写入验证码失败")
		return
	}

	//发送新的验证,清空尝试验证次数
	_ = r.limitSvc.LimitClear(ctx, regCode.TryVerifyCountKey)

	return
}

func (r *VerifyCodeSvc) SendCodeByEmail(ctx context.Context, biz biz.Biz, toEmail string) (err error) {
	code := generateCode()
	codeFat := NewRegisterCode(biz.Code, toEmail, code)
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
	codeFat := NewRegisterCode(phoneNumber, biz.Code, code)
	err = r.preCheckAndStore(ctx, codeFat)
	if err != nil {
		return err
	}

	// 2. 构建邮件内容

	phoneNumbers := []string{phoneNumber}
	params["Code"] = code

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
	//fmt.Printf("验证码 %s 已发送至 %s\n", Code, toEmail)
	return
}

func (r *VerifyCodeSvc) Verify(ctx context.Context,
	bizCode string,
	registerFeature string,
	userInputCode string,
) error {
	codeFat := NewRegisterCode(bizCode, registerFeature, userInputCode)

	//限制验证次数
	err := r.limitSvc.GetAllow(ctx, codeFat.TryVerifyCountKey, 3, 10*time.Minute) //有效期内,允许验证3次
	if err != nil {
		return err
	}

	signedCode, err := r.vcCache.VerifyCodeGet(ctx, codeFat.Key)
	if err != nil {
		return err
	}

	if signedCode != userInputCode {
		return my_err.ErrVerifyCodeNotMatch
	}

	return nil
}

func NewVerifyCodeSvc(injector do.Injector) (*VerifyCodeSvc, error) {
	rateLimiter := do.MustInvoke[*rate_limit_svc.RateLimitSvc](injector)
	vcCache := do.MustInvoke[*VerifyCodeCache](injector)
	emailSvc := do.MustInvoke[*email_svc.EmailSvc](injector)
	smsSvc := do.MustInvoke[sms_svc.SvcSmsIface](injector)

	vc := &VerifyCodeSvc{
		limitSvc: rateLimiter,
		emailSvc: emailSvc,
		smsSvc:   smsSvc,
		vcCache:  vcCache,
	}
	return vc, nil
}
