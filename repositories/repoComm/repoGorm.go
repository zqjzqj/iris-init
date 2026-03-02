package repoComm

import (
	"gorm.io/gorm"
	"iris-init/orm"
)

type RepoGorm struct {
	Orm      *gorm.DB //不要直接修改 调用SetOrm方法设置
	OrmLasts []*gorm.DB
}

func NewRepoGorm() RepoGorm {
	return RepoGorm{Orm: orm.GetDb(), OrmLasts: make([]*gorm.DB, 0, 3)}
}

func (repo *RepoGorm) GetOrm() any {
	return repo.Orm
}

func (repo *RepoGorm) SetOrm(orm any) {
	_orm, ok := orm.(*gorm.DB)
	if !ok {
		panic("orm must is gorm..")
	}
	repo.OrmLasts = append(repo.OrmLasts, repo.Orm)
	repo.Orm = _orm
}

func (repo *RepoGorm) ResetLastOrm() {
	if len(repo.OrmLasts) == 0 {
		repo.Orm = orm.GetDb()
	} else {
		_index := len(repo.OrmLasts) - 1
		repo.Orm = repo.OrmLasts[_index]
		repo.OrmLasts = repo.OrmLasts[:_index]
	}
}

func (repo *RepoGorm) ResetOrm() {
	repo.Orm = orm.GetDb()
	repo.OrmLasts = make([]*gorm.DB, 0, 3)
}

func (repo *RepoGorm) Transaction(f func() error, _repos ...RepoInterface) error {
	return repo.Orm.Transaction(func(tx *gorm.DB) error {
		repo.SetOrm(tx)
		defer repo.ResetLastOrm()
		for _, _vRepo := range _repos {
			_vRepo.SetOrm(tx)
			defer _vRepo.ResetLastOrm()
		}
		return f()
	})
}
