package start_config

import (
	"fmt"
	"github.com/samber/do/v2"
	"github.com/spf13/viper"
	"testing"
)

type TestConfig struct {
	StartConfigPath string `mapstructure:"START_CONFIG_PATH"`
	Port            uint16 `mapstructure:"PORT"`
}

func TestReadEnvEmpty(t *testing.T) {
	//从环境中读取配置
	viper.SetEnvPrefix("app") // 设置前缀
	viper.AutomaticEnv()      // 自动读取环境变量
	var envConfig TestConfig
	err := viper.Unmarshal(&envConfig)
	if err != nil {
		panic(err)
	}

	fmt.Println("path:", envConfig)
	fmt.Println("path:", envConfig.StartConfigPath == "")
	fmt.Println("path:", envConfig.Port)
}

func TestStart(t *testing.T) {
	i := do.New()
	config, err := NewStartConfig(i)
	if err != nil {
		return
	}
	fmt.Println("config:", config)
}
