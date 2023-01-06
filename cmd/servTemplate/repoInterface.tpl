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
	DeleteByWhere(query string, args ...interface{}) (rowsAffected int64, err error)
	DeleteByID(id ...uint64) (rowsAffected int64, err error)
	Save(_model *model.{{.Model}}, _select ...string) error
	SaveOmit(_model *model.{{.Model}}, _omit ...string) error
	GetByID(id uint64, _select ...string) model.{{.Model}}
	GetIDByWhere(query string, args ...interface{}) []uint64
}

type {{.Model}}SearchWhere struct {
    ID           uint64
	SelectParams repoComm.SelectFrom
}
