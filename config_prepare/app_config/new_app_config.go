package app_config

import (
	"dxxproject/pkg/nacos_ok"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/samber/do/v2"
	"gopkg.in/yaml.v2"
)

func contentToAppConfig(content string) (appConfig *AppConfig, err error) {
	appConfig = &AppConfig{}
	err = yaml.Unmarshal([]byte(content), appConfig)
	if err != nil {
		return
	}

	return
}

func NewAppConfig(injector do.Injector) (appConfig *AppConfig, err error) {
	instance := do.MustInvoke[*nacos_ok.NacosInstance](injector)
	content, err := instance.ConfigClient.GetConfig(vo.ConfigParam{
		DataId: instance.ClientConfig.DataId,
		Group:  instance.ClientConfig.GroupName,
	})
	if err != nil {
		return
	}

	appConfig, err = contentToAppConfig(content)
	if err != nil {
		return
	}

	return
}
func Provide(injector do.Injector) {
	do.Provide(injector, NewAppConfig)
}
