package start_config

import (
	"fmt"
	"github.com/samber/do/v2"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

/*
方式1：直接指定配置文件路径（相对路径或者绝对路径）
相对路径：相对执行的可执行文件的相对路径
viper.SetConfigFile("./conf/config.yaml")
绝对路径：系统中实际的文件路径
viper.SetConfigFile("/Users/liwenzhou/Desktop/app/conf/config.yaml")

方式2：指定配置文件名和配置文件的位置，viper自行查找可用的配置文件
配置文件名不需要带后缀
配置文件位置可配置多个
viper.SetConfigName("config") // 指定配置文件名（不带后缀）
viper.AddConfigPath(".") // 指定查找配置文件的路径（这里使用相对路径）
viper.AddConfigPath("./conf")      // 指定查找配置文件的路径（这里使用相对路径）

基本上是配合远程配置中心使用的，告诉viper当前的数据使用什么格式去解析
viper.SetConfigType("json")
*/

/*
START_CONFIG_FILE
NacosAddress
NacosPort
*/

func readEnvConfig() (*envConfig, error) {
	//从环境中读取配置
	//viper.SetEnvPrefix("app") // 设置前缀
	viper.AllowEmptyEnv(true) //读取空
	viper.AutomaticEnv()      // 自动读取环境变量
	envCfg := &envConfig{}
	err := viper.Unmarshal(envCfg)
	if err != nil {
		return nil, err
	}

	return envCfg, nil
}

func newStartConfig() (startCfg *StartConfig, err error) {
	//默认配置
	startCfg = &StartConfig{
		MachineID: 1,
		Port:      8080,
		Nacos: &Nacos{
			Address:        "docker.dxx.com",
			Port:           8848,
			DataId:         "dxxlianxi",
			NamespaceId:    "844ffd5b-2b7c-45d2-81c3-3d89cae92199",
			DefaultCluster: "DEFAULT",
			DefaultGroup:   "DEFAULT_GROUP",
			Username:       "nacos",
			Password:       "nacos",
			ServiceName:    "dxxlianxi_micro",
		},
	} //返回

	//指定启动配置文件
	var filePath *string //启动配置文件路径
	filePath = pflag.String("config", "", "配置文件路径")
	pflag.Parse()

	if filePath != nil && *filePath != "" {
		fmt.Println("-start-", "配置文件路径", *filePath)
		viper.SetConfigFile(*filePath) //配置文件路径
		err = viper.ReadInConfig()     // 读取配置信息
		if err != nil {
			fmt.Printf("读取配置信息失败, err:%v\n", err)
			return
		}

		err = viper.Unmarshal(startCfg) //用配置文件,覆盖默认配置
		if err != nil {
			fmt.Printf("反序列化配置失败, err:%v\n", err)
			return
		}

		/*		viper.WatchConfig()
				viper.OnConfigChange(func(in fsnotify.Event) {
					fmt.Println("配置文件修改了...")
					err := viper.Unmarshal(startCfg)
					if err != nil {
						fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
					}
				})*/
	} else {
		fmt.Println("-start-", "没有指定启动配置文件")
	}

	return
}

func NewStartConfig(injector do.Injector) (startCfg *StartConfig, err error) {
	config, err := newStartConfig()
	if err != nil {
		panic(err)
	}

	startCfg = config
	return
}
