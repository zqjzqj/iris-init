package services

import (
	"jd-fxl/model"
	"jd-fxl/repositories"
	"jd-fxl/repositories/repoInterface"
	"jd-fxl/sErr"
)

func NewAreaService() AreaService {
	return AreaService{repo: repositories.NewAreaRepo()}
}

type AreaService struct {
	repo repoInterface.AreaRepo
}

func (areaServ AreaService) GetListByPID(pid uint) []model.Area {
	return areaServ.repo.GetListByPID(pid)
}

func (areaServ AreaService) CheckAreaID(proID, cityID, regionID uint) error {
	area := areaServ.repo.GetByID(regionID)
	if area.ID == 0 {
		return sErr.ErrInvalidRegionID
	}
	if area.Pid != cityID {
		return sErr.ErrInvalidRegionID
	}
	area = areaServ.repo.GetByID(cityID)
	if area.Pid != proID {
		return sErr.ErrInvalidProvinceID
	}
	return nil
}
