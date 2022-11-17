package repoInterface

import (
	"jd-fxl/model"
	"jd-fxl/repositories/repoComm"
)

type AdminRepo interface {
	RepoInterface
	GetTotalCount(where AdmSearchWhere) int64
	Delete(query string, args ...interface{}) (rowsAffected int64, err error)
	GetList(where AdmSearchWhere) []model.Admin
	Save(adm *model.Admin, _select ...string) error
	GetByID(id uint64, _select ...string) model.Admin
	GetByPhone(phone string, _select ...string) model.Admin
	GetByToken(token string, _select ...string) model.Admin
	GetByUsername(username string, _select ...string) model.Admin
	GetByRealName(realName string, _select ...string) model.Admin
	GetIDByWhere(query string, args ...interface{}) []uint64
}

type AdmSearchWhere struct {
	ID           uint64
	Username     string
	SelectParams repoComm.SelectFrom
}
