package main

import (
	"flag"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	recover2 "github.com/kataras/iris/v12/middleware/recover"
	"github.com/zqjzqj/pRuntime"
	"iris-init/appWeb/routes"
	"iris-init/config"
	"iris-init/cron"
	"iris-init/global"
	"iris-init/logs"
	"iris-init/migrates"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
)

var configPath = flag.String("config", "./config", "配置文件路径")
var migrateCmd = flag.String("migrate", "", "迁移参数 run执行迁移 rollback回滚迁移") //rollback
var mRollbackId = flag.String("mRollbackId", "", "回滚迁移所需要的版本号")

func init() {
	flag.Parse()
	err := config.LoadConfigJson(*configPath)
	if err != nil {
		logs.Fatal("配置文件载入错误", err)
	}
	_ = os.Setenv("TZ", "Asia/Shanghai")
}

func migrateFunc() {
	if *migrateCmd == "" {
		return
	}
	if *migrateCmd == "run" {
		migrates.Migrate()
	} else if *migrateCmd == "rollback" {
		if *mRollbackId == "" {
			logs.Fatal("无效的回退版本号【请填写参数 mRollbackId】")
		}
		migrates.Rollback(*mRollbackId)
	}
	os.Exit(0)
}

func main() {
	migrateFunc()

	//后台模式运行
	if config.RunBackground() {
		if runtime.GOOS != "windows" {
			//设置一下pid文件
			pRuntime.SetPidFile("./iris-init.pid")
			pRuntime.RunBackground()
		} else {
			logs.PrintlnWarning("windows不支持后台运行模式")
		}
	}

	//子进程模式运行
	if config.RunDaemon() {
		err := pRuntime.RunDaemon(true)
		if err != nil {
			logs.PrintErr(err)
		}
	}
	//性能剖析
	if config.PprofPort() > 0 {
		go func() {
			log.Println(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", config.PprofPort()), nil))
		}()
	}

	//初始化计划任务
	err := cron.InitCron()
	if err != nil {
		logs.Fatal(err)
	}
	app := iris.New()

	err = ListenWeb(app)
	if err != nil {
		logs.Fatal(err)
	}
}

func ListenWeb(appWeb *iris.Application) error {
	appWeb.Use(recover2.New(), logger.New())
	//注册路由
	routes.RegisterRoutes(appWeb)
	port := config.GetWebCfg().GetPort()
	logs.PrintlnInfo("Http API List:")
	for _, r := range appWeb.GetRoutes() {
		if r.Method != "OPTIONS" {
			logs.PrintlnInfo(
				fmt.Sprintf("[%s] http://localhost:%d%s   [%s]", r.Method, port, r.Path, r.Name),
			)
		}
	}
	//监听http
	err := appWeb.Run(iris.Addr(fmt.Sprintf(":%d", port)), iris.WithConfiguration(iris.Configuration{
		TimeFormat:        global.DateTimeFormatStr,
		RemoteAddrHeaders: []string{"X-Real-Ip", "X-Forwarded-For"},
	}))
	if err != nil {
		return err
	}
	return nil
}
