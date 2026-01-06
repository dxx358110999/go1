package sms_provider

import (
	"context"
	"dxxproject/config_prepare/app_config"
	"encoding/json"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v5/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/samber/do/v2"
)

type AliSmsImpl struct {
	sdk *SmsAliSdk
}

func (rec *AliSmsImpl) SendSms(ctx context.Context, info SmsSendInfo) (err error) {

	//拼接多个手机号
	phoneNumberString := ""
	for _, number := range info.PhoneNumbers {
		phoneNumberString += number
		phoneNumberString += ","
	}
	phoneNumberString = phoneNumberString[:len(phoneNumberString)-1] //去掉末尾逗号

	//参数转json
	paramsJson, err := json.Marshal(info.Params)
	if err != nil {
		return err
	}
	paramsJsonStr := string(paramsJson)

	// 创建请求对象并设置入参
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		// 需替换成为您的短信模板code
		TemplateCode: tea.String(info.TemplateId),
		// 需替换成为您的短信模板变量对应的实际值，示例值：{"code":"1234"}
		TemplateParam: tea.String(paramsJsonStr),
		// 需替换成为您的接收手机号码
		PhoneNumbers: tea.String(phoneNumberString),
		// 需替换成为您的短信签名
		SignName: tea.String(info.SignedName),
	}
	err = rec.sdk.Send(sendSmsRequest)
	if err != nil {
		return err
	}
	return
}

var _ SmsProviderIF = new(AliSmsImpl)

func NewAliSmsImpl(injector do.Injector) (impl *AliSmsImpl, err error) {
	cfg := do.MustInvoke[*app_config.AppConfig](injector).AliSms

	err, sdk := NewSmsAliSdk(cfg)
	impl = &AliSmsImpl{
		sdk: sdk,
	}

	return
}
