package repoInterface

import (
	"jd-fxl/model"
	"jd-fxl/repositories/repoComm"
)

type RolesRepo interface {
	RepoInterface
	Save(role *model.Roles, _select ...string) error //要根据 role.PermIdents 来保存更新角色权限
	Delete(query string, args ...interface{}) (rowsAffected int64, err error)
	GetByID(id uint64, _select ...string) model.Roles
	GetByWhere(query string, args ...any) model.Roles
	GetList(where RolesSearchWhere) []model.Roles
	GetRolesByID(id ...uint64) []model.Roles
}

type RolesSearchWhere struct {
	RoleID       uint64
	Name         string
	SelectParams repoComm.SelectFrom
}
