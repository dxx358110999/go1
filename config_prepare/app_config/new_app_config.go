package app_config

import (
	"dxxproject/pkg/nacos_ok"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"gopkg.in/yaml.v2"
)

func contentToAppConfig(content string) (appConfig *Config, err error) {
	appConfig = &Config{}
	err = yaml.Unmarshal([]byte(content), appConfig)
	if err != nil {
		return
	}

	return
}

func NewAppConfig(instance *nacos_ok.NacosInstance) (appConfig *Config, err error) {
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
