package repositories

import (
	"gorm.io/gorm"
	"iris-init/model"
	"iris-init/orm"
	"iris-init/repositories/repoComm"
	"iris-init/repositories/repoInterface"
)

type RolesRepoGorm struct {
	repoComm.RepoGorm
}

func NewRolesRepo() repoInterface.RolesRepo {
	return &RolesRepoGorm{repoComm.NewRepoGorm()}
}

// 该方法需要自己去完善 GetSearchWhereTx方法内部
func (rolesRepo *RolesRepoGorm) GetByWhere(where repoInterface.RolesSearchWhere) model.Roles {
	roles := model.Roles{}
	_ = rolesRepo.GetSearchWhereTx(where, nil).Limit(1).Find(&roles)
	return roles
}

// 该方法需要自己去完善 GetSearchWhereTx方法内部
func (rolesRepo *RolesRepoGorm) GetIDByWhere(where repoInterface.RolesSearchWhere) []uint64 {
	var ID []uint64
	tx := rolesRepo.GetSearchWhereTx(where, nil)
	tx.Select("id").Model(model.Roles{}).Scan(&ID)
	return ID
}

func (rolesRepo *RolesRepoGorm) Create(roles *[]model.Roles) error {
	return rolesRepo.Orm.Create(roles).Error
}

func (rolesRepo *RolesRepoGorm) Save(roles *model.Roles, _select ...string) error {
	return repoComm.SaveModel(rolesRepo.Orm, roles, _select...)
}

func (rolesRepo *RolesRepoGorm) SaveOmit(roles *model.Roles, _omit ...string) error {
	return repoComm.SaveModelOmit(rolesRepo.Orm, roles, _omit...)
}

// 这里因为gorm的缘故 传入的roles主键必须不为空
func (rolesRepo *RolesRepoGorm) Delete(roles model.Roles) (rowsAffected int64, err error) {
	tx := rolesRepo.Orm.Delete(roles)
	return tx.RowsAffected, tx.Error
}

// 为了避免更换源之后的一些麻烦 该方法不建议在仓库结构RolesRepoGorm以外使用
func (rolesRepo *RolesRepoGorm) deleteByWhere(query string, args ...interface{}) (rowsAffected int64, err error) {
	tx := rolesRepo.Orm.Where(query, args...).Delete(model.Roles{})
	return tx.RowsAffected, tx.Error
}

func (rolesRepo *RolesRepoGorm) DeleteByID(ID ...uint64) (rowsAffected int64, err error) {
	if len(ID) == 1 {
		return rolesRepo.deleteByWhere("id", ID[0])
	}
	return rolesRepo.deleteByWhere("id in ?", ID)
}

func (rolesRepo *RolesRepoGorm) UpdateByWhere(where repoInterface.RolesSearchWhere, data interface{}) (rowsAffected int64, err error) {
	tx := rolesRepo.GetSearchWhereTx(where, nil)
	r := tx.Updates(data)
	return r.RowsAffected, r.Error
}

func (rolesRepo *RolesRepoGorm) DeleteByWhere(where repoInterface.RolesSearchWhere) (rowsAffected int64, err error) {
	tx := rolesRepo.GetSearchWhereTx(where, nil)
	r := tx.Delete(model.Roles{})
	return r.RowsAffected, r.Error
}

