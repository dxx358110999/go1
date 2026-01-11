package nacos_ok

import (
	"dxxproject/config_prepare/start_config"
)

func NewNacosClientConfig(startConfig *start_config.Config, localIp string) (config *NacosClientConfig, err error) {

	config = &NacosClientConfig{
		LocalIp:     localIp,
		AppPort:     startConfig.Port,
		ServiceName: startConfig.Nacos.ServiceName,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
		ClusterName: startConfig.Nacos.DefaultCluster, // default value is DEFAULT
		GroupName:   startConfig.Nacos.DefaultGroup,   // default value is DEFAULT_GROUP
		DataId:      startConfig.Nacos.DataId,
		username:    startConfig.Nacos.Username,
		password:    startConfig.Nacos.Password,
		namespaceId: startConfig.Nacos.NamespaceId,
	}

	//环境变量中的配置,覆盖配置文件中的配置
	envCfg, err := readEnvConfig()
	if err != nil {
		return nil, err
	}

	if envCfg.NacosUsername != "" && envCfg.NacosPassword != "" {
		config.username = envCfg.NacosUsername
		config.password = envCfg.NacosPassword
	}

	if envCfg.NacosNamespaceId != "" {
		config.namespaceId = envCfg.NacosNamespaceId
	}

	if envCfg.NacosCluster != "" {
		config.ClusterName = envCfg.NacosCluster
	}

	if envCfg.NacosGroup != "" {
		config.GroupName = envCfg.NacosGroup
	}

	return
}
