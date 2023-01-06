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
	return repoComm.SaveModel(rolesRepo.Orm, role, _select...)
}

func (rolesRepo RolesRepoGorm) DeleteByID(id ...uint64) (rowsAffected int64, err error) {
	if len(id) == 1 {
		return rolesRepo.Delete("id", id[0])
	}
	return rolesRepo.Delete("id in ?", id)
}

func (rolesRepo RolesRepoGorm) Delete(query string, args ...interface{}) (rowsAffected int64, err error) {
	r := rolesRepo.Orm.Where(query, args...).Delete(model.Roles{})
	return r.RowsAffected, r.Error
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
