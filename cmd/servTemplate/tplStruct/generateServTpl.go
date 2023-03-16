package tplStruct

import (
	"fmt"
	"iris-init/global"
	"iris-init/logs"
	"iris-init/sErr"
	"os"
	"os/exec"
	"runtime"
	"text/template"
)

type ServTpl struct {
	Model                string
	Alias                string
	AppRoot              string
	servTplPath          string
	repoTplPath          string
	modelTplPath         string
	modelPath            string
	repoInterfaceTplPath string
	controllerTplPath    string
	repoPath             string
	repoInterfacePath    string
	servPath             string
	controllerPath       string
	controllerDir        string
	ViewDir              string
	viewListTplPath      string
	viewItemTplPath      string
	ModelField           []Field
	UniqueField          map[string][]Field
}

func NewServTpl(_model, alias, ctrDir string) ServTpl {
	if _model == "" {
		panic("model is not empty")
	}
	st := ServTpl{Model: _model, Alias: alias, controllerDir: ctrDir}
	pwd, _ := os.Getwd()
	st.SetAppPath(pwd)
	_modelStruct, ok := Str2ModelMap[st.Model]
	if !ok {
		logs.PrintErr(fmt.Sprintf("请先在cmd/servTemplate/tplStruct/str2modelMap.go中添加model的映射关系%s=>model.%s{}, 或使用go run ./cmd/generateTpl.go -createModel=%s命令创建", _model, _model, _model))
		panic("model参数错误")
	}
	st.ModelField = RefStructField(_modelStruct)
	st.UniqueField = GetUniqueFields(st.ModelField)
	return st
}

func (servTpl *ServTpl) SetViewDir(viewDir string) {
	if viewDir != "admin" {
		panic("无效的view参数，暂只支持admin")
	}
	servTpl.ViewDir = viewDir
}

func (servTpl *ServTpl) SetAppPath(AppRoot string) {
	servTpl.AppRoot = AppRoot
	servTpl.servTplPath = fmt.Sprintf("%s/cmd/servTemplate/services.tpl", servTpl.AppRoot)
	servTpl.repoTplPath = fmt.Sprintf("%s/cmd/servTemplate/repo.tpl", servTpl.AppRoot)
	servTpl.repoInterfaceTplPath = fmt.Sprintf("%s/cmd/servTemplate/repoInterface.tpl", servTpl.AppRoot)
	servTpl.controllerTplPath = fmt.Sprintf("%s/cmd/servTemplate/controller.tpl", servTpl.AppRoot)
	servTpl.viewListTplPath = fmt.Sprintf("%s/cmd/servTemplate/viewList.tpl", servTpl.AppRoot)
	servTpl.viewItemTplPath = fmt.Sprintf("%s/cmd/servTemplate/viewItem.tpl", servTpl.AppRoot)
	servTpl.modelTplPath = fmt.Sprintf("%s/cmd/servTemplate/_model.tpl", servTpl.AppRoot)
	servTpl.RefreshModel()
}

func (servTpl *ServTpl) SetModel(model, alias string) {
	servTpl.Model = model
	servTpl.Alias = alias
	servTpl.RefreshModel()
}

func (servTpl *ServTpl) RefreshModel() {
	_m := global.StringFirstLower(servTpl.Model)
	servTpl.repoPath = fmt.Sprintf("%s/repositories/%s.go", servTpl.AppRoot, _m+"Repo")
	servTpl.repoInterfacePath = fmt.Sprintf("%s/repositories/repoInterface/%s.go", servTpl.AppRoot, _m+"Repo")
	servTpl.servPath = fmt.Sprintf("%s/services/%s.go", servTpl.AppRoot, _m+"Service")
	servTpl.modelPath = fmt.Sprintf("%s/model/_%s.go", servTpl.AppRoot, _m)
	if servTpl.controllerDir == "" {
		servTpl.controllerPath = fmt.Sprintf("%s/appWeb/controller/%s.go", servTpl.AppRoot, _m+"Controller")
	} else {
		servTpl.controllerPath = fmt.Sprintf("%s/appWeb/controller/%s/%s.go",
			servTpl.AppRoot,
			servTpl.controllerDir,
			_m+"Controller",
		)
	}
	if servTpl.Alias == "" {
		servTpl.Alias = _m
	}
}

