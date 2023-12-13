package flag

import sysflag "flag"

type Option struct {
	DB bool
}

func Parse() Option {
	db := sysflag.Bool("db", false, "show db")
	// 将命令行中传递的参数解析并存储到相应的变量中
	sysflag.Parse()
	return Option{
		DB: *db,
	}
}

func IsWebStop(option Option) bool {
	if option.DB {
		return true
	}
	return false
}

func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
	}
}
