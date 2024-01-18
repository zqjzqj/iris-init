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

func (rAdmRepo RolesAdmRepoGorm) DeleteByRoleID(RoleID ...uint64) (rowsAffected int64, err error) {
	if len(RoleID) == 1 {
		return rAdmRepo.Delete("role_id", RoleID[0])
	}
	return rAdmRepo.Delete("role_id in ?", RoleID)
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
	if adm.RolesID == "" || adm.RolesID == model.RoleAdmin {
		return rAdmRepo.Orm.Where("admin_id", adm.ID).Delete(model.RolesAdmin{}).Error
	}
	return rAdmRepo.Orm.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("admin_id", adm.ID).Delete(model.RolesAdmin{}).Error
		if err != nil {
			return err
		}
		RolesID := adm.RefreshRoleIDSlices()
		RolesID = global.RemoveDuplicateElement(RolesID)
		rAdm := make([]model.RolesAdmin, 0, len(RolesID))
		for _, v := range RolesID {
			i, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				continue
			}
			rAdm = append(rAdm, model.RolesAdmin{
				RoleID:  i,
				AdminID: adm.ID,
			})
		}
		return tx.Create(&rAdm).Error
	})
}