func (servTpl ServTpl) GenerateFile(ignoreErr bool) error {
	defer func() {
		if runtime.GOOS == "windows" {
			_ = exec.Command("cmd", "/C", fmt.Sprintf("gofmt -l -w ./")).Run()
		} else {
			_ = exec.Command("bash", "-c", fmt.Sprintf("gofmt -l -w ./")).Run()
		}
	}()
	err := servTpl.GenerateRepoInterface()
	if err != nil {
		if !ignoreErr {
			return fmt.Errorf("repoInterface err %v", err)
		} else {
			logs.PrintlnWarning(fmt.Sprintf("repoInterface err %v", err))
		}
	} else {
		logs.PrintlnSuccess("create repoInterface success")
	}
	err = servTpl.GenerateRepo()
	if err != nil {
		if !ignoreErr {
			return fmt.Errorf("repo err %v", err)
		} else {
			logs.PrintlnWarning(fmt.Sprintf("repo err %v", err))
		}
	} else {
		logs.PrintlnSuccess("create repo success")
	}
	err = servTpl.GenerateService()
	if err != nil {
		if !ignoreErr {
			return fmt.Errorf("service err %v", err)
		} else {
			logs.PrintlnWarning(fmt.Sprintf("service err %v", err))
		}
	} else {
		logs.PrintlnSuccess("create service success")
	}
	err = servTpl.GenerateController()
	if err != nil {
		if !ignoreErr {
			return fmt.Errorf("controller err %v", err)
		} else {
			logs.PrintlnWarning(fmt.Sprintf("controller err %v", err))
		}
	} else if servTpl.controllerDir != "" {
		logs.PrintlnSuccess("create controller success")
	}
	if servTpl.ViewDir != "" {
		err = servTpl.GenerateView()
		if err != nil {
			if !ignoreErr {
				return fmt.Errorf("view err %v", err)
			} else {
				logs.PrintlnWarning(fmt.Sprintf("view err %v", err))
			}
		} else {
			logs.PrintlnSuccess("create view success")
		}
	}
	return nil
}

func (servTpl ServTpl) GenerateRepoInterface() error {
	return servTpl.generateFile(servTpl.repoInterfaceTplPath, servTpl.repoInterfacePath, map[string]any{
		"ModelField":  servTpl.ModelField,
		"UniqueField": servTpl.UniqueField,
	})
}

func (servTpl ServTpl) GenerateRepo() error {
	return servTpl.generateFile(servTpl.repoTplPath, servTpl.repoPath, map[string]any{
		"ModelField":  servTpl.ModelField,
		"UniqueField": servTpl.UniqueField,
	})
}

func (servTpl ServTpl) GenerateService() error {
	return servTpl.generateFile(servTpl.servTplPath, servTpl.servPath, map[string]any{
		"ModelField":  servTpl.ModelField,
		"UniqueField": servTpl.UniqueField,
	})
}

func (servTpl ServTpl) GenerateModel() error {
	return servTpl.generateFile(servTpl.modelTplPath, servTpl.modelPath, map[string]any{
		"ModelField": servTpl.ModelField,
		"TableName":  global.SnakeString(servTpl.Model),
	})
}

func (servTpl ServTpl) GenerateController() error {
	if servTpl.controllerDir == "" {
		return nil
	}
	_package := servTpl.controllerDir
	if _package == "/" {
		_package = "controller"
	}
	return servTpl.generateFile(servTpl.controllerTplPath, servTpl.controllerPath, map[string]any{
		"Package": _package,
		"View":    servTpl.ViewDir,
	})
}

func (servTpl ServTpl) GenerateView() error {
	viewRoot := fmt.Sprintf("%s/views/%s/%s", servTpl.AppRoot, servTpl.ViewDir, servTpl.Alias)
	if !global.FileExists(viewRoot) {
		err := os.Mkdir(viewRoot, os.ModePerm)
		if err != nil {
			return err
		}
	}
	viewList := fmt.Sprintf("%s/list.html", viewRoot)
	viewItem := fmt.Sprintf("%s/item.html", viewRoot)
	err := servTpl.generateFile(servTpl.viewListTplPath, viewList, map[string]any{
		"View":       servTpl.ViewDir,
		"ModelField": servTpl.ModelField,
	})
	if err != nil {
		return err
	}
	return servTpl.generateFile(servTpl.viewItemTplPath, viewItem, map[string]any{
		"View":       servTpl.ViewDir,
		"ModelField": servTpl.ModelField,
	})
}

func (servTpl ServTpl) generateFile(tplPath, filePath string, data map[string]any) error {
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
	data["Model"] = servTpl.Model
	data["Alias"] = servTpl.Alias
	return tpl.Execute(f, data)
}
