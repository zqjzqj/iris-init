package repoInterface

import (
	"iris-init/model"
	"iris-init/repositories/repoComm"
)

type RolesPermissionsRepo interface {
	repoComm.RepoInterface
	SaveByRole(role model.Roles) error //当len(role.PermIdents)==0时 应当清空对应的数据
	GetPermissionsByRoles(RoleID ...uint64) []string
}
