package main

import (
	"flag"
	"iris-init/cmd/servTemplate/tplStruct"
	"iris-init/logs"
	"strings"
)

var migrate = flag.String("migrate", "", "迁移models ','多个逗号分割")
var model = flag.String("model", "", "model名")
var _model = flag.String("_model", "", "model名,在创建services和repo的时候也生成一个临时用于复制ShowMap方法的model")
var createModel = flag.String("createModel", "", "创建model")
var tableName = flag.String("tableName", "", "创建model的表名, 与createModel关联使用 为空则使用model的蛇形作为表名")
var view = flag.String("view", "", "创建view 默认空 不创建")
var alias = flag.String("alias", "", "alias")
var appRoot = flag.String("appRoot", "", "appRoot项目路径，默认为当前目录") //rollback
var ctrDir = flag.String("ctrDir", "", "控制器生成的子目录，默认为空")        //rollback
var force = flag.Bool("force", false, "重名是否强制生成， 为true则会选择一个新的文件名生成，但不会覆盖已存在的文件")

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
	models := strings.Split(*model, ",")
	for _, modelItem := range models {
		servTpl := tplStruct.NewServTpl(modelItem, *alias, *ctrDir)
		if *appRoot != "" {
			servTpl.SetAppPath(*appRoot)
		}
		if *view != "" {
			servTpl.SetViewDir(*view)
		}
		servTpl.Force = *force
		_ = servTpl.GenerateFile(true)
		if modelItem != "" {
			err := servTpl.GenerateModel()
			if err != nil {
				logs.PrintErr(err)
			}
		}
		logs.PrintlnSuccess("GenerateFile OK...", modelItem)
	}

}
