package repoInterface

import (
	"iris-init/model"
	"iris-init/repositories/repoComm"
)

type AreaRepo interface {
	repoComm.RepoInterface
	GetByID(id uint, _select ...string) model.Area
	GetListByPID(pid uint, _select ...string) []model.Area
}
