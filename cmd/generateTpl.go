package main

import (
	"flag"
	"iris-init/cmd/servTemplate"
	"iris-init/logs"
	"strings"
)

var migrate = flag.String("migrate", "", "迁移models ','多个逗号分割")
var model = flag.String("model", "", "model名")
var view = flag.String("view", "", "创建view 默认空 不创建")
var alias = flag.String("alias", "", "alias")
var appRoot = flag.String("appRoot", "", "appRoot项目路径，默认为当前目录") //rollback
var ctrDir = flag.String("ctrDir", "", "控制器生成的子目录，默认为空")        //rollback

func init() {
	flag.Parse()
	logs.IsPrintLog = true
}

func main() {
	if *migrate != "" {
		models := strings.Split(*migrate, ",")
		migrateTpl := servTemplate.NewMigrateTpl(models)
		err := migrateTpl.GenerateFile()
		if err != nil {
			logs.Fatal(err)
		}
		logs.PrintlnSuccess("Migrate OK...")
		if *model == "" {
			return
		}
	}
	servTpl := servTemplate.NewServTpl(*model, *alias, *ctrDir)
	if *appRoot != "" {
		servTpl.SetAppPath(*appRoot)
	}
	if *view != "" {
		servTpl.SetViewDir(*view)
	}
	_ = servTpl.GenerateFile(true)
	logs.PrintlnSuccess("GenerateFile OK...")
}
