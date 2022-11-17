package repositories

import (
	"jd-fxl/model"
	"jd-fxl/repositories/repoComm"
	"jd-fxl/repositories/repoInterface"
)

type AreaRepoGorm struct {
	repoComm.RepoGorm
}

func NewAreaRepo() repoInterface.AreaRepo {
	return &AreaRepoGorm{repoComm.NewRepoGorm()}
}

func (areaRepo AreaRepoGorm) GetByID(id uint, _select ...string) model.Area {
	tx := areaRepo.Orm.Where("id", id)
	if len(_select) > 0 {
		tx.Select(_select)
	}
	area := model.Area{}
	tx.First(&area)
	return area
}

func (areaRepo AreaRepoGorm) GetListByPID(pid uint, _select ...string) []model.Area {
	tx := areaRepo.Orm.Where("pid", pid)
	if len(_select) > 0 {
		tx.Select(_select)
	}
	areas := make([]model.Area, 0)
	tx.Find(&areas)
	return areas
}