func (rolesRepo *RolesRepoGorm) GetSearchWhereTx(where repoInterface.RolesSearchWhere, tx0 *gorm.DB) *gorm.DB {
	var tx *gorm.DB
	if tx0 == nil {
		tx = rolesRepo.Orm.Model(model.Roles{})
	} else {
		tx = tx0.Model(model.Roles{})
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
	if where.NameLt != "" {
		tx.Where("name < ?", where.NameLt)
	}
	if where.NameElt != "" {
		tx.Where("name <= ?", where.NameElt)
	}
	if where.NameGt != "" {
		tx.Where("name > ?", where.NameGt)
	}
	if where.NameEgt != "" {
		tx.Where("name >= ?", where.NameEgt)
	}
	if len(where.NameNotIn) > 0 {
		tx.Where("name not in ?", where.NameNotIn)
	}
	if where.NameSort != "" {
		if where.NameSort == "asc" {
			tx.Order("name asc")
		} else {
			tx.Order("name desc")
		}
	}
	//需要额外调整
	if where.Remark != "" {
		tx.Where("remark", where.Remark)
	}
	if where.RemarkNeq != "" {
		tx.Where("remark <> ?", where.RemarkNeq)
	}
	if where.RemarkNull {
		tx.Where("remark is null")
	}
	if where.RemarkLike != "" {
		tx.Where("remark like ?", "%"+where.RemarkLike+"%")
	}

	if len(where.RemarkIn) > 0 {
		tx.Where("remark in ?", where.RemarkIn)
	}

	if where.RemarkNotNull {
		tx.Where("remark is not null")
	}
	if where.RemarkLt != "" {
		tx.Where("remark < ?", where.RemarkLt)
	}
	if where.RemarkElt != "" {
		tx.Where("remark <= ?", where.RemarkElt)
	}
	if where.RemarkGt != "" {
		tx.Where("remark > ?", where.RemarkGt)
	}
	if where.RemarkEgt != "" {
		tx.Where("remark >= ?", where.RemarkEgt)
	}
	if len(where.RemarkNotIn) > 0 {
		tx.Where("remark not in ?", where.RemarkNotIn)
	}
	if where.RemarkSort != "" {
		if where.RemarkSort == "asc" {
			tx.Order("remark asc")
		} else {
			tx.Order("remark desc")
		}
	}
	//需要额外调整
	if where.PermIdents != "" {
		tx.Where("perm_idents", where.PermIdents)
	}
	if where.PermIdentsNeq != "" {
		tx.Where("perm_idents <> ?", where.PermIdentsNeq)
	}
	if where.PermIdentsNull {
		tx.Where("perm_idents is null")
	}

	if len(where.PermIdentsIn) > 0 {
		tx.Where("perm_idents in ?", where.PermIdentsIn)
	}

	if where.PermIdentsNotNull {
		tx.Where("perm_idents is not null")
	}
	if where.PermIdentsLt != "" {
		tx.Where("perm_idents < ?", where.PermIdentsLt)
	}
	if where.PermIdentsElt != "" {
		tx.Where("perm_idents <= ?", where.PermIdentsElt)
	}
	if where.PermIdentsGt != "" {
		tx.Where("perm_idents > ?", where.PermIdentsGt)
	}
	if where.PermIdentsEgt != "" {
		tx.Where("perm_idents >= ?", where.PermIdentsEgt)
	}
	if len(where.PermIdentsNotIn) > 0 {
		tx.Where("perm_idents not in ?", where.PermIdentsNotIn)
	}
	if where.PermIdentsSort != "" {
		if where.PermIdentsSort == "asc" {
			tx.Order("perm_idents asc")
		} else {
			tx.Order("perm_idents desc")
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
func (rolesRepo *RolesRepoGorm) GetTotalCount(where repoInterface.RolesSearchWhere) int64 {
	tx := rolesRepo.GetSearchWhereTx(where, nil)
	var r int64
	tx.Count(&r)
	return r
}

func (rolesRepo *RolesRepoGorm) GetList(where repoInterface.RolesSearchWhere) []model.Roles {
	roles := make([]model.Roles, 0, where.SelectParams.RetSize)
	tx := rolesRepo.GetSearchWhereTx(where, nil)
	tx.Find(&roles)
	return roles
}

func (rolesRepo *RolesRepoGorm) GetByID(ID uint64, _select ...string) model.Roles {
	roles := model.Roles{}
	tx := rolesRepo.Orm.Where("id", ID)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&roles)
	return roles
}
func (rolesRepo *RolesRepoGorm) GetByName(name string, _select ...string) []model.Roles {
	roles := make([]model.Roles, 0)
	tx := rolesRepo.Orm.
		Where("name", name)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&roles)
	return roles
}

func (rolesRepo *RolesRepoGorm) DeleteByName(name string) (rowsAffected int64, err error) {
	tx := rolesRepo.Orm.
		Where("name", name)
	r := tx.Delete(model.Roles{})
	return r.RowsAffected, r.Error
}

func (rolesRepo *RolesRepoGorm) GetByIDLock(ID uint64, _select ...string) (model.Roles, repoComm.ReleaseLock) {
	if !orm.IsBeginTransaction(rolesRepo.Orm) {
		panic("rolesRepo.GetByIDLock is must beginTransaction")
	}
	roles := model.Roles{}
	tx := orm.LockForUpdate(rolesRepo.Orm.Where("id", ID))
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&roles)

	//这里返回一个空的释放锁方法 因为gorm在事务提交或回滚后会自动释放
	return roles, func() {}
}

func (rolesRepo *RolesRepoGorm) ScanByWhere(where repoInterface.RolesSearchWhere, dest any) error {
	return rolesRepo.GetSearchWhereTx(where, nil).Find(dest).Error
}

func (rolesRepo *RolesRepoGorm) ScanByOrWhere(dest any, where ...repoInterface.RolesSearchWhere) error {
	tx := rolesRepo.Orm.Model(model.Roles{})
	for _, v := range where {
		tx.Or(rolesRepo.GetSearchWhereTx(v, nil))
	}
	return tx.Find(dest).Error
}

func (rolesRepo RolesRepoGorm) GetRolesByID(id ...uint64) []model.Roles {
	r := make([]model.Roles, 0, len(id))
	rolesRepo.Orm.Where("id in ?", id).Find(&r)
	return r
}
