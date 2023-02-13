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
	GetByID(id uint64, _select ...string) model.{{.Model}}
	GetByIDLock(id uint64, _select ...string) (model.{{.Model}}, repoComm.ReleaseLock)
	GetByWhere(where {{.Model}}SearchWhere) model.{{.Model}}
	GetIDByWhere(where {{.Model}}SearchWhere) []uint64
}

type {{.Model}}SearchWhere struct {
    {{- range .ModelField}}
    {{.Name}}   {{.Type}}
    {{- end}}
	SelectParams repoComm.SelectFrom
}
