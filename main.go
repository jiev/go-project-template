package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"go-project-template/config"
	"os"
)

func main() {
	fmt.Printf("Module Start...")

	// 读取配置
	var configPath string
	flag.StringVar(&configPath,"config","conf/module.conf","module config")

	var config config.Config
	if _,err := toml.DecodeFile(configPath,&config); err != nil {
		fmt.Printf("read config :%v error,exit!",configPath)
		os.Exit(1)
		return
	}

	// 初始化



}

