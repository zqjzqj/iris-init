package repoComm

import (
	"jd-fxl/orm"
	"gorm.io/gorm"
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
