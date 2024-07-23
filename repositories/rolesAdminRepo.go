package repositories

import (
	"gorm.io/gorm"
	"iris-init/global"
	"iris-init/model"
	"iris-init/orm"
	"iris-init/repositories/repoComm"
	"iris-init/repositories/repoInterface"
	"strconv"
)

type RolesAdminRepoGorm struct {
	repoComm.RepoGorm
}

func NewRolesAdminRepo() repoInterface.RolesAdminRepo {
	return &RolesAdminRepoGorm{repoComm.NewRepoGorm()}
}

// 该方法需要自己去完善 GetSearchWhereTx方法内部
func (rolesAdminRepo *RolesAdminRepoGorm) GetByWhere(where repoInterface.RolesAdminSearchWhere) model.RolesAdmin {
	rolesAdmin := model.RolesAdmin{}
	_ = rolesAdminRepo.GetSearchWhereTx(where, nil).Limit(1).Find(&rolesAdmin)
	return rolesAdmin
}

// 该方法需要自己去完善 GetSearchWhereTx方法内部
func (rolesAdminRepo *RolesAdminRepoGorm) GetIDByWhere(where repoInterface.RolesAdminSearchWhere) [][]uint8 {
	var ID [][]uint8
	tx := rolesAdminRepo.GetSearchWhereTx(where, nil)
	tx.Select("id").Model(model.RolesAdmin{}).Scan(&ID)
	return ID
}

func (rolesAdminRepo *RolesAdminRepoGorm) Create(rolesAdmin *[]model.RolesAdmin) error {
	return rolesAdminRepo.Orm.Create(rolesAdmin).Error
}

func (rolesAdminRepo *RolesAdminRepoGorm) Save(rolesAdmin *model.RolesAdmin, _select ...string) error {
	return repoComm.SaveModel(rolesAdminRepo.Orm, rolesAdmin, _select...)
}

func (rolesAdminRepo *RolesAdminRepoGorm) SaveOmit(rolesAdmin *model.RolesAdmin, _omit ...string) error {
	return repoComm.SaveModelOmit(rolesAdminRepo.Orm, rolesAdmin, _omit...)
}

// 这里因为gorm的缘故 传入的rolesAdmin主键必须不为空
func (rolesAdminRepo *RolesAdminRepoGorm) Delete(rolesAdmin model.RolesAdmin) (rowsAffected int64, err error) {
	tx := rolesAdminRepo.Orm.Delete(rolesAdmin)
	return tx.RowsAffected, tx.Error
}

// 为了避免更换源之后的一些麻烦 该方法不建议在仓库结构RolesAdminRepoGorm以外使用
func (rolesAdminRepo *RolesAdminRepoGorm) deleteByWhere(query string, args ...interface{}) (rowsAffected int64, err error) {
	tx := rolesAdminRepo.Orm.Where(query, args...).Delete(&model.RolesAdmin{})
	return tx.RowsAffected, tx.Error
}

func (rolesAdminRepo *RolesAdminRepoGorm) DeleteByID(ID ...[]uint8) (rowsAffected int64, err error) {
	if len(ID) == 1 {
		return rolesAdminRepo.deleteByWhere("id", ID[0])
	}
	return rolesAdminRepo.deleteByWhere("id in ?", ID)
}

func (rolesAdminRepo *RolesAdminRepoGorm) UpdateByWhere(where repoInterface.RolesAdminSearchWhere, data interface{}) (rowsAffected int64, err error) {
	tx := rolesAdminRepo.GetSearchWhereTx(where, nil)
	r := tx.Updates(data)
	return r.RowsAffected, r.Error
}

func (rolesAdminRepo *RolesAdminRepoGorm) DeleteByWhere(where repoInterface.RolesAdminSearchWhere) (rowsAffected int64, err error) {
	tx := rolesAdminRepo.GetSearchWhereTx(where, nil)
	r := tx.Delete(&model.RolesAdmin{})
	return r.RowsAffected, r.Error
}

