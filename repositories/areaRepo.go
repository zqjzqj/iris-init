package repositories

import (
	"gorm.io/gorm"
	"iris-init/model"
	"iris-init/orm"
	"iris-init/repositories/repoComm"
	"iris-init/repositories/repoInterface"
)

type AreaRepoGorm struct {
	repoComm.RepoGorm
}

func NewAreaRepo() repoInterface.AreaRepo {
	return &AreaRepoGorm{repoComm.NewRepoGorm()}
}

// 该方法需要自己去完善 GetSearchWhereTx方法内部
func (areaRepo *AreaRepoGorm) GetByWhere(where repoInterface.AreaSearchWhere) model.Area {
	area := model.Area{}
	_ = areaRepo.GetSearchWhereTx(where, nil).Limit(1).Find(&area)
	return area
}

// 该方法需要自己去完善 GetSearchWhereTx方法内部
func (areaRepo *AreaRepoGorm) GetIDByWhere(where repoInterface.AreaSearchWhere) []uint64 {
	var ID []uint64
	tx := areaRepo.GetSearchWhereTx(where, nil)
	tx.Select("id").Model(model.Area{}).Scan(&ID)
	return ID
}

func (areaRepo *AreaRepoGorm) Create(area *[]model.Area) error {
	return areaRepo.Orm.Create(area).Error
}

func (areaRepo *AreaRepoGorm) Save(area *model.Area, _select ...string) error {
	return repoComm.SaveModel(areaRepo.Orm, area, _select...)
}

func (areaRepo *AreaRepoGorm) SaveOmit(area *model.Area, _omit ...string) error {
	return repoComm.SaveModelOmit(areaRepo.Orm, area, _omit...)
}

// 这里因为gorm的缘故 传入的area主键必须不为空
func (areaRepo *AreaRepoGorm) Delete(area model.Area) (rowsAffected int64, err error) {
	tx := areaRepo.Orm.Delete(area)
	return tx.RowsAffected, tx.Error
}

// 为了避免更换源之后的一些麻烦 该方法不建议在仓库结构AreaRepoGorm以外使用
func (areaRepo *AreaRepoGorm) deleteByWhere(query string, args ...interface{}) (rowsAffected int64, err error) {
	tx := areaRepo.Orm.Where(query, args...).Delete(&model.Area{})
	return tx.RowsAffected, tx.Error
}

func (areaRepo *AreaRepoGorm) DeleteByID(ID ...uint64) (rowsAffected int64, err error) {
	if len(ID) == 1 {
		return areaRepo.deleteByWhere("id", ID[0])
	}
	return areaRepo.deleteByWhere("id in ?", ID)
}

func (areaRepo *AreaRepoGorm) UpdateByWhere(where repoInterface.AreaSearchWhere, data interface{}) (rowsAffected int64, err error) {
	tx := areaRepo.GetSearchWhereTx(where, nil)
	r := tx.Updates(data)
	return r.RowsAffected, r.Error
}

func (areaRepo *AreaRepoGorm) DeleteByWhere(where repoInterface.AreaSearchWhere) (rowsAffected int64, err error) {
	tx := areaRepo.GetSearchWhereTx(where, nil)
	r := tx.Delete(&model.Area{})
	return r.RowsAffected, r.Error
}

