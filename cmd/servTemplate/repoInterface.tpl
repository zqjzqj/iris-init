package repoInterface

import (
	"iris-init/model"
	"iris-init/repositories/repoComm"
)

type {{.Model}}Repo interface {
	repoComm.RepoInterface
	GetTotalCount(where {{.Model}}SearchWhere) int64
	GetList(where {{.Model}}SearchWhere) []model.{{.Model}}
	Delete(_model model.{{.Model}}) (rowsAffected int64, err error)
	DeleteByID({{.Pk.Name}} ...{{.Pk.Type}}) (rowsAffected int64, err error)
	Save(_model *model.{{.Model}}, _select ...string) error
	SaveOmit(_model *model.{{.Model}}, _omit ...string) error
	Create(_model *[]model.{{.Model}}) error
	GetBy{{.Pk.Name}}({{.Pk.Name}} {{.Pk.Type}}, _select ...string) model.{{.Model}}
	GetBy{{.Pk.Name}}Lock({{.Pk.Name}} {{.Pk.Type}}, _select ...string) (model.{{.Model}}, repoComm.ReleaseLock)
	GetByWhere(where {{.Model}}SearchWhere) model.{{.Model}}
	GetIDByWhere(where {{.Model}}SearchWhere) []{{.Pk.Type}}
    {{- range $key, $item := .UniqueField}}
    GetBy{{$key}}({{- range $item}}{{.NameFirstLower}} {{.Type}}, {{- end}} _select ...string) model.{{$.Model}}
    GetBy{{$key}}Lock({{- range $item}}{{.NameFirstLower}} {{.Type}}, {{- end}} _select ...string) (model.{{$.Model}}, repoComm.ReleaseLock)
    {{- end}}
    {{- range $key, $item := .IndexField}}
    GetBy{{$key}}({{- range $item}}{{.NameFirstLower}} {{.Type}}, {{- end}} _select ...string) []model.{{$.Model}}
    {{- end}}
    {{- range $key, $item := .UniqueField}}
    DeleteBy{{$key}}({{- range $item}}{{.NameFirstLower}} {{.Type}}, {{- end}}) (rowsAffected int64, err error)
    {{- end}}
    {{- range $key, $item := .IndexField}}
    DeleteBy{{$key}}({{- range $item}}{{.NameFirstLower}} {{.Type}}, {{- end}}) (rowsAffected int64, err error)
    {{- end}}
    UpdateByWhere(where {{.Model}}SearchWhere, data interface{}) (rowsAffected int64, err error)
    DeleteByWhere(where {{.Model}}SearchWhere) (rowsAffected int64, err error)
    ScanByWhere(where {{.Model}}SearchWhere, dest any) error
    ScanByOrWhere(dest any, where ...{{.Model}}SearchWhere) error
}

type {{.Model}}SearchWhere struct {
    {{- range .ModelField}}
        {{- if eq .IsStruct false}}
            {{.Name}}     string
            {{.Name}}Neq string //不等于条件
            {{.Name}}Null bool
            {{.Name}}NotNull bool
            {{.Name}}Any any
            {{- if eq .Type "string" }}
                {{.Name}}Like string
            {{- end}}
            {{.Name}}Lt   string // {{.Name}} < {{.Name}}Lt
            {{.Name}}Gt   string // {{.Name}} > {{.Name}}Gt
            {{.Name}}Elt  string // {{.Name}} <= {{.Name}}Elt
            {{.Name}}Egt  string // {{.Name}} >= {{.Name}}Egt
            {{- if eq .Type "uint8"}}
                {{.Name}}NotIn []int // not in查询
                {{.Name}}In []int // in查询
            {{- else}}
                {{.Name}}NotIn []{{.Type}} // not in查询
                {{.Name}}In []{{.Type}} // in查询
            {{- end}}
            {{.Name}}Sort string // 排序
        {{- end}}
        {{- if eq .IsSoftDelete true}}
        Unscoped_Del bool
        {{- end}}
    {{- end}}
	SelectParams repoComm.SelectFrom
}
