package main

import (
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	// 鉴权key
	Key string
	// 该patch服务的端口
	ServerPort int
	// Bot端服务端口
	BotServerPort int
)

func init() {
	cfg, err := ini.ShadowLoad("CoreConf.conf")
	if err != nil {
		fmt.Printf("读取配置文件(CoreConf.conf)失败, %v\n", err)
	}
	sec, _ := cfg.GetSection(ini.DEFAULT_SECTION)

	Key = sec.Key("PatchKey").MustString("12345678")
	ServerPort = sec.Key("PatchServerPort").MustInt(8899)
	// 处理bot端口 原来的格式: 0.0.0.0:8888
	temp := sec.Key("Port").MustString("0.0.0.0:8888")
	tempSplits := strings.Split(temp, ":")
	if len(tempSplits) == 2 {
		portStr := tempSplits[1]
		if port, err := strconv.Atoi(portStr); err == nil {
			BotServerPort = port
		}
	} else {
		BotServerPort = 8888
	}
	// prompt
	fmt.Println("======================================================")
	fmt.Println("当前配置(如有错误请修改配置更改):")
	fmt.Println("鉴权key(PatchKey):                      ", Key)
	fmt.Println("该Patch服务运行的端口(PatchServerPort): ", ServerPort)
	fmt.Println("原Bot端运行的端口(Port):                ", BotServerPort)
	fmt.Println("======================================================")
}
