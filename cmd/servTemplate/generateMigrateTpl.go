package servTemplate

import (
	"fmt"
	"iris-init/global"
	"iris-init/sErr"
	"os"
	"strings"
	"text/template"
	"time"
)

type MigrateTpl struct {
	Name           string
	Models         []string
	AppRoot        string
	migrateTplPath string
	migratePath    string
}

func NewMigrateTpl(_models []string) MigrateTpl {
	if len(_models) == 0 {
		panic("model is not empty")
	}
	mt := MigrateTpl{Models: _models}
	pwd, _ := os.Getwd()
	mt.SetAppPath(pwd)
	return mt
}

func (migrateTpl *MigrateTpl) SetAppPath(AppRoot string) {
	migrateTpl.AppRoot = AppRoot
	migrateTpl.migrateTplPath = fmt.Sprintf("%s/cmd/servTemplate/migrate.tpl", migrateTpl.AppRoot)
	migrateTpl.RefreshName()
}

func (migrateTpl *MigrateTpl) addModel(models ...string) {
	migrateTpl.Models = append(migrateTpl.Models, models...)
	migrateTpl.RefreshName()
}

func (migrateTpl *MigrateTpl) SetModel(models []string) {
	migrateTpl.Models = models
	migrateTpl.RefreshName()
}

func (migrateTpl *MigrateTpl) RefreshName() {
	migrateTpl.Name = strings.Join(migrateTpl.Models, "_")
	migrateTpl.Name = strings.ToLower(migrateTpl.Name)
	migrateTpl.Name = fmt.Sprintf("%s_%s", time.Now().Format(time.DateTimeCompact), global.StringFirstUpper(migrateTpl.Name))
	migrateTpl.migratePath = fmt.Sprintf("%s/migrates/migrate_%s.go", migrateTpl.AppRoot, strings.ToLower(migrateTpl.Name))
}

func (migrateTpl MigrateTpl) GenerateFile() error {
	err := migrateTpl.generateFile(migrateTpl.migrateTplPath, migrateTpl.migratePath)
	if err != nil {
		return fmt.Errorf("repoInterface err %v", err)
	}
	return nil
}

func (migrateTpl MigrateTpl) generateFile(tplPath, filePath string) error {
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
	return tpl.Execute(f, map[string]interface{}{
		"Models":      migrateTpl.Models,
		"MigrateName": migrateTpl.Name,
	})
}
