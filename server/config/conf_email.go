package config

type Email struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
	// 发件人邮箱
	UserAddr string `json:"user_addr" yaml:"user_addr"`
	Password string `json:"password" yaml:"password"`
	// 默认发件人名字
	DefaultFromEmail string `json:"default_from_email" yaml:"default_from_email"`
	UserSSL          bool   `json:"user_ssl" yaml:"user_ssl"`
	UserTls          bool   `json:"user_tls" yaml:"user_tls"`
}
