package repositories

import (
	"gorm.io/gorm"
	"iris-init/model"
	"iris-init/repositories/repoComm"
	"iris-init/repositories/repoInterface"
)

type RolesPermissionsRepoGorm struct {
	repoComm.RepoGorm
}

func NewRolesPermissionsRepo() repoInterface.RolesPermissionsRepo {
	return &RolesPermissionsRepoGorm{repoComm.NewRepoGorm()}
}

func (rPermRepo RolesPermissionsRepoGorm) SaveByRole(role model.Roles) error {
	if role.ID == 0 {
		panic("SaveByRole role.ID is 0")
	}
	if len(role.PermIdents) == 0 {
		return rPermRepo.Orm.Where("role_id", role.ID).Delete(model.RolesPermissions{}).Error
	}
	return rPermRepo.Orm.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("role_id", role.ID).Delete(model.RolesPermissions{}).Error
		if err != nil {
			return err
		}
		rps := make([]model.RolesPermissions, 0, len(role.PermIdents))
		for _, v := range role.PermIdents {
			rps = append(rps, model.RolesPermissions{
				RoleId:          role.ID,
				PermissionIdent: v,
			})
		}
		return tx.Create(&rps).Error
	})
}

func (rPermRepo RolesPermissionsRepoGorm) GetPermissionsByRoles(roleId ...uint64) []string {
	r := []string{}
	rPermRepo.Orm.Model(model.RolesPermissions{}).
		Where("role_id in ?", roleId).
		Select("permission_ident").
		Group("permission_ident").Scan(&r)
	return r
}
