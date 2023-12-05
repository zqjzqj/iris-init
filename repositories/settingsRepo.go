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

// 该方法需要自己去完善 GetSearchWhereTx方法内部
func (settingsRepo *SettingsRepoGorm) GetByWhere(where repoInterface.SettingsSearchWhere) model.Settings {
	settings := model.Settings{}
	_ = settingsRepo.GetSearchWhereTx(where, nil).Limit(1).Find(&settings)
	return settings
}

// 该方法需要自己去完善 GetSearchWhereTx方法内部
func (settingsRepo *SettingsRepoGorm) GetIDByWhere(where repoInterface.SettingsSearchWhere) []uint64 {
	var ID []uint64
	tx := settingsRepo.GetSearchWhereTx(where, nil)
	tx.Select("id").Model(model.Settings{}).Scan(&ID)
	return ID
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

// 这里因为gorm的缘故 传入的settings主键必须不为空
func (settingsRepo *SettingsRepoGorm) Delete(settings model.Settings) (rowsAffected int64, err error) {
	tx := settingsRepo.Orm.Delete(settings)
	return tx.RowsAffected, tx.Error
}

// 为了避免更换源之后的一些麻烦 该方法不建议在仓库结构SettingsRepoGorm以外使用
func (settingsRepo *SettingsRepoGorm) deleteByWhere(query string, args ...interface{}) (rowsAffected int64, err error) {
	tx := settingsRepo.Orm.Where(query, args...).Delete(model.Settings{})
	return tx.RowsAffected, tx.Error
}

func (settingsRepo *SettingsRepoGorm) DeleteByID(ID ...uint64) (rowsAffected int64, err error) {
	if len(ID) == 1 {
		return settingsRepo.deleteByWhere("id", ID[0])
	}
	return settingsRepo.deleteByWhere("id in ?", ID)
}

func (settingsRepo *SettingsRepoGorm) UpdateByWhere(where repoInterface.SettingsSearchWhere, data interface{}) (rowsAffected int64, err error) {
	tx := settingsRepo.GetSearchWhereTx(where, nil)
	r := tx.Updates(data)
	return r.RowsAffected, r.Error
}

func (settingsRepo *SettingsRepoGorm) DeleteByWhere(where repoInterface.SettingsSearchWhere) (rowsAffected int64, err error) {
	tx := settingsRepo.GetSearchWhereTx(where, nil)
	r := tx.Delete(model.Settings{})
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
	if where.Key != "" {
		tx.Where("key", where.Key)
	}
	if where.KeyNeq != "" {
		tx.Where("key <> ?", where.KeyNeq)
	}
	if where.KeyNull {
		tx.Where("key is null")
	}
	if where.KeyLike != "" {
		tx.Where("key like ?", "%"+where.KeyLike+"%")
	}

	if len(where.KeyIn) > 0 {
		tx.Where("key in ?", where.KeyIn)
	}

	if where.KeyNotNull {
		tx.Where("key is not null")
	}
	if where.KeySort != "" {
		if where.KeySort == "asc" {
			tx.Order("key asc")
		} else {
			tx.Order("key desc")
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
	if where.Desc != "" {
		tx.Where("desc", where.Desc)
	}
	if where.DescNeq != "" {
		tx.Where("desc <> ?", where.DescNeq)
	}
	if where.DescNull {
		tx.Where("desc is null")
	}
	if where.DescLike != "" {
		tx.Where("desc like ?", "%"+where.DescLike+"%")
	}

	if len(where.DescIn) > 0 {
		tx.Where("desc in ?", where.DescIn)
	}

	if where.DescNotNull {
		tx.Where("desc is not null")
	}
	if where.DescSort != "" {
		if where.DescSort == "asc" {
			tx.Order("desc asc")
		} else {
			tx.Order("desc desc")
		}
	}
	//需要额外调整
	if where.Value != "" {
		tx.Where("value", where.Value)
	}
	if where.ValueNeq != "" {
		tx.Where("value <> ?", where.ValueNeq)
	}
	if where.ValueNull {
		tx.Where("value is null")
	}
	if where.ValueLike != "" {
		tx.Where("value like ?", "%"+where.ValueLike+"%")
	}

	if len(where.ValueIn) > 0 {
		tx.Where("value in ?", where.ValueIn)
	}

	if where.ValueNotNull {
		tx.Where("value is not null")
	}
	if where.ValueSort != "" {
		if where.ValueSort == "asc" {
			tx.Order("value asc")
		} else {
			tx.Order("value desc")
		}
	}
	//需要额外调整
	if where.InputType != "" {
		tx.Where("input_type", where.InputType)
	}
	if where.InputTypeNeq != "" {
		tx.Where("input_type <> ?", where.InputTypeNeq)
	}
	if where.InputTypeNull {
		tx.Where("input_type is null")
	}
	if where.InputTypeLike != "" {
		tx.Where("input_type like ?", "%"+where.InputTypeLike+"%")
	}

	if len(where.InputTypeIn) > 0 {
		tx.Where("input_type in ?", where.InputTypeIn)
	}

	if where.InputTypeNotNull {
		tx.Where("input_type is not null")
	}
	if where.InputTypeSort != "" {
		if where.InputTypeSort == "asc" {
			tx.Order("input_type asc")
		} else {
			tx.Order("input_type desc")
		}
	}
	//需要额外调整
	if where.CreatedAt != "" {
		tx.Where("created_at", where.CreatedAt)
	}
	if where.CreatedAtNeq != "" {
		tx.Where("created_at <> ?", where.CreatedAtNeq)
	}
	if where.CreatedAtNull {
		tx.Where("created_at is null")
	}

	if len(where.CreatedAtIn) > 0 {
		tx.Where("created_at in ?", where.CreatedAtIn)
	}

	if where.CreatedAtNotNull {
		tx.Where("created_at is not null")
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
	if len(where.CreatedAtNotIn) > 0 {
		tx.Where("created_at not in ?", where.CreatedAtNotIn)
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
	if where.UpdatedAtNeq != "" {
		tx.Where("updated_at <> ?", where.UpdatedAtNeq)
	}
	if where.UpdatedAtNull {
		tx.Where("updated_at is null")
	}

	if len(where.UpdatedAtIn) > 0 {
		tx.Where("updated_at in ?", where.UpdatedAtIn)
	}

	if where.UpdatedAtNotNull {
		tx.Where("updated_at is not null")
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
	if len(where.UpdatedAtNotIn) > 0 {
		tx.Where("updated_at not in ?", where.UpdatedAtNotIn)
	}
	if where.UpdatedAtSort != "" {
		if where.UpdatedAtSort == "asc" {
			tx.Order("updated_at asc")
		} else {
			tx.Order("updated_at desc")
		}
	}
	where.SelectParams.SetTxGorm(tx)
	return tx
}

// 返回数据总数
func (settingsRepo *SettingsRepoGorm) GetTotalCount(where repoInterface.SettingsSearchWhere) int64 {
	tx := settingsRepo.GetSearchWhereTx(where, nil)
	var r int64
	tx.Count(&r)
	return r
}

func (settingsRepo *SettingsRepoGorm) GetList(where repoInterface.SettingsSearchWhere) []model.Settings {
	settings := make([]model.Settings, 0, where.SelectParams.RetSize)
	tx := settingsRepo.GetSearchWhereTx(where, nil)
	tx.Find(&settings)
	return settings
}

func (settingsRepo *SettingsRepoGorm) GetByID(ID uint64, _select ...string) model.Settings {
	settings := model.Settings{}
	tx := settingsRepo.Orm.Where("id", ID)
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

func (settingsRepo *SettingsRepoGorm) GetByIDLock(ID uint64, _select ...string) (model.Settings, repoComm.ReleaseLock) {
	if !orm.IsBeginTransaction(settingsRepo.Orm) {
		panic("settingsRepo.GetByIDLock is must beginTransaction")
	}
	settings := model.Settings{}
	tx := orm.LockForUpdate(settingsRepo.Orm.Where("id", ID))
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&settings)

	//这里返回一个空的释放锁方法 因为gorm在事务提交或回滚后会自动释放
	return settings, func() {}
}

func (settingsRepo *SettingsRepoGorm) ScanByWhere(where repoInterface.SettingsSearchWhere, dest any) error {
	return settingsRepo.GetSearchWhereTx(where, nil).Find(dest).Error
}

func (settingsRepo *SettingsRepoGorm) ScanByOrWhere(dest any, where ...repoInterface.SettingsSearchWhere) error {
	tx := settingsRepo.Orm.Model(model.Settings{})
	for _, v := range where {
		tx.Or(settingsRepo.GetSearchWhereTx(v, nil))
	}
	return tx.Find(dest).Error
}
