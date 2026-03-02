package main

import (
	"flag"
	"iris-init/config"
	"iris-init/logs"
	"iris-init/services"
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
