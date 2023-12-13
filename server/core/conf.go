package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/fs"
	"log"
	"os"
	"server/config"
	"server/global"
)

const ConfigFile = "conf.yaml"

func UMYaml() {
	// 把 c 作为参数传递给其他函数时，这些函数可以直接修改 c 指向的变量的值
	// var c *config.Config
	c := &config.Config{}
	// 在执行 Go 程序时的工作目录下查找名为 conf.yaml 的文件
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("read yaml error: %v\n", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config init unmarshal: %v\n", err)
	}
	log.Println("config init success")
	global.Config = c
}

func UpdateYaml() error {
	marshal, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}

	err = os.WriteFile(ConfigFile, marshal, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Info("config update success")
	return nil
}
