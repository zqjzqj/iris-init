package repositories

import (
	"gorm.io/gorm"
	"iris-init/model"
	"iris-init/repositories/repoComm"
	"iris-init/repositories/repoInterface"
)

type AdminRepoGorm struct {
	repoComm.RepoGorm
}

func NewAdminRepo() repoInterface.AdminRepo {
	return &AdminRepoGorm{repoComm.NewRepoGorm()}
}

func (admRepo AdminRepoGorm) Save(admin *model.Admin, _select ...string) error {
	return repoComm.SaveModel(admRepo.Orm, admin, _select...)
}

func (admRepo AdminRepoGorm) DeleteByID(id ...uint64) (rowsAffected int64, err error) {
	if len(id) == 1 {
		return admRepo.Delete("id", id[0])
	}
	return admRepo.Delete("id in ?", id)
}

func (admRepo AdminRepoGorm) Delete(query string, args ...interface{}) (rowsAffected int64, err error) {
	tx := admRepo.Orm.Where(query, args...).Delete(model.Admin{})
	return tx.RowsAffected, tx.Error
}

func (admRepo AdminRepoGorm) GetSearchWhereTx(where repoInterface.AdmSearchWhere, tx0 *gorm.DB) *gorm.DB {
	var tx *gorm.DB
	if tx0 == nil {
		tx = admRepo.Orm.Model(model.Admin{})
	} else {
		tx = tx0.Model(model.Admin{})
	}
	if where.ID > 0 {
		tx.Where("id", where.ID)
	}
	if where.Username != "" {
		tx.Where("username like ?", where.Username+"%")
	}
	return tx
}

// 返回数据总数
func (admRepo AdminRepoGorm) GetTotalCount(where repoInterface.AdmSearchWhere) int64 {
	tx := admRepo.GetSearchWhereTx(where, nil)
	var r int64
	tx.Count(&r)
	return r
}

func (admRepo AdminRepoGorm) GetList(where repoInterface.AdmSearchWhere) []model.Admin {
	adm := make([]model.Admin, 0, where.SelectParams.RetSize)
	tx := admRepo.GetSearchWhereTx(where, nil)
	where.SelectParams.SetTxGorm(tx).Find(&adm)
	return adm
}

func (admRepo AdminRepoGorm) GetByID(id uint64, _select ...string) model.Admin {
	if id == 0 {
		return model.Admin{}
	}
	adm := model.Admin{}
	tx := admRepo.Orm.Where("id", id)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.First(&adm)
	return adm
}

func (admRepo AdminRepoGorm) GetByRealName(realName string, _select ...string) model.Admin {
	if realName == "" {
		return model.Admin{}
	}
	adm := model.Admin{}
	tx := admRepo.Orm.Where("real_name", realName)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.First(&adm)
	return adm
}

func (admRepo AdminRepoGorm) GetByPhone(phone string, _select ...string) model.Admin {
	if phone == "" {
		return model.Admin{}
	}
	adm := model.Admin{}
	tx := admRepo.Orm.Where("phone", phone)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.First(&adm)
	return adm
}

func (admRepo AdminRepoGorm) GetByToken(token string, _select ...string) model.Admin {
	if token == "" {
		return model.Admin{}
	}
	adm := model.Admin{}
	tx := admRepo.Orm.Where("token", token)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.First(&adm)
	return adm
}

func (admRepo AdminRepoGorm) GetByUsername(username string, _select ...string) model.Admin {
	if username == "" {
		return model.Admin{}
	}
	adm := model.Admin{}
	tx := admRepo.Orm.Where("username", username)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.First(&adm)
	return adm
}
