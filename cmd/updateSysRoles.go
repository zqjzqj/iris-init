package main

import (
	"9xbet_risk/config"
	"9xbet_risk/logs"
	"9xbet_risk/services"
	"flag"
	"os"
)

var configPath = flag.String("config", "./config", "配置文件路径")

func init() {
	flag.Parse()
	err := config.LoadConfigJson(*configPath)
	if err != nil {
		logs.Fatal("配置文件载入错误", err)
	}
	_ = os.Setenv("TZ", "Asia/Shanghai")
}

func main() {
	services.NewRolesService().UpdateSysRole()
	logs.PrintlnSuccess("OK")
}
