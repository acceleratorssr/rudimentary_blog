package main

import (
	"server/core"
	"server/flag"
	"server/global"
	"server/routers"
)

func main() {
	// 读取配置文件
	core.UMYaml()
	// 初始化日志
	global.Log = core.InitLogger()
	// 初始化数据库
	global.DB = core.Gorm()

	// 命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	r := routers.InitRouters()
	serAddr := global.Config.System.Addr()
	global.Log.Info("server run addr:", serAddr)
	err := r.Run(serAddr)
	if err != nil {
		return
	}
}
