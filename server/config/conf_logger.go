package config

type Logger struct {
	// 设置日志级别
	Level string `yaml:"level"`
	// 设置日志前缀
	Prefix string `yaml:"prefix"`
	// 设置日志记录方向
	Director string `yaml:"director"`
	// 是否显示行号
	ShowLine bool `yaml:"show-line"`
	// 是否在控制台输出日志
	LogInConsole bool `yaml:"log-in-console"`
}
