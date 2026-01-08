package start_config

//type envConfig struct {
//	StartConfigPath string `mapstructure:"START_CONFIG_PATH"`
//}

type Nacos struct {
	Address        string `mapstructure:"address"`
	Port           uint64 `mapstructure:"port"`
	DataId         string `mapstructure:"dataId"`
	NamespaceId    string `mapstructure:"namespaceId"`
	DefaultCluster string `mapstructure:"defaultCluster"`
	DefaultGroup   string `mapstructure:"defaultGroup"`
	Username       string `mapstructure:"username"`
	Password       string `mapstructure:"password"`
	ServiceName    string `mapstructure:"serviceName"`
}

type StartConfig struct {
	MachineID int64  `mapstructure:"machine_id"`
	Port      uint64 `mapstructure:"port"`

	*Nacos `mapstructure:"nacos"`
}
