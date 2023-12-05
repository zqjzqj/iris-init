package repoInterface

import (
	"big_data_new/model"
	"big_data_new/repositories/repoComm"
)

type RolesPermissionsRepo interface {
	repoComm.RepoInterface
	SaveByRole(role model.Roles) error //当len(role.PermIdents)==0时 应当清空对应的数据
	GetPermissionsByRoles(roleId ...uint64) []string
}
