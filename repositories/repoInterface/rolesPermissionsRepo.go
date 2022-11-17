package repoInterface

import "jd-fxl/model"

type RolesPermissionsRepo interface {
	RepoInterface
	SaveByRole(role model.Roles) error //当len(role.PermIdents)==0时 应当清空对应的数据
	GetPermissionsByRoles(roleId ...uint64) []string
}
