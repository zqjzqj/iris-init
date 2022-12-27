package main

import (
	"flag"
	"iris-init/cmd/servTemplate"
	"iris-init/logs"
)

var model = flag.String("model", "", "model名")
var alias = flag.String("alias", "", "alias")
var appRoot = flag.String("appRoot", "", "appRoot项目路径，默认为当前目录") //rollback

func init() {
	flag.Parse()
	logs.IsPrintLog = true
}

func main() {
	servTpl := servTemplate.NewServTpl(*model, *alias)
	if *appRoot != "" {
		servTpl.SetAppPath(*appRoot)
	}
	err := servTpl.GenerateFile()
	if err != nil {
		logs.Fatal(err)
	}
	logs.PrintlnSuccess("OK...")
}