func (areaRepo *AreaRepoGorm) GetSearchWhereTx(where repoInterface.AreaSearchWhere, tx0 *gorm.DB) *gorm.DB {
	var tx *gorm.DB
	if tx0 == nil {
		tx = areaRepo.Orm.Model(model.Area{})
	} else {
		tx = tx0.Model(model.Area{})
	}
	//需要额外调整
	if where.ID != "" {
		tx.Where("id", where.ID)
	}
	if where.IDNeq != "" {
		tx.Where("id <> ?", where.IDNeq)
	}
	if where.IDNull {
		tx.Where("id is null")
	}

	if len(where.IDIn) > 0 {
		tx.Where("id in ?", where.IDIn)
	}

	if where.IDNotNull {
		tx.Where("id is not null")
	}
	if where.IDLt != "" {
		tx.Where("id < ?", where.IDLt)
	}
	if where.IDElt != "" {
		tx.Where("id <= ?", where.IDElt)
	}
	if where.IDGt != "" {
		tx.Where("id > ?", where.IDGt)
	}
	if where.IDEgt != "" {
		tx.Where("id >= ?", where.IDEgt)
	}
	if len(where.IDNotIn) > 0 {
		tx.Where("id not in ?", where.IDNotIn)
	}
	if where.IDSort != "" {
		if where.IDSort == "asc" {
			tx.Order("id asc")
		} else {
			tx.Order("id desc")
		}
	}
	//需要额外调整
	if where.Pid != "" {
		tx.Where("pid", where.Pid)
	}
	if where.PidNeq != "" {
		tx.Where("pid <> ?", where.PidNeq)
	}
	if where.PidNull {
		tx.Where("pid is null")
	}

	if len(where.PidIn) > 0 {
		tx.Where("pid in ?", where.PidIn)
	}

	if where.PidNotNull {
		tx.Where("pid is not null")
	}
	if where.PidLt != "" {
		tx.Where("pid < ?", where.PidLt)
	}
	if where.PidElt != "" {
		tx.Where("pid <= ?", where.PidElt)
	}
	if where.PidGt != "" {
		tx.Where("pid > ?", where.PidGt)
	}
	if where.PidEgt != "" {
		tx.Where("pid >= ?", where.PidEgt)
	}
	if len(where.PidNotIn) > 0 {
		tx.Where("pid not in ?", where.PidNotIn)
	}
	if where.PidSort != "" {
		if where.PidSort == "asc" {
			tx.Order("pid asc")
		} else {
			tx.Order("pid desc")
		}
	}
	//需要额外调整
	if where.ShortName != "" {
		tx.Where("short_name", where.ShortName)
	}
	if where.ShortNameNeq != "" {
		tx.Where("short_name <> ?", where.ShortNameNeq)
	}
	if where.ShortNameNull {
		tx.Where("short_name is null")
	}
	if where.ShortNameLike != "" {
		tx.Where("short_name like ?", "%"+where.ShortNameLike+"%")
	}

	if len(where.ShortNameIn) > 0 {
		tx.Where("short_name in ?", where.ShortNameIn)
	}

	if where.ShortNameNotNull {
		tx.Where("short_name is not null")
	}
	if where.ShortNameSort != "" {
		if where.ShortNameSort == "asc" {
			tx.Order("short_name asc")
		} else {
			tx.Order("short_name desc")
		}
	}
	//需要额外调整
	if where.Name != "" {
		tx.Where("name", where.Name)
	}
	if where.NameNeq != "" {
		tx.Where("name <> ?", where.NameNeq)
	}
	if where.NameNull {
		tx.Where("name is null")
	}
	if where.NameLike != "" {
		tx.Where("name like ?", "%"+where.NameLike+"%")
	}

	if len(where.NameIn) > 0 {
		tx.Where("name in ?", where.NameIn)
	}

	if where.NameNotNull {
		tx.Where("name is not null")
	}
	if where.NameSort != "" {
		if where.NameSort == "asc" {
			tx.Order("name asc")
		} else {
			tx.Order("name desc")
		}
	}
	//需要额外调整
	if where.MergerName != "" {
		tx.Where("merger_name", where.MergerName)
	}
	if where.MergerNameNeq != "" {
		tx.Where("merger_name <> ?", where.MergerNameNeq)
	}
	if where.MergerNameNull {
		tx.Where("merger_name is null")
	}
	if where.MergerNameLike != "" {
		tx.Where("merger_name like ?", "%"+where.MergerNameLike+"%")
	}

	if len(where.MergerNameIn) > 0 {
		tx.Where("merger_name in ?", where.MergerNameIn)
	}

	if where.MergerNameNotNull {
		tx.Where("merger_name is not null")
	}
	if where.MergerNameSort != "" {
		if where.MergerNameSort == "asc" {
			tx.Order("merger_name asc")
		} else {
			tx.Order("merger_name desc")
		}
	}
	//需要额外调整
	if where.Level != "" {
		tx.Where("level", where.Level)
	}
	if where.LevelNeq != "" {
		tx.Where("level <> ?", where.LevelNeq)
	}
	if where.LevelNull {
		tx.Where("level is null")
	}

	if len(where.LevelIn) > 0 {
		tx.Where("level in ?", where.LevelIn)
	}

	if where.LevelNotNull {
		tx.Where("level is not null")
	}
	if where.LevelLt != "" {
		tx.Where("level < ?", where.LevelLt)
	}
	if where.LevelElt != "" {
		tx.Where("level <= ?", where.LevelElt)
	}
	if where.LevelGt != "" {
		tx.Where("level > ?", where.LevelGt)
	}
	if where.LevelEgt != "" {
		tx.Where("level >= ?", where.LevelEgt)
	}
	if len(where.LevelNotIn) > 0 {
		tx.Where("level not in ?", where.LevelNotIn)
	}
	if where.LevelSort != "" {
		if where.LevelSort == "asc" {
			tx.Order("level asc")
		} else {
			tx.Order("level desc")
		}
	}
	//需要额外调整
	if where.PinYin != "" {
		tx.Where("pin_yin", where.PinYin)
	}
	if where.PinYinNeq != "" {
		tx.Where("pin_yin <> ?", where.PinYinNeq)
	}
	if where.PinYinNull {
		tx.Where("pin_yin is null")
	}
	if where.PinYinLike != "" {
		tx.Where("pin_yin like ?", "%"+where.PinYinLike+"%")
	}

	if len(where.PinYinIn) > 0 {
		tx.Where("pin_yin in ?", where.PinYinIn)
	}

	if where.PinYinNotNull {
		tx.Where("pin_yin is not null")
	}
	if where.PinYinSort != "" {
		if where.PinYinSort == "asc" {
			tx.Order("pin_yin asc")
		} else {
			tx.Order("pin_yin desc")
		}
	}
	//需要额外调整
	if where.Code != "" {
		tx.Where("code", where.Code)
	}
	if where.CodeNeq != "" {
		tx.Where("code <> ?", where.CodeNeq)
	}
	if where.CodeNull {
		tx.Where("code is null")
	}
	if where.CodeLike != "" {
		tx.Where("code like ?", "%"+where.CodeLike+"%")
	}

	if len(where.CodeIn) > 0 {
		tx.Where("code in ?", where.CodeIn)
	}

	if where.CodeNotNull {
		tx.Where("code is not null")
	}
	if where.CodeSort != "" {
		if where.CodeSort == "asc" {
			tx.Order("code asc")
		} else {
			tx.Order("code desc")
		}
	}
	//需要额外调整
	if where.ZipCode != "" {
		tx.Where("zip_code", where.ZipCode)
	}
	if where.ZipCodeNeq != "" {
		tx.Where("zip_code <> ?", where.ZipCodeNeq)
	}
	if where.ZipCodeNull {
		tx.Where("zip_code is null")
	}
	if where.ZipCodeLike != "" {
		tx.Where("zip_code like ?", "%"+where.ZipCodeLike+"%")
	}

	if len(where.ZipCodeIn) > 0 {
		tx.Where("zip_code in ?", where.ZipCodeIn)
	}

	if where.ZipCodeNotNull {
		tx.Where("zip_code is not null")
	}
	if where.ZipCodeSort != "" {
		if where.ZipCodeSort == "asc" {
			tx.Order("zip_code asc")
		} else {
			tx.Order("zip_code desc")
		}
	}
	//需要额外调整
	if where.First != "" {
		tx.Where("first", where.First)
	}
	if where.FirstNeq != "" {
		tx.Where("first <> ?", where.FirstNeq)
	}
	if where.FirstNull {
		tx.Where("first is null")
	}
	if where.FirstLike != "" {
		tx.Where("first like ?", "%"+where.FirstLike+"%")
	}

	if len(where.FirstIn) > 0 {
		tx.Where("first in ?", where.FirstIn)
	}

	if where.FirstNotNull {
		tx.Where("first is not null")
	}
	if where.FirstSort != "" {
		if where.FirstSort == "asc" {
			tx.Order("first asc")
		} else {
			tx.Order("first desc")
		}
	}
	//需要额外调整
	if where.Lng != "" {
		tx.Where("lng", where.Lng)
	}
	if where.LngNeq != "" {
		tx.Where("lng <> ?", where.LngNeq)
	}
	if where.LngNull {
		tx.Where("lng is null")
	}
	if where.LngLike != "" {
		tx.Where("lng like ?", "%"+where.LngLike+"%")
	}

	if len(where.LngIn) > 0 {
		tx.Where("lng in ?", where.LngIn)
	}

	if where.LngNotNull {
		tx.Where("lng is not null")
	}
	if where.LngSort != "" {
		if where.LngSort == "asc" {
			tx.Order("lng asc")
		} else {
			tx.Order("lng desc")
		}
	}
	//需要额外调整
	if where.Lat != "" {
		tx.Where("lat", where.Lat)
	}
	if where.LatNeq != "" {
		tx.Where("lat <> ?", where.LatNeq)
	}
	if where.LatNull {
		tx.Where("lat is null")
	}
	if where.LatLike != "" {
		tx.Where("lat like ?", "%"+where.LatLike+"%")
	}

	if len(where.LatIn) > 0 {
		tx.Where("lat in ?", where.LatIn)
	}

	if where.LatNotNull {
		tx.Where("lat is not null")
	}
	if where.LatSort != "" {
		if where.LatSort == "asc" {
			tx.Order("lat asc")
		} else {
			tx.Order("lat desc")
		}
	}
	//需要额外调整
	if where.Children != "" {
		tx.Where("children", where.Children)
	}
	if where.ChildrenNeq != "" {
		tx.Where("children <> ?", where.ChildrenNeq)
	}
	if where.ChildrenNull {
		tx.Where("children is null")
	}

	if len(where.ChildrenIn) > 0 {
		tx.Where("children in ?", where.ChildrenIn)
	}

	if where.ChildrenNotNull {
		tx.Where("children is not null")
	}
	if where.ChildrenSort != "" {
		if where.ChildrenSort == "asc" {
			tx.Order("children asc")
		} else {
			tx.Order("children desc")
		}
	}
	where.SelectParams.SetTxGorm(tx)
	return tx
}

