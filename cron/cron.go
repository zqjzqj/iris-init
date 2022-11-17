package cron

import (
	cron2 "github.com/robfig/cron/v3"
)

var cron *cron2.Cron

func InitCron() error {
	return nil
	cron = cron2.New(cron2.WithSeconds())
	cron.Start()
	return nil
}
