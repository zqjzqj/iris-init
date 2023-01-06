package repoInterface

import (
	"iris-init/model"
	"iris-init/repositories/repoComm"
)

type RolesRepo interface {
	repoComm.RepoInterface
	Save(role *model.Roles, _select ...string) error //要根据 role.PermIdents 来保存更新角色权限
	Delete(query string, args ...interface{}) (rowsAffected int64, err error)
	DeleteByID(id ...uint64) (rowsAffected int64, err error)
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
