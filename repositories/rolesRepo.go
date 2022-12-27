package repositories

import (
	"gorm.io/gorm"
	"iris-init/model"
	"iris-init/repositories/repoComm"
	"iris-init/repositories/repoInterface"
)

type RolesRepoGorm struct {
	repoComm.RepoGorm
}

func NewRolesRepo() repoInterface.RolesRepo {
	return &RolesRepoGorm{repoComm.NewRepoGorm()}
}

func (rolesRepo RolesRepoGorm) Save(role *model.Roles, _select ...string) error {
	if len(role.PermIdents) == 0 {
		return repoComm.SaveModel(rolesRepo.Orm, role, _select...)
	}
	return rolesRepo.Orm.Transaction(func(tx *gorm.DB) error {
		err := repoComm.SaveModel(tx, role, _select...)
		if err != nil {
			return err
		}
		rPermRepo := NewRolesPermissionsRepo()
		rPermRepo.SetOrm(tx)
		defer rolesRepo.ResetOrm()
		return rPermRepo.SaveByRole(*role)
	})
}

func (rolesRepo RolesRepoGorm) Delete(query string, args ...interface{}) (rowsAffected int64, err error) {
	role := rolesRepo.GetByWhere(query, args...)
	if role.ID == 0 {
		return 0, nil
	}
	err = rolesRepo.Orm.Transaction(func(tx *gorm.DB) error {
		rPermRepo := NewRolesPermissionsRepo()
		rPermRepo.SetOrm(tx)
		defer rolesRepo.ResetOrm()

		role.PermIdents = nil
		_err := rPermRepo.SaveByRole(role)
		if _err != nil {
			return _err
		}
		r := tx.Delete(&role)
		rowsAffected = r.RowsAffected
		return r.Error
	})
	return rowsAffected, err
}

func (rolesRepo RolesRepoGorm) GetSearchWhereTx(where repoInterface.RolesSearchWhere, tx0 *gorm.DB) *gorm.DB {
	var tx *gorm.DB
	if tx0 == nil {
		tx = rolesRepo.Orm.Model(model.Roles{})
	} else {
		tx = tx0.Model(model.Roles{})
	}
	if where.RoleID > 0 {
		tx.Where("id", where.RoleID)
	}
	if where.Name != "" {
		tx.Where("name like ?", where.Name+"%")
	}
	return tx
}

func (rolesRepo RolesRepoGorm) GetList(where repoInterface.RolesSearchWhere) []model.Roles {
	roles := make([]model.Roles, 0, where.SelectParams.RetSize)
	tx := rolesRepo.GetSearchWhereTx(where, nil)
	where.SelectParams.SetTxGorm(tx).Find(&roles)
	return roles
}

func (rolesRepo RolesRepoGorm) GetByWhere(query string, args ...any) model.Roles {
	role := model.Roles{}
	rolesRepo.Orm.Where(query, args...).First(&role)
	return role
}

func (rolesRepo RolesRepoGorm) GetByID(id uint64, _select ...string) model.Roles {
	if id == 0 {
		return model.Roles{}
	}
	roles := model.Roles{}
	tx := rolesRepo.Orm.Where("id", id)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.First(&roles)
	return roles
}

func (rolesRepo RolesRepoGorm) GetRolesByID(id ...uint64) []model.Roles {
	r := make([]model.Roles, 0, len(id))
	rolesRepo.Orm.Where("id in ?", id).Find(&r)
	return r
}
