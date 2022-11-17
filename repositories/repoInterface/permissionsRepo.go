package repoInterface

import (
	"jd-fxl/model"
	"jd-fxl/repositories/repoComm"
)

type PermissionsRepo interface {
	RepoInterface
	Save(perm *model.Permissions, _select ...string) error
	GetByIdent(ident string, _select ...string) model.Permissions
	GetByID(id uint64, _select ...string) model.Permissions
	GetListAsMenu(idents []string) []model.Permissions
	GetListPreloadChildren_2() []model.Permissions
	TruncateTable()
	GetOrCreatePermissionByName(name string, pid uint64, level uint8) (model.Permissions, error)
	GetList(where PermissionsSearchWhere) []model.Permissions
	GetListPreloadChildren(where PermissionsSearchWhere) []model.Permissions
}

type PermissionsSearchWhere struct {
	ID         uint64
	Pid        int64
	Level      int8
	Name       string
	Ident      []string
	SelectFrom repoComm.SelectFrom
}
