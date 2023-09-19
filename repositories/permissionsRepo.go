package repositories

import (
	"gorm.io/gorm"
	"iris-init/model"
	"iris-init/orm"
	"iris-init/repositories/repoComm"
	"iris-init/repositories/repoInterface"
)

type PermissionsRepoGorm struct {
	repoComm.RepoGorm
}

func NewPermissionsRepo() repoInterface.PermissionsRepo {
	return &PermissionsRepoGorm{repoComm.NewRepoGorm()}
}

// 截断表
func (permRepo PermissionsRepoGorm) TruncateTable() {
	permRepo.Orm.Exec("truncate table " + new(model.Permissions).TableName())
}

func (permRepo PermissionsRepoGorm) GetByID(id uint64, _select ...string) model.Permissions {
	if id == 0 {
		return model.Permissions{}
	}
	perm := model.Permissions{}
	if len(_select) > 0 {
		permRepo.Orm.Select(_select).Where("id", id).First(&perm)
		return perm
	}
	permRepo.Orm.Where("id", id).First(&perm)
	return perm
}

func (permRepo PermissionsRepoGorm) GetByIdent(ident string, _select ...string) model.Permissions {
	if ident == "" {
		return model.Permissions{}
	}
	perm := model.Permissions{}
	if len(_select) > 0 {
		permRepo.Orm.Select(_select).Where("ident", ident).First(&perm)
		return perm
	}
	permRepo.Orm.Where("ident", ident).First(&perm)
	return perm
}

func (permRepo PermissionsRepoGorm) Save(perm *model.Permissions, _select ...string) error {
	return repoComm.SaveModel(permRepo.Orm, perm, _select...)
}

