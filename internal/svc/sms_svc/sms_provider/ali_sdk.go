package sms_provider

import (
	"dxxproject/config_prepare/app_config"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v5/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
)

type AliYunSmsSdk struct {
}

// Description:
//
// 使用凭据初始化账号 Client
//
// @return Client
//
// @throws Exception

type SmsAliSdk struct {
	client *dysmsapi20170525.Client
}

func NewSmsAliSdk(cfg *app_config.AliSms) (err error, sdk *SmsAliSdk) {
	sdk = &SmsAliSdk{}
	err = sdk.createClient(cfg)
	if err != nil {
		return
	}
	return
}

func (rec *SmsAliSdk) createClient(cfg *app_config.AliSms) (err error) {

	credConfig := new(credentials.Config).
		SetType("access_key").
		SetAccessKeyId(cfg.AKId).
		SetAccessKeySecret(cfg.AKSecret)

	akCredential, err := credentials.NewCredential(credConfig)
	if err != nil {
		return
	}

	config := &openapi.Config{}
	config.Credential = akCredential
	config.Endpoint = tea.String(cfg.Host)

	client, err := dysmsapi20170525.NewClient(config)
	if err != nil {
		return err
	}

	rec.client = client

	return

}

func (rec *SmsAliSdk) Send(request *dysmsapi20170525.SendSmsRequest) (err error) {

	//这里估计可能抛出panic......瞎几把搞
	err = func() (_err error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_err = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, err = rec.client.SendSmsWithOptions(request, &util.RuntimeOptions{})
		if err != nil {
			return err
		}

		return nil
	}()

	//if err != nil {
	//	var sErr = &tea.SDKError{}
	//	if _t, ok := err.(*tea.SDKError); ok {
	//		sErr = _t
	//	} else {
	//		sErr.Message = tea.String(err.Error())
	//	}
	//	// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
	//	// 错误 message
	//	fmt.Println(tea.StringValue(sErr.Message))
	//	// 诊断地址
	//	var data interface{}
	//	d := json.NewDecoder(strings.NewReader(tea.StringValue(sErr.Data)))
	//	d.Decode(&data)
	//	if m, ok := data.(map[string]interface{}); ok {
	//		recommend, _ := m["Recommend"]
	//		fmt.Println(recommend)
	//	}
	//	_, err = util.AssertAsString(sErr.Message)
	//	if err != nil {
	//		return err
	//	}
	//}
	return err
}
