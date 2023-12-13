package config

import "strconv"

type Mysql struct {
	// Host of the mysql server
	Host string `yaml:"host"`
	// Port of the mysql server
	Port   int    `yaml:"port"`
	Config string `yaml:"config"`
	// 选择连接的数据库
	DB string `yaml:"db"`
	// 连接数据库的用户名
	Username string `yaml:"user"`
	// 连接数据库的密码
	Password string `yaml:"password"`
	// 选择日志等级
	LogLevel string `yaml:"log_level"`
	//// Log slow queries
	//LogSlowQueries int `yaml:"log_slow_queries"`
	//// Log slow queries time
	//LogSlowQueriesTime int `yaml:"log_slow_queries_time"`
	// 最大并发连接数
	MaxConnections int `yaml:"max_connections"`
	// 最大空闲连接数
	MaxIdleConnections int `yaml:"max-idle-connections"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config
}
