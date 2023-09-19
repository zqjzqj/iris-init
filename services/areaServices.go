package services

import (
	"iris-init/model"
	"iris-init/repositories"
	"iris-init/repositories/repoComm"
	"iris-init/repositories/repoInterface"
	"iris-init/sErr"
	"sync"
	"time"
)

var areaLayeredList []model.Area
var areaLayeredListMutex sync.Mutex
var areaUpdateIng bool

// 作用于修改地区库后设置为true 方便更新缓存
// 暂时还没用 因为还没有编辑地区的功能
var AreaIsUpdate = false

func NewAreaService() AreaService {
	return AreaService{repo: repositories.NewAreaRepo()}
}

func NewAreaServiceByOrm(orm any) AreaService {
	r := AreaService{repo: repositories.NewAreaRepo()}
	r.repo.SetOrm(orm)
	return r
}

func NewAreaServiceByRepo(repo repoInterface.AreaRepo) AreaService {
	return AreaService{repo: repo}
}

type AreaService struct {
	repo repoInterface.AreaRepo
}

func (areaServ AreaService) GetListByPID(pid uint) []model.Area {
	return areaServ.repo.GetListByPID(pid)
}

func (areaServ AreaService) GetByID(id uint, _select ...string) model.Area {
	return areaServ.repo.GetByID(uint64(id), _select...)
}

func (areaServ AreaService) GetByWhere(where repoInterface.AreaSearchWhere) model.Area {
	return areaServ.repo.GetByWhere(where)
}

func (areaServ AreaService) ScanByWhere(where repoInterface.AreaSearchWhere, dest any) error {
	return areaServ.repo.ScanByWhere(where, dest)
}

func (areaServ AreaService) ScanByOrWhere(dest any, where ...repoInterface.AreaSearchWhere) error {
	return areaServ.repo.ScanByOrWhere(dest, where...)
}

func (areaServ AreaService) UpdateByWhere(where repoInterface.AreaSearchWhere, data interface{}) (rowsAffected int64, err error) {
	return areaServ.repo.UpdateByWhere(where, data)
}

func (areaServ AreaService) GetByFirst(first string, _select ...string) []model.Area {
	return areaServ.repo.GetByFirst(first, _select...)
}

func (areaServ AreaService) DeleteByFirst(first string) error {
	_, err := areaServ.repo.DeleteByFirst(first)
	return err
}
func (areaServ AreaService) GetByLevel(level uint8, _select ...string) []model.Area {
	return areaServ.repo.GetByLevel(level, _select...)
}

func (areaServ AreaService) DeleteByLevel(level uint8) error {
	_, err := areaServ.repo.DeleteByLevel(level)
	return err
}
func (areaServ AreaService) GetByPid(pid uint, _select ...string) []model.Area {
	return areaServ.repo.GetByPid(pid, _select...)
}

func (areaServ AreaService) DeleteByPid(pid uint) error {
	_, err := areaServ.repo.DeleteByPid(pid)
	return err
}

func (areaServ AreaService) GetByIDLock(ID uint64, _select ...string) (model.Area, repoComm.ReleaseLock) {
	return areaServ.repo.GetByIDLock(ID, _select...)
}

func (areaServ AreaService) ListByWhere(where repoInterface.AreaSearchWhere) []model.Area {
	return areaServ.repo.GetList(where)
}

func (areaServ AreaService) TotalCount(where repoInterface.AreaSearchWhere) int64 {
	return areaServ.repo.GetTotalCount(where)
}

func (areaServ AreaService) ListLayered() []model.Area {
	for areaUpdateIng {
		time.Sleep(1 * time.Second)
	}
	if AreaIsUpdate || len(areaLayeredList) == 0 {
		areaLayeredListMutex.Lock()
		areaUpdateIng = true
		defer func() {
			areaUpdateIng = false
			areaLayeredListMutex.Unlock()
		}()
		if len(areaLayeredList) > 0 && !AreaIsUpdate {
			return areaLayeredList
		}
		areaLayeredList = areaServ.GetListByPID(0)
		for k := range areaLayeredList {
			areaServ.RefreshChildren(&areaLayeredList[k])
		}
		AreaIsUpdate = false
	}
	return areaLayeredList
}

func (areaServ AreaService) RefreshChildren(area *model.Area) bool {
	area.Children = []model.Area{}
	area.Children = areaServ.GetListByPID(uint(area.ID))
	for k := range area.Children {
		areaServ.RefreshChildren(&area.Children[k])
	}
	return true
}

func (areaServ AreaService) CheckAreaID(proID, cityID, regionID uint) error {
	area := areaServ.GetByID(regionID)
	if area.ID == 0 {
		return sErr.ErrInvalidRegionID
	}
	if area.Pid != cityID {
		return sErr.ErrInvalidRegionID
	}
	area = areaServ.GetByID(cityID)
	if area.Pid != proID {
		return sErr.ErrInvalidProvinceID
	}
	return nil
}
