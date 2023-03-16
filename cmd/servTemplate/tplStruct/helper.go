package tplStruct

import (
	"iris-init/global"
	"reflect"
	"strings"
)

type Field struct {
	Name           string
	NameSnake      string
	NameFirstLower string
	Type           string
	Label          string
	ValidateLabel  string //主要用于service的验证标签
	Unique         string
	OnlyRead       bool
	Search         bool
}

func GetUniqueFields(fields []Field) map[string][]Field {
	r := make(map[string][]Field)
	for _, v := range fields {
		if v.Unique != "" {
			if r[v.Unique] == nil {
				r[v.Unique] = make([]Field, 0, 3)
			}
			r[v.Unique] = append(r[v.Unique], v)
		}
	}
	if len(r) == 0 {
		return nil
	}
	rr := make(map[string][]Field, len(r))
	for _, v := range r {
		name := ""
		for _, vv := range v {
			name += vv.Name + "_"
		}
		name = strings.TrimRight(name, "_")
		rr[name] = v
	}
	return rr
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
			_tag := ref.Field(i).Tag
			_validate, _unique := GetValidateStrByGormLabel(_tag.Get("gorm"))
			_f := Field{
				Name:   ref.Field(i).Name,
				Type:   ref.Field(i).Type.String(),
				Label:  ref.Field(i).Tag.Get("label"),
				Unique: _unique,
			}
			_f.NameFirstLower = global.StringFirstLower(_f.Name)
			if _validate != "" && _f.Name != "ID" {
				_f.ValidateLabel = `validate:"` + _validate + `"`
			}
			_f.NameSnake = global.SnakeString(_f.Name)
			if _f.Label == "" {
				_f.Label = _f.Name
			}
			if _tag.Get("Unique") != "" {
				_f.Unique = _tag.Get("Unique")
			}
			if _tag.Get("OnlyRead") == "true" {
				_f.OnlyRead = true
			}
			if _tag.Get("Search") == "true" {
				_f.Search = true
			}
			fields = append(fields, _f)
		}
	}
	return fields
}

func GetValidateStrByGormLabel(gormLabel string) (validate string, unique string) {
	_gormLabel := strings.Split(gormLabel, ";")
	var required = true
	for _, v := range _gormLabel {
		if strings.HasPrefix(v, "size:") {
			validate += "max=" + strings.TrimLeft(v, "size:") + ","
		}
		if strings.HasPrefix(v, "default:") {
			required = false
		}
		//如果包含唯一索引
		if strings.HasPrefix(v, "index:") && strings.Contains(v, "unique") {
			_unique := strings.Split(v, ",")
			unique = strings.TrimLeft(_unique[0], "index:")
		}
	}
	if required {
		validate += "required"
	}
	if validate != "" {
		validate = strings.TrimRight(validate, ",")
	}
	return
}