// 返回数据总数
func (areaRepo *AreaRepoGorm) GetTotalCount(where repoInterface.AreaSearchWhere) int64 {
	tx := areaRepo.GetSearchWhereTx(where, nil)
	var r int64
	tx.Count(&r)
	return r
}

func (areaRepo *AreaRepoGorm) GetList(where repoInterface.AreaSearchWhere) []model.Area {
	area := make([]model.Area, 0, where.SelectParams.RetSize)
	tx := areaRepo.GetSearchWhereTx(where, nil)
	tx.Find(&area)
	return area
}

func (areaRepo *AreaRepoGorm) GetByID(ID uint64, _select ...string) model.Area {
	area := model.Area{}
	tx := areaRepo.Orm.Where("id", ID)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&area)
	return area
}
func (areaRepo *AreaRepoGorm) GetByFirst(first string, _select ...string) []model.Area {
	area := make([]model.Area, 0)
	tx := areaRepo.Orm.
		Where("first", first)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&area)
	return area
}

func (areaRepo *AreaRepoGorm) DeleteByFirst(first string) (rowsAffected int64, err error) {
	tx := areaRepo.Orm.
		Where("first", first)
	r := tx.Delete(&model.Area{})
	return r.RowsAffected, r.Error
}
func (areaRepo *AreaRepoGorm) GetByLevel(level uint8, _select ...string) []model.Area {
	area := make([]model.Area, 0)
	tx := areaRepo.Orm.
		Where("level", level)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&area)
	return area
}

