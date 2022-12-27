package repositories

import (
	"gorm.io/gorm"
	"iris-init/model"
	"iris-init/repositories/repoComm"
	"iris-init/repositories/repoInterface"
	"strconv"
	"strings"
)

type RolesAdmRepoGorm struct {
	repoComm.RepoGorm
}

func NewRolesAdmRepo() repoInterface.RolesAdmRepo {
	return &RolesAdmRepoGorm{repoComm.NewRepoGorm()}
}

func (rAdmRepo RolesAdmRepoGorm) SaveByAdm(adm model.Admin) error {
	if adm.ID == 0 {
		panic("SaveByAdm adm.ID is 0")
	}
	if adm.RolesId == "" {
		return rAdmRepo.Orm.Where("admin_id", adm.ID).Delete(model.RolesAdmin{}).Error
	}
	return rAdmRepo.Orm.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("admin_id", adm.ID).Delete(model.RolesAdmin{}).Error
		if err != nil {
			return err
		}
		rolesId := strings.Split(adm.RolesId, ",")
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
