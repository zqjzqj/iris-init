package repositories

import (
	"gorm.io/gorm"
	"iris-init/global"
	"iris-init/model"
	"iris-init/repositories/repoComm"
	"iris-init/repositories/repoInterface"
	"strconv"
)

type RolesAdmRepoGorm struct {
	repoComm.RepoGorm
}

func NewRolesAdmRepo() repoInterface.RolesAdmRepo {
	return &RolesAdmRepoGorm{repoComm.NewRepoGorm()}
}

func (rAdmRepo RolesAdmRepoGorm) DeleteByRoleID(roleID ...uint64) (rowsAffected int64, err error) {
	if len(roleID) == 1 {
		return rAdmRepo.Delete("role_id", roleID[0])
	}
	return rAdmRepo.Delete("role_id in ?", roleID)
}

func (rAdmRepo RolesAdmRepoGorm) DeleteByAdmID(admID ...uint64) (rowsAffected int64, err error) {
	if len(admID) == 1 {
		return rAdmRepo.Delete("admin_id", admID[0])
	}
	return rAdmRepo.Delete("admin_id in ?", admID)
}

func (rAdmRepo RolesAdmRepoGorm) Delete(query string, args ...interface{}) (rowsAffected int64, err error) {
	tx := rAdmRepo.Orm.Where(query, args...).Delete(model.RolesAdmin{})
	return tx.RowsAffected, tx.Error
}

func (rAdmRepo RolesAdmRepoGorm) SaveByAdm(adm model.Admin) error {
	if adm.ID == 0 {
		panic("SaveByAdm adm.ID is 0")
	}
	if adm.RolesId == "" || adm.RolesId == model.RoleAdmin {
		return rAdmRepo.Orm.Where("admin_id", adm.ID).Delete(model.RolesAdmin{}).Error
	}
	return rAdmRepo.Orm.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("admin_id", adm.ID).Delete(model.RolesAdmin{}).Error
		if err != nil {
			return err
		}
		rolesId := adm.RefreshRoleIDSlices()
		rolesId = global.RemoveDuplicateElement(rolesId)
		rAdm := make([]model.RolesAdmin, 0, len(rolesId))
		for _, v := range rolesId {
			i, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				continue
			}
			rAdm = append(rAdm, model.RolesAdmin{
				RoleId:  i,
				AdminId: adm.ID,
			})
		}
		return tx.Create(&rAdm).Error
	})
}