func (areaRepo *AreaRepoGorm) DeleteByLevel(level uint8) (rowsAffected int64, err error) {
	tx := areaRepo.Orm.
		Where("level", level)
	r := tx.Delete(&model.Area{})
	return r.RowsAffected, r.Error
}
func (areaRepo *AreaRepoGorm) GetByPid(pid uint, _select ...string) []model.Area {
	area := make([]model.Area, 0)
	tx := areaRepo.Orm.
		Where("pid", pid)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&area)
	return area
}

func (areaRepo *AreaRepoGorm) DeleteByPid(pid uint) (rowsAffected int64, err error) {
	tx := areaRepo.Orm.
		Where("pid", pid)
	r := tx.Delete(&model.Area{})
	return r.RowsAffected, r.Error
}

func (areaRepo *AreaRepoGorm) GetByIDLock(ID uint64, _select ...string) (model.Area, repoComm.ReleaseLock) {
	if !orm.IsBeginTransaction(areaRepo.Orm) {
		panic("areaRepo.GetByIDLock is must beginTransaction")
	}
	area := model.Area{}
	tx := orm.LockForUpdate(areaRepo.Orm.Where("id", ID))
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&area)

	//这里返回一个空的释放锁方法 因为gorm在事务提交或回滚后会自动释放
	return area, func() {}
}

func (areaRepo *AreaRepoGorm) ScanByWhere(where repoInterface.AreaSearchWhere, dest any) error {
	return areaRepo.GetSearchWhereTx(where, nil).Find(dest).Error
}

func (areaRepo *AreaRepoGorm) ScanByOrWhere(dest any, where ...repoInterface.AreaSearchWhere) error {
	tx := areaRepo.Orm.Model(model.Area{})
	for _, v := range where {
		tx.Or(areaRepo.GetSearchWhereTx(v, nil))
		v.SelectParams.SetTxGorm(tx)
	}
	return tx.Find(dest).Error
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
