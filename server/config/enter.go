package config

// Config 注意：对应的yaml标签必须一样
type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `yaml:"site_info"`
	Wechat   Wechat   `yaml:"wechat"`
	QiNiu    QiNiu    `yaml:"qi_niu"`
	Jwt      Jwt      `yaml:"jwt"`
	Email163 Email    `yaml:"email_163"`
	//EmailQQ  Email    `yaml:"email_qq"`
	Upload        Upload        `yaml:"upload"`
	Redis         Redis         `yaml:"redis"`
	Elasticsearch Elasticsearch `yaml:"elasticsearch"`
}
