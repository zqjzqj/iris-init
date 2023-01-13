package main

import (
	"flag"
	"iris-init/cmd/servTemplate/tplStruct"
	"iris-init/logs"
	"strings"
)

var migrate = flag.String("migrate", "", "迁移models ','多个逗号分割")
var model = flag.String("model", "", "model名")
var createModel = flag.String("createModel", "", "创建model")
var tableName = flag.String("tableName", "", "创建model的表名, 与createModel关联使用 为空则使用model的蛇形作为表名")
var view = flag.String("view", "", "创建view 默认空 不创建")
var alias = flag.String("alias", "", "alias")
var appRoot = flag.String("appRoot", "", "appRoot项目路径，默认为当前目录") //rollback
var ctrDir = flag.String("ctrDir", "", "控制器生成的子目录，默认为空")        //rollback

func init() {
	flag.Parse()
	logs.IsPrintLog = true
}

func main() {
	if *createModel != "" {
		modelTpl := tplStruct.NewModelTpl(*createModel, *tableName)
		err := modelTpl.GenerateFile()
		if err != nil {
			logs.Fatal(err)
		}
		logs.PrintlnSuccess("create model ok...")
		return
	}
	if *migrate != "" {
		models := strings.Split(*migrate, ",")
		migrateTpl := tplStruct.NewMigrateTpl(models)
		err := migrateTpl.GenerateFile()
		if err != nil {
			logs.Fatal(err)
		}
		logs.PrintlnSuccess("Migrate OK...")
		if *model == "" {
			return
		}
	}
	servTpl := tplStruct.NewServTpl(*model, *alias, *ctrDir)
	if *appRoot != "" {
		servTpl.SetAppPath(*appRoot)
	}
	if *view != "" {
		servTpl.SetViewDir(*view)
	}
	_ = servTpl.GenerateFile(true)
	logs.PrintlnSuccess("GenerateFile OK...")
}
