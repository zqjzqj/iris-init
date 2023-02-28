package tplStruct

import (
	"iris-init/global"
	"reflect"
	"strings"
)

type Field struct {
	Name          string
	NameSnake     string
	Type          string
	Label         string
	ValidateLabel string //主要用于service的验证标签
}

func RefStructField(_struct any) []Field {
	var ref reflect.Type
	v, ok := _struct.(reflect.Type)
	if ok {
		ref = v
	} else {
		ref = reflect.TypeOf(_struct)
	}
	fields := make([]Field, 0, ref.NumField())
	for i := 0; i < ref.NumField(); i++ {
		//扩展字段将被忽略
		if ref.Field(i).Type.Name() == "FieldsExtendsJsonType" {
			continue
		}
		//对于组合的结构 只反射mField包下的
		if ref.Field(i).Type.Kind() == reflect.Struct &&
			strings.HasSuffix(ref.Field(i).Type.PkgPath(), "model/mField") || ref.Field(i).Tag.Get("ref") == "true" {
			fields = append(fields, RefStructField(ref.Field(i).Type)...)
		} else {
			_validate := GetValidateStrByGormLabel(ref.Field(i).Tag.Get("gorm"))
			_f := Field{
				Name:  ref.Field(i).Name,
				Type:  ref.Field(i).Type.String(),
				Label: ref.Field(i).Tag.Get("label"),
			}
			if _validate != "" && _f.Name != "ID" {
				_f.ValidateLabel = `validate:"` + _validate + `"`
			}
			_f.NameSnake = global.SnakeString(_f.Name)
			if _f.Label == "" {
				_f.Label = _f.Name
			}
			fields = append(fields, _f)
		}
	}
	return fields
}

func GetValidateStrByGormLabel(gormLabel string) string {
	_gormLabel := strings.Split(gormLabel, ";")
	validate := ""
	var required = true
	for _, v := range _gormLabel {
		if strings.HasPrefix(v, "size:") {
			validate += "max=" + strings.TrimLeft(v, "size:") + ","
		}
		if strings.HasPrefix(v, "default:") {
			required = false
		}
	}
	if required {
		validate += "required"
	}
	if validate != "" {
		validate = strings.TrimRight(validate, ",")
	}
	return validate
}
