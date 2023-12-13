package config

type QiNiu struct {
	Enable    string `json:"enable" yaml:"enable"`
	AccessKey string `json:"access_key" yaml:"access_key"`
	SecretKey string `json:"secret_key" yaml:"secret_key"`
	// 存储桶的名字
	Bucket string `json:"bucket" yaml:"bucket"`
	// 访问图片的地址前缀
	CDN string `json:"cdn" yaml:"cdn"`
	// 存储的地区
	Zone string `json:"zone" yaml:"zone"`
	// 存储的大小限制，单位是MB
	Size float64 `json:"size" yaml:"size"`
}
