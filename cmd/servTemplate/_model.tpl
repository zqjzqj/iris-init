package model

import (
	"big_data_new/model/mField"
)

//主要是用于在编辑实际model的字段后 创建一个临时的model文件 用于复制ShowMap内容 完成后删除即可
type _{{.Model}} struct {
	{{- range .ModelField}}
    {{.Name}}   {{.Type}} `label:"{{.Label}}"`
    {{- end}}
}

func ({{.Alias}} _{{.Model}}) TableName() string {
	return "{{.TableName}}"
}

func ({{.Alias}} _{{.Model}}) ShowMap() map[string]interface{} {
	return map[string]interface{}{
	    {{- range .ModelField}}
        "{{.Name}}": {{$.Alias}}.{{.Name}},
        {{- end}}
	}
}
