package config

type Redis struct {
	Addr     string `json:"addr" yaml:"addr"`
	Password string `json:"password" yaml:"password"`
	PoolSize int    `json:"pool_size" yaml:"pool_size"`
}
