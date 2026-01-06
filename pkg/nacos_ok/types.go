package nacos_ok

type NacosServerConfig struct {
	address string
	port    uint64
}

type NacosClientConfig struct {
	LocalIp     string
	AppPort     uint64
	ServiceName string
	Weight      float64
	Enable      bool
	Healthy     bool
	Ephemeral   bool
	Metadata    map[string]string
	ClusterName string
	GroupName   string

	DataId      string
	username    string
	password    string
	namespaceId string

	ServerAddress string
	ServerPort    uint64
}