func (rolesAdminRepo *RolesAdminRepoGorm) GetSearchWhereTx(where repoInterface.RolesAdminSearchWhere, tx0 *gorm.DB) *gorm.DB {
	var tx *gorm.DB
	if tx0 == nil {
		tx = rolesAdminRepo.Orm.Model(model.RolesAdmin{})
	} else {
		tx = tx0.Model(model.RolesAdmin{})
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
	if where.RoleID != "" {
		tx.Where("role_id", where.RoleID)
	}
	if where.RoleIDNeq != "" {
		tx.Where("role_id <> ?", where.RoleIDNeq)
	}
	if where.RoleIDNull {
		tx.Where("role_id is null")
	}

	if len(where.RoleIDIn) > 0 {
		tx.Where("role_id in ?", where.RoleIDIn)
	}

	if where.RoleIDNotNull {
		tx.Where("role_id is not null")
	}
	if where.RoleIDLt != "" {
		tx.Where("role_id < ?", where.RoleIDLt)
	}
	if where.RoleIDElt != "" {
		tx.Where("role_id <= ?", where.RoleIDElt)
	}
	if where.RoleIDGt != "" {
		tx.Where("role_id > ?", where.RoleIDGt)
	}
	if where.RoleIDEgt != "" {
		tx.Where("role_id >= ?", where.RoleIDEgt)
	}
	if len(where.RoleIDNotIn) > 0 {
		tx.Where("role_id not in ?", where.RoleIDNotIn)
	}
	if where.RoleIDSort != "" {
		if where.RoleIDSort == "asc" {
			tx.Order("role_id asc")
		} else {
			tx.Order("role_id desc")
		}
	}
	//需要额外调整
	if where.AdminID != "" {
		tx.Where("admin_id", where.AdminID)
	}
	if where.AdminIDNeq != "" {
		tx.Where("admin_id <> ?", where.AdminIDNeq)
	}
	if where.AdminIDNull {
		tx.Where("admin_id is null")
	}

	if len(where.AdminIDIn) > 0 {
		tx.Where("admin_id in ?", where.AdminIDIn)
	}

	if where.AdminIDNotNull {
		tx.Where("admin_id is not null")
	}
	if where.AdminIDLt != "" {
		tx.Where("admin_id < ?", where.AdminIDLt)
	}
	if where.AdminIDElt != "" {
		tx.Where("admin_id <= ?", where.AdminIDElt)
	}
	if where.AdminIDGt != "" {
		tx.Where("admin_id > ?", where.AdminIDGt)
	}
	if where.AdminIDEgt != "" {
		tx.Where("admin_id >= ?", where.AdminIDEgt)
	}
	if len(where.AdminIDNotIn) > 0 {
		tx.Where("admin_id not in ?", where.AdminIDNotIn)
	}
	if where.AdminIDSort != "" {
		if where.AdminIDSort == "asc" {
			tx.Order("admin_id asc")
		} else {
			tx.Order("admin_id desc")
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
func (rolesAdminRepo *RolesAdminRepoGorm) GetTotalCount(where repoInterface.RolesAdminSearchWhere) int64 {
	tx := rolesAdminRepo.GetSearchWhereTx(where, nil)
	var r int64
	tx.Count(&r)
	return r
}

func (rolesAdminRepo *RolesAdminRepoGorm) GetList(where repoInterface.RolesAdminSearchWhere) []model.RolesAdmin {
	rolesAdmin := make([]model.RolesAdmin, 0, where.SelectParams.RetSize)
	tx := rolesAdminRepo.GetSearchWhereTx(where, nil)
	tx.Find(&rolesAdmin)
	return rolesAdmin
}

func (rolesAdminRepo *RolesAdminRepoGorm) GetByID(ID []uint8, _select ...string) model.RolesAdmin {
	rolesAdmin := model.RolesAdmin{}
	tx := rolesAdminRepo.Orm.Where("id", ID)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&rolesAdmin)
	return rolesAdmin
}
func (rolesAdminRepo *RolesAdminRepoGorm) GetByAdminID(adminID uint64, _select ...string) []model.RolesAdmin {
	rolesAdmin := make([]model.RolesAdmin, 0)
	tx := rolesAdminRepo.Orm.
		Where("admin_id", adminID)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&rolesAdmin)
	return rolesAdmin
}

func (rolesAdminRepo *RolesAdminRepoGorm) DeleteByAdminID(adminID uint64) (rowsAffected int64, err error) {
	tx := rolesAdminRepo.Orm.
		Where("admin_id", adminID)
	r := tx.Delete(&model.RolesAdmin{})
	return r.RowsAffected, r.Error
}
func (rolesAdminRepo *RolesAdminRepoGorm) GetByRoleID(roleID uint64, _select ...string) []model.RolesAdmin {
	rolesAdmin := make([]model.RolesAdmin, 0)
	tx := rolesAdminRepo.Orm.
		Where("role_id", roleID)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&rolesAdmin)
	return rolesAdmin
}

func (rolesAdminRepo *RolesAdminRepoGorm) DeleteByRoleID(roleID uint64) (rowsAffected int64, err error) {
	tx := rolesAdminRepo.Orm.
		Where("role_id", roleID)
	r := tx.Delete(&model.RolesAdmin{})
	return r.RowsAffected, r.Error
}

func (rolesAdminRepo *RolesAdminRepoGorm) GetByIDLock(ID []uint8, _select ...string) (model.RolesAdmin, repoComm.ReleaseLock) {
	if !orm.IsBeginTransaction(rolesAdminRepo.Orm) {
		panic("rolesAdminRepo.GetByIDLock is must beginTransaction")
	}
	rolesAdmin := model.RolesAdmin{}
	tx := orm.LockForUpdate(rolesAdminRepo.Orm.Where("id", ID))
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&rolesAdmin)

	//这里返回一个空的释放锁方法 因为gorm在事务提交或回滚后会自动释放
	return rolesAdmin, func() {}
}

