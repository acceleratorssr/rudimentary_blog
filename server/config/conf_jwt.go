package config

type Jwt struct {
	// 密钥
	Secret string `json:"secret" yaml:"secret"`
	// 到期时间
	Expires int `json:"expires" yaml:"expires"`
	// 颁发人
	Issuer string `json:"issuer" yaml:"issuer"`
}
