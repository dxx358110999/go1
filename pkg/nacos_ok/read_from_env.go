package nacos_ok

import "github.com/spf13/viper"

type envConfig struct {
	NacosAddress     string `mapstructure:"NACOS_ADDRESS"`
	NacosPort        uint64 `mapstructure:"NACOS_PORT"`
	NacosUsername    string `mapstructure:"NACOS_USERNAME"`
	NacosPassword    string `mapstructure:"NACOS_PASSWORD"`
	NacosNamespaceId string `mapstructure:"NACOS_NAMESPACE_ID"`
	NacosCluster     string `mapstructure:"NACOS_CLUSTER"`
	NacosGroup       string `mapstructure:"NACOS_GROUP"`
}

func readEnvConfig() (envCfg *envConfig, err error) {
	//从环境中读取配置
	//viper.SetEnvPrefix("app") // 设置前缀
	viper.AllowEmptyEnv(true) //读取空
	viper.AutomaticEnv()      // 自动读取环境变量
	envCfg = &envConfig{}
	err = viper.Unmarshal(envCfg)
	if err != nil {
		return
	}

	return
}
