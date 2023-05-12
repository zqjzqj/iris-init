package repositories

import (
	"gorm.io/gorm"
	"iris-init/model"
	"iris-init/orm"
	"iris-init/repositories/repoComm"
	"iris-init/repositories/repoInterface"
)

type SettingsRepoGorm struct {
	repoComm.RepoGorm
}

func NewSettingsRepo() repoInterface.SettingsRepo {
	return &SettingsRepoGorm{repoComm.NewRepoGorm()}
}

//该方法需要自己去完善 GetSearchWhereTx方法内部
func (settingsRepo *SettingsRepoGorm) GetByWhere(where repoInterface.SettingsSearchWhere) model.Settings {
	settings := model.Settings{}
	_ = settingsRepo.GetSearchWhereTx(where, nil).Find(&settings)
	return settings
}

//该方法需要自己去完善 GetSearchWhereTx方法内部
func (settingsRepo *SettingsRepoGorm) GetIDByWhere(where repoInterface.SettingsSearchWhere) []uint64 {
	var id []uint64
	tx := settingsRepo.GetSearchWhereTx(where, nil)
	tx.Select("id").Model(model.Settings{}).Scan(&id)
	return id
}

func (settingsRepo *SettingsRepoGorm) Create(settings *[]model.Settings) error {
	return settingsRepo.Orm.Create(settings).Error
}

func (settingsRepo *SettingsRepoGorm) Save(settings *model.Settings, _select ...string) error {
	return repoComm.SaveModel(settingsRepo.Orm, settings, _select...)
}

func (settingsRepo *SettingsRepoGorm) SaveOmit(settings *model.Settings, _omit ...string) error {
	return repoComm.SaveModelOmit(settingsRepo.Orm, settings, _omit...)
}

//这里因为gorm的缘故 传入的settings主键必须不为空
func (settingsRepo *SettingsRepoGorm) Delete(settings model.Settings) (rowsAffected int64, err error) {
	tx := settingsRepo.Orm.Delete(settings)
	return tx.RowsAffected, tx.Error
}

//为了避免更换源之后的一些麻烦 该方法不建议在仓库结构SettingsRepoGorm以外使用
func (settingsRepo *SettingsRepoGorm) deleteByWhere(query string, args ...interface{}) (rowsAffected int64, err error) {
	tx := settingsRepo.Orm.Where(query, args...).Delete(model.Settings{})
	return tx.RowsAffected, tx.Error
}

func (settingsRepo *SettingsRepoGorm) DeleteByID(id ...uint64) (rowsAffected int64, err error) {
	if len(id) == 1 {
		return settingsRepo.deleteByWhere("id", id[0])
	}
	return settingsRepo.deleteByWhere("id in ?", id)
}

func (settingsRepo *SettingsRepoGorm) UpdateByWhere(where repoInterface.SettingsSearchWhere, data interface{}) (rowsAffected int64, err error) {
	tx := settingsRepo.GetSearchWhereTx(where, nil)
	r := tx.Updates(data)
	return r.RowsAffected, r.Error
}

