package servTemplate

import (
	"fmt"
	"iris-init/global"
	"iris-init/sErr"
	"os"
	"text/template"
)

type ServTpl struct {
	Model                string
	Alias                string
	AppRoot              string
	servTplPath          string
	repoTplPath          string
	repoInterfaceTplPath string
	repoPath             string
	repoInterfacePath    string
	servPath             string
}

func NewServTpl(_model, alias string) ServTpl {
	if _model == "" {
		panic("model is not empty")
	}
	st := ServTpl{Model: _model, Alias: alias}
	pwd, _ := os.Getwd()
	st.SetAppPath(pwd)
	return st
}

func (servTpl *ServTpl) SetAppPath(AppRoot string) {
	servTpl.AppRoot = AppRoot
	servTpl.servTplPath = fmt.Sprintf("%s/cmd/servTemplate/services.tpl", servTpl.AppRoot)
	servTpl.repoTplPath = fmt.Sprintf("%s/cmd/servTemplate/repo.tpl", servTpl.AppRoot)
	servTpl.repoInterfaceTplPath = fmt.Sprintf("%s/cmd/servTemplate/repoInterface.tpl", servTpl.AppRoot)
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
	if servTpl.Alias == "" {
		servTpl.Alias = _m
	}
}

func (servTpl ServTpl) GenerateFile() error {
	err := servTpl.GenerateRepoInterface()
	if err != nil {
		return fmt.Errorf("repoInterface err %v", err)
	}
	err = servTpl.GenerateRepo()
	if err != nil {
		return fmt.Errorf("repo err %v", err)
	}
	err = servTpl.GenerateService()
	if err != nil {
		return fmt.Errorf("service err %v", err)
	}
	return nil
}

func (servTpl ServTpl) GenerateRepoInterface() error {
	return servTpl.generateFile(servTpl.repoInterfaceTplPath, servTpl.repoInterfacePath)
}

func (servTpl ServTpl) GenerateRepo() error {
	return servTpl.generateFile(servTpl.repoTplPath, servTpl.repoPath)
}

func (servTpl ServTpl) GenerateService() error {
	return servTpl.generateFile(servTpl.servTplPath, servTpl.servPath)
}

func (servTpl ServTpl) generateFile(tplPath, filePath string) error {
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
	return tpl.Execute(f, map[string]string{
		"Model": servTpl.Model,
		"Alias": servTpl.Alias,
	})
}