func (rolesAdminRepo *RolesAdminRepoGorm) ScanByWhere(where repoInterface.RolesAdminSearchWhere, dest any) error {
	return rolesAdminRepo.GetSearchWhereTx(where, nil).Find(dest).Error
}

func (rolesAdminRepo *RolesAdminRepoGorm) ScanByOrWhere(dest any, where ...repoInterface.RolesAdminSearchWhere) error {
	tx := rolesAdminRepo.Orm.Model(model.RolesAdmin{})
	for _, v := range where {
		tx.Or(rolesAdminRepo.GetSearchWhereTx(v, tx))
	}
	return tx.Find(dest).Error
}

func (rolesAdminRepo *RolesAdminRepoGorm) SaveByAdm(adm model.Admin) error {
	if adm.ID == 0 {
		panic("SaveByAdm adm.ID is 0")
	}
	if adm.RolesID == "" || adm.RolesID == model.RoleAdmin {
		return rolesAdminRepo.Orm.Where("admin_id", adm.ID).Delete(&model.RolesAdmin{}).Error
	}
	return rolesAdminRepo.Orm.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("admin_id", adm.ID).Delete(&model.RolesAdmin{}).Error
		if err != nil {
			return err
		}
		RolesID := adm.RefreshRoleIDSlices()
		RolesID = global.RemoveDuplicateElement(RolesID)
		rAdm := make([]model.RolesAdmin, 0, len(RolesID))
		for _, v := range RolesID {
			i, _err := strconv.ParseUint(v, 10, 64)
			if _err != nil {
				continue
			}
			rAdm = append(rAdm, model.RolesAdmin{
				RoleID:  i,
				AdminID: adm.ID,
			})
		}
		return tx.Create(&rAdm).Error
	})
}
