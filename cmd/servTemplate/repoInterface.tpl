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
	DeleteByID(id ...uint64) (rowsAffected int64, err error)
	Save(_model *model.{{.Model}}, _select ...string) error
	SaveOmit(_model *model.{{.Model}}, _omit ...string) error
	Create(_model *[]model.{{.Model}}) error
	GetByID(id uint64, _select ...string) model.{{.Model}}
	GetByIDLock(id uint64, _select ...string) (model.{{.Model}}, repoComm.ReleaseLock)
	GetByWhere(where {{.Model}}SearchWhere) model.{{.Model}}
	GetIDByWhere(where {{.Model}}SearchWhere) []uint64
    {{- range $key, $item := .UniqueField}}
    GetBy{{$key}}({{- range $item}}{{.NameFirstLower}} {{.Type}}, {{- end}} _select ...string) model.{{$.Model}}
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
}

type {{.Model}}SearchWhere struct {
    {{- range .ModelField}}
    {{.Name}}   string
    {{- if eq .Type "string" }}
    {{.Name}}Like string
    {{- end}}
    {{- end}}
	SelectParams repoComm.SelectFrom
}
