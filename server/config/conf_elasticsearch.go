package config

type Elasticsearch struct {
	Addr                  string
	Username              string
	Password              string
	MaxRetries            int
	Sniff                 bool // 指示是否启用嗅探功能，用于发现集群中其他节点的功能
	SniffOnStart          bool // 指示是否在启动时执行一次嗅探，用于发现集群中其他节点的功能
	SniffOnConnectionFail bool // 指示是否在连接失败时执行一次嗅探，用于发现集群中其他节点的功能
	SniffClose            bool // 指示是否在关闭连接时执行一次嗅探，用于发现集群中其他节点的功能
	MaxIdleConns          int  // 连接池中的最大空闲连接数
	MaxActiveConns        int  // 连接池中的最大活动连接数
	MaxStaleConns         int  // 连接池中的最大过期连接数
}
