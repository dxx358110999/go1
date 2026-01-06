package nacos_ok

import (
	"dxxproject/config_prepare/start_config"
	"dxxproject/main2/basic_info"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/samber/do/v2"
	"log"
)

type NacosInstance struct {
	ClientConfig *NacosClientConfig
	ServerConfig *NacosServerConfig

	//内部创建
	NamingClient naming_client.INamingClient
	ConfigClient config_client.IConfigClient
}

func NewNacosInstance(injector do.Injector) (*NacosInstance, error) {
	/*
		需要考虑读取环境
	*/

	//取出依赖
	startConfig := do.MustInvoke[*start_config.StartConfig](injector)
	bi := do.MustInvoke[*basic_info.BasicInfo](injector)

	serverConfig, err := NewNacosServerConfig(startConfig)
	if err != nil {
		return nil, err
	}

	clientConfig, err := NewNacosClientConfig(startConfig, bi.LocalIp)
	if err != nil {
		return nil, err
	}

	instance := &NacosInstance{
		ClientConfig: clientConfig,
		ServerConfig: serverConfig,
		NamingClient: nil,
		ConfigClient: nil,
	}

	err = instance.build()
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (r *NacosInstance) build() (err error) {

	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(r.ClientConfig.namespaceId), //When namespace is public, fill in the blank string here.
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		//constant.WithLogDir("/tmp/nacos/log"),
		//constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
		constant.WithUsername(r.ClientConfig.username),
		constant.WithPassword(r.ClientConfig.password),
	)

	//Another way of create serverConfigs
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(
			r.ServerConfig.address,
			r.ServerConfig.port,
			constant.WithScheme("http"),
			constant.WithContextPath("/nacos"),
		),
	}

	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	r.NamingClient = namingClient

	if err != nil {
		return
	}

	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	r.ConfigClient = configClient
	return
}

func (r *NacosInstance) Register() (err error) {

	param := vo.RegisterInstanceParam{
		Ip:          r.ClientConfig.LocalIp,
		Port:        r.ClientConfig.AppPort,
		ServiceName: r.ClientConfig.ServiceName,
		Weight:      r.ClientConfig.Weight,
		Enable:      r.ClientConfig.Enable,
		Healthy:     r.ClientConfig.Healthy,
		Ephemeral:   r.ClientConfig.Ephemeral,
		Metadata:    r.ClientConfig.Metadata,
		ClusterName: r.ClientConfig.ClusterName, // 默认值DEFAULT
		GroupName:   r.ClientConfig.GroupName,   // 默认值DEFAULT_GROUP
	}

	success, err := r.NamingClient.RegisterInstance(param)
	if err != nil {
		log.Fatalf("注册服务失败: %s", err.Error())
	} else {
		log.Printf("nacos注册成功!%v", success)
	}
	return
}

func (r *NacosInstance) Deregister() (err error) {

	param := vo.DeregisterInstanceParam{
		Ip:          r.ClientConfig.LocalIp,
		Port:        r.ClientConfig.AppPort,
		ServiceName: r.ClientConfig.ServiceName,
		Ephemeral:   r.ClientConfig.Ephemeral,
		Cluster:     r.ClientConfig.ClusterName, // default value is DEFAULT
		GroupName:   r.ClientConfig.GroupName,   // default value is DEFAULT_GROUP
	}

	success, err := r.NamingClient.DeregisterInstance(param)
	if err != nil {
		log.Fatalf("nacos反注册失败: %s", err.Error())
	} else {
		log.Printf("nacos反注册成功! %b", success)
	}
	return
}

//func (r *NacosInstance) ListenConfig() (err error) {
//	//Listen cConfig change,key=dataId+group+namespaceId.
//	err = r.ConfigClient.ListenConfig(vo.ConfigParam{
//		DataId: r.cConfig.DataId,
//		Group:  r.cConfig.GroupName,
//		OnChange: func(namespace, group, dataId, data string) {
//			fmt.Println("cConfig changed group:" + group + ", dataId:" + dataId + ", content:" + data)
//			err = r.contentToAppConfig(data)
//			if err != nil {
//				fmt.Println("解析错误:", r.AppConfig)
//				return
//			}
//		},
//	})
//	if err != nil {
//		return
//	}
//
//	return
//}
