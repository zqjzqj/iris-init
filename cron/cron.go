package cron

import (
	cron2 "github.com/robfig/cron/v3"
	"iris-init/logs"
	"iris-init/services"
)

var cron *cron2.Cron

func before() {
	logs.PrintlnInfo("执行启动前任务...")
	_ = services.NewSettingsService().ReloadSettings()
	logs.PrintlnSuccess("执行启动前任务完成")
}

func InitCron() error {
	before()
	return nil
	cron = cron2.New(cron2.WithSeconds())
	cron.Start()
	return nil
}
