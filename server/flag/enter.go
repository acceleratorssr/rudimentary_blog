package flag

import sysflag "flag"

type Option struct {
	DB   bool
	User string // -u admin / -u user
}

func Parse() Option {
	db := sysflag.Bool("db", false, "show db")
	user := sysflag.String("u", "", "创建用户")
	// 将命令行中传递的参数解析并存储到相应的变量中
	sysflag.Parse()
	return Option{
		DB:   *db,
		User: *user,
	}
}

func IsWebStop(option Option) bool {
	if option.DB || option.User != "" {
		return true
	}
	return false
}

func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
		return
	}
	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
		return
	}
	sysflag.Usage()
}