func (settingsRepo *SettingsRepoGorm) GetSearchWhereTx(where repoInterface.SettingsSearchWhere, tx0 *gorm.DB) *gorm.DB {
	var tx *gorm.DB
	if tx0 == nil {
		tx = settingsRepo.Orm.Model(model.Settings{})
	} else {
		tx = tx0.Model(model.Settings{})
	}
	//需要额外调整
	if where.ID != "" {
		tx.Where("id", where.ID)
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
	if where.IDSort != "" {
		if where.IDSort == "asc" {
			tx.Order("id asc")
		} else {
			tx.Order("id desc")
		}
	}
	//需要额外调整
	if where.Key != "" {
		tx.Where("key", where.Key)
	}
	if where.KeyLike != "" {
		tx.Where("key like ?", "%"+where.KeyLike+"%")
	}
	//需要额外调整
	if where.Name != "" {
		tx.Where("name", where.Name)
	}
	if where.NameLike != "" {
		tx.Where("name like ?", "%"+where.NameLike+"%")
	}
	//需要额外调整
	if where.Desc != "" {
		tx.Where("desc", where.Desc)
	}
	if where.DescLike != "" {
		tx.Where("desc like ?", "%"+where.DescLike+"%")
	}
	//需要额外调整
	if where.Value != "" {
		tx.Where("value", where.Value)
	}
	if where.ValueLike != "" {
		tx.Where("value like ?", "%"+where.ValueLike+"%")
	}
	//需要额外调整
	if where.CreatedAt != "" {
		tx.Where("created_at", where.CreatedAt)
	}
	if where.CreatedAtLt != "" {
		tx.Where("created_at < ?", where.CreatedAtLt)
	}
	if where.CreatedAtElt != "" {
		tx.Where("created_at <= ?", where.CreatedAtElt)
	}
	if where.CreatedAtGt != "" {
		tx.Where("created_at > ?", where.CreatedAtGt)
	}
	if where.CreatedAtEgt != "" {
		tx.Where("created_at >= ?", where.CreatedAtEgt)
	}
	if where.CreatedAtSort != "" {
		if where.CreatedAtSort == "asc" {
			tx.Order("created_at asc")
		} else {
			tx.Order("created_at desc")
		}
	}
	//需要额外调整
	if where.UpdatedAt != "" {
		tx.Where("updated_at", where.UpdatedAt)
	}
	if where.UpdatedAtLt != "" {
		tx.Where("updated_at < ?", where.UpdatedAtLt)
	}
	if where.UpdatedAtElt != "" {
		tx.Where("updated_at <= ?", where.UpdatedAtElt)
	}
	if where.UpdatedAtGt != "" {
		tx.Where("updated_at > ?", where.UpdatedAtGt)
	}
	if where.UpdatedAtEgt != "" {
		tx.Where("updated_at >= ?", where.UpdatedAtEgt)
	}
	if where.UpdatedAtSort != "" {
		if where.UpdatedAtSort == "asc" {
			tx.Order("updated_at asc")
		} else {
			tx.Order("updated_at desc")
		}
	}
	return tx
}

//返回数据总数
func (settingsRepo *SettingsRepoGorm) GetTotalCount(where repoInterface.SettingsSearchWhere) int64 {
	tx := settingsRepo.GetSearchWhereTx(where, nil)
	var r int64
	tx.Count(&r)
	return r
}

func (settingsRepo *SettingsRepoGorm) GetList(where repoInterface.SettingsSearchWhere) []model.Settings {
	settings := make([]model.Settings, 0, where.SelectParams.RetSize)
	tx := settingsRepo.GetSearchWhereTx(where, nil)
	where.SelectParams.SetTxGorm(tx).Find(&settings)
	return settings
}

func (settingsRepo *SettingsRepoGorm) GetByID(id uint64, _select ...string) model.Settings {
	if id == 0 {
		return model.Settings{}
	}
	settings := model.Settings{}
	tx := settingsRepo.Orm.Where("id", id)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&settings)
	return settings
}
func (settingsRepo *SettingsRepoGorm) GetByKey(key string, _select ...string) model.Settings {
	settings := model.Settings{}
	tx := settingsRepo.Orm.
		Where("key", key)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&settings)
	return settings
}

func (settingsRepo *SettingsRepoGorm) DeleteByKey(key string) (rowsAffected int64, err error) {
	tx := settingsRepo.Orm.
		Where("key", key)
	r := tx.Delete(model.Settings{})
	return r.RowsAffected, r.Error
}

func (settingsRepo *SettingsRepoGorm) GetByIDLock(id uint64, _select ...string) (model.Settings, repoComm.ReleaseLock) {
	if id == 0 {
		panic("settingsRepo.GetByIDLock id must > 0")
	}
	if !orm.IsBeginTransaction(settingsRepo.Orm) {
		panic("settingsRepo.GetByIDLock is must beginTransaction")
	}
	settings := model.Settings{}
	tx := orm.LockForUpdate(settingsRepo.Orm.Where("id", id))
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&settings)

	//这里返回一个空的释放锁方法 因为gorm在事务提交或回滚后会自动释放
	return settings, func() {}
}
