package repoInterface

import (
	"iris-init/model"
	"iris-init/repositories/repoComm"
)

type {{.Model}}Repo interface {
	RepoInterface
	GetTotalCount(where {{.Model}}SearchWhere) int64
	GetList(where {{.Model}}SearchWhere) []model.{{.Model}}
	Delete(query string, args ...interface{}) (rowsAffected int64, err error)
	Save(_model *model.{{.Model}}, _select ...string) error
	GetByID(id uint64, _select ...string) model.{{.Model}}
	GetIDByWhere(query string, args ...interface{}) []uint64
}

type {{.Model}}SearchWhere struct {
	SelectParams repoComm.SelectFrom
}