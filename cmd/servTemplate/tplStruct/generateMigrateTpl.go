package tplStruct

import (
	"big_data_new/global"
	"big_data_new/logs"
	"big_data_new/migrates"
	"big_data_new/sErr"
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/template"
	"time"
)

type MigrateTpl struct {
	Name                   string
	Models                 []string
	AppRoot                string
	migrateTplPath         string
	migratePath            string
	aMigratesSlicesTplPath string
	aMigratesSlicesPath    string
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
	migrateTpl.aMigratesSlicesTplPath = fmt.Sprintf("%s/cmd/servTemplate/aMigratesSlices.tpl", migrateTpl.AppRoot)
	migrateTpl.aMigratesSlicesPath = fmt.Sprintf("%s/migrates/aMigratesSlices.go", migrateTpl.AppRoot)
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
	for _, v := range migrateTpl.Models {
		migrateTpl.Name = migrateTpl.Name + "_" + v
	}
	migrateTpl.Name = strings.TrimLeft(migrateTpl.Name, "_")
	migrateTpl.Name = fmt.Sprintf("%s_%s", time.Now().Format(global.DateTimeFormatStrCompact), global.StringFirstUpper(migrateTpl.Name))
	migrateTpl.migratePath = fmt.Sprintf("%s/migrates/migrate_%s.go", migrateTpl.AppRoot, global.StringFirstLower(migrateTpl.Name))
}

func (migrateTpl MigrateTpl) GenerateFile() error {
	err := migrateTpl.generateFile(migrateTpl.migrateTplPath, migrateTpl.migratePath, nil)
	if err != nil {
		return fmt.Errorf("repoInterface err %v", err)
	}
	_aMigratesSlices := make([]string, 0, 10)
	for _, v := range migrates.MM {
		_aMigratesSlices = append(_aMigratesSlices, fmt.Sprintf("%s{}", reflect.TypeOf(v).Name()))
	}
	_aMigratesSlices = append(_aMigratesSlices, fmt.Sprintf("Migrate_%s{}", migrateTpl.Name))

	_tmpAMigratesSlicesPath := fmt.Sprintf("%s.tmp", migrateTpl.aMigratesSlicesPath)
	err = os.Rename(migrateTpl.aMigratesSlicesPath, _tmpAMigratesSlicesPath)
	if err != nil {
		logs.PrintErr("create aMigratesSlices fail")
		return err
	}
	err = migrateTpl.generateFile(migrateTpl.aMigratesSlicesTplPath, migrateTpl.aMigratesSlicesPath, map[string]any{
		"AMigratesSlices": _aMigratesSlices,
	})
	if err == nil {
		_ = os.Remove(_tmpAMigratesSlicesPath)
		return nil
	}
	return os.Rename(_tmpAMigratesSlicesPath, migrateTpl.aMigratesSlicesPath)
}

func (migrateTpl MigrateTpl) generateFile(tplPath, filePath string, data map[string]any) error {
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
	data["Models"] = migrateTpl.Models
	data["MigrateName"] = migrateTpl.Name
	return tpl.Execute(f, data)
}
