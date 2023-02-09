package repositories

import (
	"gorm.io/gorm"
	"iris-init/model"
	"iris-init/repositories/repoComm"
	"iris-init/repositories/repoInterface"
)

type PermissionsRepoGorm struct {
	repoComm.RepoGorm
}

func NewPermissionsRepo() repoInterface.PermissionsRepo {
	return &PermissionsRepoGorm{repoComm.NewRepoGorm()}
}

// 截断表
func (permRepo PermissionsRepoGorm) TruncateTable() {
	permRepo.Orm.Exec("truncate table " + new(model.Permissions).TableName())
}

func (permRepo PermissionsRepoGorm) GetByID(id uint64, _select ...string) model.Permissions {
	if id == 0 {
		return model.Permissions{}
	}
	perm := model.Permissions{}
	if len(_select) > 0 {
		permRepo.Orm.Select(_select).Where("id", id).First(&perm)
		return perm
	}
	permRepo.Orm.Where("id", id).First(&perm)
	return perm
}

func (permRepo PermissionsRepoGorm) GetByIdent(ident string, _select ...string) model.Permissions {
	if ident == "" {
		return model.Permissions{}
	}
	perm := model.Permissions{}
	if len(_select) > 0 {
		permRepo.Orm.Select(_select).Where("ident", ident).First(&perm)
		return perm
	}
	permRepo.Orm.Where("ident", ident).First(&perm)
	return perm
}

func (permRepo PermissionsRepoGorm) Save(perm *model.Permissions, _select ...string) error {
	return repoComm.SaveModel(permRepo.Orm, perm, _select...)
}

func (permRepo PermissionsRepoGorm) GetSearchWhereTx(where repoInterface.PermissionsSearchWhere, tx0 *gorm.DB) *gorm.DB {
	var tx *gorm.DB
	if tx0 == nil {
		tx = permRepo.Orm.Model(model.Permissions{})
	} else {
		tx = tx0.Model(model.Permissions{})
	}
	if where.ID > 0 {
		tx.Where("id", where.ID)
	}
	if where.Level > -1 {
		tx.Where("level", where.Level)
	}
	if where.Pid > -1 {
		tx.Where("pid", where.Pid)
	}
	if where.Name != "" {
		tx.Where("name like ?", where.Name+"%")
	}
	if len(where.Ident) > 0 {
		tx.Where("ident in ?", where.Ident)
	}
	return tx
}

func (permRepo PermissionsRepoGorm) GetListPreloadChildren_2() []model.Permissions {
	return permRepo.GetList(repoInterface.PermissionsSearchWhere{
		Pid: 0,
		SelectFrom: repoComm.SelectFrom{
			RetSize: 30,
			OrderBy: []repoComm.OrderByParams{{
				Column: "sort",
				Desc:   false,
			}, {
				Column: "id",
				Desc:   false,
			}},
			Preload: []repoComm.PreloadParams{
				{
					Query: "Children.Children",
				},
			},
		},
	})
}

func (permRepo PermissionsRepoGorm) GetListAsMenu(idents []string) []model.Permissions {
	return permRepo.GetList(repoInterface.PermissionsSearchWhere{
		Pid:   0,
		Ident: idents,
		SelectFrom: repoComm.SelectFrom{
			RetSize: 30,
			OrderBy: []repoComm.OrderByParams{{
				Column: "sort",
				Desc:   false,
			}, {
				Column: "id",
				Desc:   false,
			}},
			Preload: []repoComm.PreloadParams{
				{
					Query: "Children",
					Args: []interface{}{
						func(tx *gorm.DB) *gorm.DB {
							if idents == nil { //不限制权限idents
								return tx.Order("sort asc, id asc")
							}
							return tx.Order("sort asc, id asc").Where("ident in ?", idents)
						},
					},
				},
			},
		},
	})
}

func (permRepo PermissionsRepoGorm) GetListPreloadChildren(where repoInterface.PermissionsSearchWhere) []model.Permissions {
	where.SelectFrom.Preload = []repoComm.PreloadParams{
		{Query: "Children"},
	}
	return permRepo.GetList(where)
}

func (permRepo PermissionsRepoGorm) GetList(where repoInterface.PermissionsSearchWhere) []model.Permissions {
	var perm []model.Permissions
	if where.SelectFrom.Limit > 0 {
		perm = make([]model.Permissions, 0, where.SelectFrom.Limit)
	} else {
		if where.SelectFrom.RetSize > 0 {
			perm = make([]model.Permissions, 0, where.SelectFrom.RetSize)
		} else {
			perm = make([]model.Permissions, 0, 15)
		}
	}
	tx := permRepo.GetSearchWhereTx(where, nil)
	where.SelectFrom.SetTxGorm(tx).Find(&perm)
	return perm
}

func (permRepo PermissionsRepoGorm) GetOrCreatePermissionByName(name string, pid uint64, level uint8, sort uint) (model.Permissions, error) {
	perm := model.Permissions{}
	if permRepo.Orm.
		Where("pid", pid).
		Where("name", name).
		First(&perm).RowsAffected == 0 {
		//写入
		perm.Pid = pid
		perm.Name = name
		perm.Level = level
		perm.Sort = sort
		perm.GenerateIdent()
		err := permRepo.Save(&perm)
		return perm, err
	}
	return perm, nil
}
