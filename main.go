package main

import (
	"flag"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	recover2 "github.com/kataras/iris/v12/middleware/recover"
	"github.com/rs/cors"
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
	"time"
)

var configPath = flag.String("config", "./config", "配置文件路径")
var migrateCmd = flag.String("migrate-back", "", "回滚迁移所需要的版本号") //rollback

func init() {
	flag.Parse()
	err := config.LoadConfigJson(*configPath)
	if err != nil {
		logs.Fatal("配置文件载入错误", err)
	}
	_ = os.Setenv("TZ", "Asia/Shanghai")
}

func migrateFunc() {
	if *migrateCmd != "" {
		migrates.Rollback(*migrateCmd)
	} else {
		migrates.Migrate()
		return
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
			logs.PrintlnWarning(err)
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
	//进程退出时
	end := pRuntime.HandleEndSignal(func() {
		logs.PrintlnInfo("Exiting...")
		global.HandleAppEndFunc(app)
		logs.PrintlnInfo("Exit OK.")
	})
	err = ListenWeb(app)
	if err != nil {
		logs.Fatal(err)
	}
	<-end
}

func ListenWeb(appWeb *iris.Application) error {
	cOption := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}
	if config.EnvIsDev() {
		cOption.Debug = true
	}
	c := cors.New(cOption)
	appWeb.WrapRouter(c.ServeHTTP)
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
		TimeFormat:        time.DateTime,
		RemoteAddrHeaders: []string{"X-Real-Ip", "X-Forwarded-For"},
	}))
	if err != nil {
		return err
	}
	return nil
}
