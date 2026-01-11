package nacos_ok

import (
	"dxxproject/config_prepare/start_config"
)

func NewNacosServerConfig(startConfig *start_config.Config) (config *NacosServerConfig, err error) {
	var address string
	var port uint64
	address = startConfig.Nacos.Address
	port = startConfig.Nacos.Port

	config = &NacosServerConfig{
		address: address,
		port:    port,
	}

	envCfg, err := readEnvConfig()
	if err != nil {
		return
	}

	if envCfg.NacosAddress != "" {
		config.address = envCfg.NacosAddress
	}

	if envCfg.NacosPort != 0 {
		config.port = envCfg.NacosPort
	}

	return
}