func (permissionsRepo *PermissionsRepoGorm) GetSearchWhereTx(where repoInterface.PermissionsSearchWhere, tx0 *gorm.DB) *gorm.DB {
	var tx *gorm.DB
	if tx0 == nil {
		tx = permissionsRepo.Orm.Model(model.Permissions{})
	} else {
		tx = tx0.Model(model.Permissions{})
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
	if where.Method != "" {
		tx.Where("method", where.Method)
	}
	if where.MethodNeq != "" {
		tx.Where("method <> ?", where.MethodNeq)
	}
	if where.MethodNull {
		tx.Where("method is null")
	}
	if where.MethodLike != "" {
		tx.Where("method like ?", "%"+where.MethodLike+"%")
	}

	if len(where.MethodIn) > 0 {
		tx.Where("method in ?", where.MethodIn)
	}

	if where.MethodNotNull {
		tx.Where("method is not null")
	}
	if where.MethodSort != "" {
		if where.MethodSort == "asc" {
			tx.Order("method asc")
		} else {
			tx.Order("method desc")
		}
	}
	//需要额外调整
	if where.Path != "" {
		tx.Where("path", where.Path)
	}
	if where.PathNeq != "" {
		tx.Where("path <> ?", where.PathNeq)
	}
	if where.PathNull {
		tx.Where("path is null")
	}
	if where.PathLike != "" {
		tx.Where("path like ?", "%"+where.PathLike+"%")
	}

	if len(where.PathIn) > 0 {
		tx.Where("path in ?", where.PathIn)
	}

	if where.PathNotNull {
		tx.Where("path is not null")
	}
	if where.PathSort != "" {
		if where.PathSort == "asc" {
			tx.Order("path asc")
		} else {
			tx.Order("path desc")
		}
	}
	//需要额外调整
	if where.Sort != "" {
		tx.Where("sort", where.Sort)
	}
	if where.SortNeq != "" {
		tx.Where("sort <> ?", where.SortNeq)
	}
	if where.SortNull {
		tx.Where("sort is null")
	}

	if len(where.SortIn) > 0 {
		tx.Where("sort in ?", where.SortIn)
	}

	if where.SortNotNull {
		tx.Where("sort is not null")
	}
	if where.SortLt != "" {
		tx.Where("sort < ?", where.SortLt)
	}
	if where.SortElt != "" {
		tx.Where("sort <= ?", where.SortElt)
	}
	if where.SortGt != "" {
		tx.Where("sort > ?", where.SortGt)
	}
	if where.SortEgt != "" {
		tx.Where("sort >= ?", where.SortEgt)
	}
	if len(where.SortNotIn) > 0 {
		tx.Where("sort not in ?", where.SortNotIn)
	}
	if where.SortSort != "" {
		if where.SortSort == "asc" {
			tx.Order("sort asc")
		} else {
			tx.Order("sort desc")
		}
	}
	//需要额外调整
	if where.Ident != "" {
		tx.Where("ident", where.Ident)
	}
	if where.IdentNeq != "" {
		tx.Where("ident <> ?", where.IdentNeq)
	}
	if where.IdentNull {
		tx.Where("ident is null")
	}
	if where.IdentLike != "" {
		tx.Where("ident like ?", "%"+where.IdentLike+"%")
	}

	if len(where.IdentIn) > 0 {
		tx.Where("ident in ?", where.IdentIn)
	}

	if where.IdentNotNull {
		tx.Where("ident is not null")
	}
	if where.IdentSort != "" {
		if where.IdentSort == "asc" {
			tx.Order("ident asc")
		} else {
			tx.Order("ident desc")
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

func (permRepo PermissionsRepoGorm) GetListPreloadChildren_2() []model.Permissions {
	return permRepo.GetList(repoInterface.PermissionsSearchWhere{
		Pid: "0",
		SelectParams: repoComm.SelectFrom{
			RetSize: 30,
			OrderBy: []repoComm.OrderByParams{{
				Column: "sort",
				Desc:   false,
			}, {
				Column: "id",
				Desc:   false,
			}},
			Preload: []repoComm.PreloadParams{
				{
					Query: "Children.Children",
				},
			},
		},
	})
}

func (permRepo PermissionsRepoGorm) GetListAsMenu(idents []string) []model.Permissions {
	return permRepo.GetList(repoInterface.PermissionsSearchWhere{
		Pid:     "0",
		IdentIn: idents,
		SelectParams: repoComm.SelectFrom{
			RetSize: 30,
			OrderBy: []repoComm.OrderByParams{{
				Column: "sort",
				Desc:   false,
			}, {
				Column: "id",
				Desc:   false,
			}},
			Preload: []repoComm.PreloadParams{
				{
					Query: "Children",
					Args: func() repoComm.SelectFrom {
						var where []repoComm.WhereParams
						if len(idents) > 0 {
							where = []repoComm.WhereParams{
								{
									Query: "ident in ?",
									Args:  []interface{}{idents},
								},
							}
						}
						return repoComm.SelectFrom{
							Where: where,
							OrderBy: []repoComm.OrderByParams{
								{
									Column: "sort",
									Desc:   false,
								},
								{
									Column: "id",
									Desc:   false,
								},
							}}
					},
				},
			},
		},
	})
}

func (permRepo PermissionsRepoGorm) GetListPreloadChildren(where repoInterface.PermissionsSearchWhere) []model.Permissions {
	where.SelectParams.Preload = []repoComm.PreloadParams{
		{Query: "Children"},
	}
	return permRepo.GetList(where)
}

func (permRepo PermissionsRepoGorm) GetList(where repoInterface.PermissionsSearchWhere) []model.Permissions {
	var perm []model.Permissions
	if where.SelectParams.Limit > 0 {
		perm = make([]model.Permissions, 0, where.SelectParams.Limit)
	} else {
		if where.SelectParams.RetSize > 0 {
			perm = make([]model.Permissions, 0, where.SelectParams.RetSize)
		} else {
			perm = make([]model.Permissions, 0, 15)
		}
	}
	tx := permRepo.GetSearchWhereTx(where, nil)
	where.SelectParams.SetTxGorm(tx).Find(&perm)
	return perm
}

func (permRepo PermissionsRepoGorm) GetOrCreatePermissionByName(name string, pid uint64, level uint8, sort uint) (model.Permissions, error) {
	perm := model.Permissions{}
	if permRepo.Orm.
		Where("pid", pid).
		Where("name", name).
		First(&perm).RowsAffected == 0 {
		//写入
		perm.Pid = pid
		perm.Name = name
		perm.Level = level
		perm.Sort = sort
		perm.GenerateIdent()
		err := permRepo.Save(&perm)
		return perm, err
	}
	return perm, nil
}

func (permissionsRepo *PermissionsRepoGorm) DeleteByLevel(level uint8) (rowsAffected int64, err error) {
	tx := permissionsRepo.Orm.
		Where("level", level)
	r := tx.Delete(model.Permissions{})
	return r.RowsAffected, r.Error
}

func (permissionsRepo *PermissionsRepoGorm) GetByIDLock(ID uint64, _select ...string) (model.Permissions, repoComm.ReleaseLock) {
	if !orm.IsBeginTransaction(permissionsRepo.Orm) {
		panic("permissionsRepo.GetByIDLock is must beginTransaction")
	}
	permissions := model.Permissions{}
	tx := orm.LockForUpdate(permissionsRepo.Orm.Where("id", ID))
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&permissions)

	//这里返回一个空的释放锁方法 因为gorm在事务提交或回滚后会自动释放
	return permissions, func() {}
}

func (permissionsRepo *PermissionsRepoGorm) ScanByWhere(where repoInterface.PermissionsSearchWhere, dest any) error {
	return permissionsRepo.GetSearchWhereTx(where, nil).Find(dest).Error
}

func (permissionsRepo *PermissionsRepoGorm) ScanByOrWhere(dest any, where ...repoInterface.PermissionsSearchWhere) error {
	tx := permissionsRepo.Orm.Model(model.Permissions{})
	for _, v := range where {
		tx.Or(permissionsRepo.GetSearchWhereTx(v, nil))
	}
	return tx.Find(dest).Error
}

func (permissionsRepo *PermissionsRepoGorm) DeleteByIdent(ident string) (rowsAffected int64, err error) {
	tx := permissionsRepo.Orm.
		Where("ident", ident)
	r := tx.Delete(model.Permissions{})
	return r.RowsAffected, r.Error
}
func (permissionsRepo *PermissionsRepoGorm) GetByPid_Name(pid uint64, name string, _select ...string) model.Permissions {
	permissions := model.Permissions{}
	tx := permissionsRepo.Orm.
		Where("pid", pid).
		Where("name", name)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&permissions)
	return permissions
}

func (permissionsRepo *PermissionsRepoGorm) DeleteByPid_Name(pid uint64, name string) (rowsAffected int64, err error) {
	tx := permissionsRepo.Orm.
		Where("pid", pid).
		Where("name", name)
	r := tx.Delete(model.Permissions{})
	return r.RowsAffected, r.Error
}
func (permissionsRepo *PermissionsRepoGorm) GetByLevel(level uint8, _select ...string) []model.Permissions {
	permissions := make([]model.Permissions, 0)
	tx := permissionsRepo.Orm.
		Where("level", level)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&permissions)
	return permissions
}

// 返回数据总数
func (permissionsRepo *PermissionsRepoGorm) GetTotalCount(where repoInterface.PermissionsSearchWhere) int64 {
	tx := permissionsRepo.GetSearchWhereTx(where, nil)
	var r int64
	tx.Count(&r)
	return r
}

func (permissionsRepo *PermissionsRepoGorm) UpdateByWhere(where repoInterface.PermissionsSearchWhere, data interface{}) (rowsAffected int64, err error) {
	tx := permissionsRepo.GetSearchWhereTx(where, nil)
	r := tx.Updates(data)
	return r.RowsAffected, r.Error
}

func (permissionsRepo *PermissionsRepoGorm) DeleteByWhere(where repoInterface.PermissionsSearchWhere) (rowsAffected int64, err error) {
	tx := permissionsRepo.GetSearchWhereTx(where, nil)
	r := tx.Delete(model.Permissions{})
	return r.RowsAffected, r.Error
}

// 为了避免更换源之后的一些麻烦 该方法不建议在仓库结构PermissionsRepoGorm以外使用
func (permissionsRepo *PermissionsRepoGorm) deleteByWhere(query string, args ...interface{}) (rowsAffected int64, err error) {
	tx := permissionsRepo.Orm.Where(query, args...).Delete(model.Permissions{})
	return tx.RowsAffected, tx.Error
}

func (permissionsRepo *PermissionsRepoGorm) DeleteByID(ID ...uint64) (rowsAffected int64, err error) {
	if len(ID) == 1 {
		return permissionsRepo.deleteByWhere("id", ID[0])
	}
	return permissionsRepo.deleteByWhere("id in ?", ID)
}

// 该方法需要自己去完善 GetSearchWhereTx方法内部
func (permissionsRepo *PermissionsRepoGorm) GetByWhere(where repoInterface.PermissionsSearchWhere) model.Permissions {
	permissions := model.Permissions{}
	_ = permissionsRepo.GetSearchWhereTx(where, nil).Limit(1).Find(&permissions)
	return permissions
}

// 该方法需要自己去完善 GetSearchWhereTx方法内部
func (permissionsRepo *PermissionsRepoGorm) GetIDByWhere(where repoInterface.PermissionsSearchWhere) []uint64 {
	var ID []uint64
	tx := permissionsRepo.GetSearchWhereTx(where, nil)
	tx.Select("id").Model(model.Permissions{}).Scan(&ID)
	return ID
}

func (permissionsRepo *PermissionsRepoGorm) SaveOmit(permissions *model.Permissions, _omit ...string) error {
	return repoComm.SaveModelOmit(permissionsRepo.Orm, permissions, _omit...)
}

// 这里因为gorm的缘故 传入的permissions主键必须不为空
func (permissionsRepo *PermissionsRepoGorm) Delete(permissions model.Permissions) (rowsAffected int64, err error) {
	tx := permissionsRepo.Orm.Delete(permissions)
	return tx.RowsAffected, tx.Error
}

func (permissionsRepo *PermissionsRepoGorm) Create(permissions *[]model.Permissions) error {
	return permissionsRepo.Orm.Create(permissions).Error
}
