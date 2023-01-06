package repoComm

import (
	"gorm.io/gorm"
	"iris-init/orm"
)

type RepoGorm struct {
	Orm *gorm.DB //不要直接修改 调用SetOrm方法设置
}

func NewRepoGorm() RepoGorm {
	return RepoGorm{orm.GetDb()}
}

func (repo *RepoGorm) SetOrm(orm any) {
	_orm, ok := orm.(*gorm.DB)
	if !ok {
		panic("orm must is gorm..")
	}
	repo.Orm = _orm
}

func (repo *RepoGorm) ResetOrm() {
	repo.Orm = orm.GetDb()
}

func (repo *RepoGorm) Transaction(f func() error, _repo RepoInterface, _repos ...RepoInterface) error {
	if _repos == nil || len(_repos) == 0 {
		_repos = []RepoInterface{_repo}
	} else {
		_repos = append(_repos, _repo)
	}
	return repo.Orm.Transaction(func(tx *gorm.DB) error {
		repo.SetOrm(tx)
		defer repo.ResetOrm()
		for _, _vRepo := range _repos {
			_vRepo.SetOrm(tx)
			defer _vRepo.ResetOrm()
		}
		return f()
	})
}
