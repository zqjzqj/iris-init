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
	TypeOrigin     string
	Label          string
	ValidateLabel  string //主要用于service的验证标签
	Unique         string
	Index          string //用于生成 如 GetByField() []model.Model 这样的方法
	OnlyRead       bool
	Search         bool
	IsNumber       bool
	IsPk           bool //是否是主键
	IsStruct       bool
	ReferencesName string
}

func GetIndexFields(fields []Field) map[string][]Field {
	r := make(map[string][]Field)
	for _, v := range fields {
		if v.Index != "" {
			if r[v.Index] == nil {
				r[v.Index] = make([]Field, 0, 3)
			}
			r[v.Index] = append(r[v.Index], v)
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

func GetReferences(fields []Field) []Field {
	r := make([]Field, 0)
	for _, v := range fields {
		if !v.IsStruct {
			continue
		}
		for _, v2 := range fields {
			if v.Name+"ID" == v2.Name {
				v.ReferencesName = v2.Name
				r = append(r, v)
				break
			}
		}
	}
	return r
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
			gormLabelStruct := GetValidateStrByGormLabel(_tag.Get("gorm"))
			_f := Field{
				Name:       ref.Field(i).Name,
				Type:       ref.Field(i).Type.String(),
				TypeOrigin: ref.Field(i).Type.String(),
				Label:      ref.Field(i).Tag.Get("label"),
				Unique:     gormLabelStruct.Unique,
				Index:      gormLabelStruct.Index,
				IsPk:       gormLabelStruct.IsPk,
				IsNumber:   global.IsNumber(ref.Field(i).Type),
				IsStruct:   ref.Field(i).Type.Kind() == reflect.Struct,
			}
			if _f.Label == "" {
				_f.Label = gormLabelStruct.Comment
			}
			_f.NameFirstLower = global.StringFirstLower(_f.Name)
			//关键字处理
			if _f.NameFirstLower == "type" {
				_f.NameFirstLower = "_type"
			}
			if _f.Type == "sql.NullString" {
				_f.Type = "string"
				_f.IsStruct = false
			}
			if _f.Type == "time.Time" {
				_f.Type = "string"
				_f.IsStruct = false
			}
			if gormLabelStruct.Validate != "" && _f.Name != "ID" {
				_f.ValidateLabel = `validate:"` + gormLabelStruct.Validate + `"`
			}
			_f.NameSnake = global.SnakeString(_f.Name)
			if _f.Label == "" {
				_f.Label = _f.Name
			}
			if _tag.Get("Unique") != "" {
				_f.Unique = _tag.Get("Unique")
			}
			if _tag.Get("Index") != "" {
				_f.Index = _tag.Get("Index")
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

func GetValidateStrByGormLabel(gormLabel string) GormLabelStruct {
	_gormLabel := strings.Split(gormLabel, ";")
	var required = true
	gormLabelStruct := GormLabelStruct{}
	for _, v := range _gormLabel {
		if strings.HasPrefix(v, "size:") {
			gormLabelStruct.Validate += "max=" + strings.TrimLeft(v, "size:") + ","
		}
		if strings.HasPrefix(v, "default:") {
			required = false
		}
		if strings.HasPrefix(v, "primarykey") {
			gormLabelStruct.IsPk = true
		}
		if strings.HasPrefix(v, "comment:") {
			gormLabelStruct.Comment = strings.TrimLeft(v, "comment:")
		}
		//如果包含索引
		if strings.HasPrefix(v, "index:") {
			_index := strings.Split(v, ",")
			if strings.Contains(v, ",unique") { //唯一索引
				gormLabelStruct.Unique = strings.TrimLeft(_index[0], "index:")
			} else { //普通的索引
				gormLabelStruct.Index = strings.TrimLeft(_index[0], "index:")
			}
		}
	}
	if required {
		gormLabelStruct.Validate += "required"
	}
	if gormLabelStruct.Validate != "" {
		gormLabelStruct.Validate = strings.TrimRight(gormLabelStruct.Validate, ",")
	}
	return gormLabelStruct
}

type GormLabelStruct struct {
	Validate string
	Unique   string
	Index    string
	Comment  string
	IsPk     bool
}
