package repoInterface

import "jd-fxl/model"

type AreaRepo interface {
	GetByID(id uint, _select ...string) model.Area
	GetListByPID(pid uint, _select ...string) []model.Area
}
