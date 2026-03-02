package tplStruct

import (
	"fmt"
	"iris-init/global"
	"iris-init/logs"
	"iris-init/sErr"
	"os"
	"reflect"
	"text/template"
)

type ModelTpl struct {
	Model            string
	TableName        string
	alias            string
	AppRoot          string
	modelTplPath     string
	modelPath        string
	str2modelTplPath string
	str2modelPath    string
}

func NewModelTpl(_model, tableName string) ModelTpl {
	if _model == "" {
		panic("model is not empty")
	}
	if tableName == "" {
		tableName = global.SnakeString(_model)
	}
	mt := ModelTpl{Model: _model, alias: global.StringFirstLower(_model), TableName: tableName}
	pwd, _ := os.Getwd()
	mt.SetAppPath(pwd)
	return mt
}

func (modelTpl *ModelTpl) SetAppPath(AppRoot string) {
	modelTpl.AppRoot = AppRoot
	modelTpl.modelTplPath = fmt.Sprintf("%s/cmd/servTemplate/model.tpl", modelTpl.AppRoot)
	modelTpl.modelPath = fmt.Sprintf("%s/model/%s.go", modelTpl.AppRoot, modelTpl.alias)

	modelTpl.str2modelTplPath = fmt.Sprintf("%s/cmd/servTemplate/str2modelMapTpl", modelTpl.AppRoot)
	modelTpl.str2modelPath = fmt.Sprintf("%s/cmd/servTemplate/tplStruct/str2modelMap.go", modelTpl.AppRoot)
}

func (modelTpl ModelTpl) GenerateFile() error {
	err := modelTpl.generateFile(modelTpl.modelTplPath, modelTpl.modelPath, nil)
	if err != nil {
		return fmt.Errorf("model err %v", err)
	}
	_str2ModelMap := make(map[string]string)
	for k, v := range Str2ModelMap {
		_str2ModelMap[k] = fmt.Sprintf("model.%s{}", reflect.TypeOf(v).Name())
	}
	_str2ModelMap[modelTpl.Model] = fmt.Sprintf("model.%s{}", modelTpl.Model)
	_tmpStr2ModelPath := fmt.Sprintf("%s.tmp", modelTpl.str2modelPath)
	err = os.Rename(modelTpl.str2modelPath, _tmpStr2ModelPath)
	if err != nil {
		logs.PrintErr("create str2model fail")
		return err
	}
	err = modelTpl.generateFile(modelTpl.str2modelTplPath, modelTpl.str2modelPath, map[string]any{
		"Str2ModelMap": _str2ModelMap,
	})
	if err == nil {
		_ = os.Remove(_tmpStr2ModelPath)
		return nil
	}
	return os.Rename(_tmpStr2ModelPath, modelTpl.str2modelPath)
}

func (modelTpl ModelTpl) generateFile(tplPath, filePath string, data map[string]any) error {
	_riTplBytes, err := os.ReadFile(tplPath)
	if err != nil {
		return err
	}
	tpl, err := template.New(filePath).
		Parse(string(_riTplBytes))
	if err != nil {
		return err
	}
	if global.FileExists(filePath) {
		return sErr.New("tpl file exists " + filePath)
	}
	f, err := os.Create(filePath)
	defer func() { _ = f.Close() }()
	if err != nil {
		return err
	}
	if data == nil {
		data = make(map[string]interface{})
	}
	data["Model"] = modelTpl.Model
	data["Alias"] = modelTpl.alias
	data["TableName"] = modelTpl.TableName
	return tpl.Execute(f, data)
}
