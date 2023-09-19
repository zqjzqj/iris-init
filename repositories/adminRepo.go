package repositories

import (
	"gorm.io/gorm"
	"iris-init/model"
	"iris-init/orm"
	"iris-init/repositories/repoComm"
	"iris-init/repositories/repoInterface"
)

type AdminRepoGorm struct {
	repoComm.RepoGorm
}

func NewAdminRepo() repoInterface.AdminRepo {
	return &AdminRepoGorm{repoComm.NewRepoGorm()}
}

// 该方法需要自己去完善 GetSearchWhereTx方法内部
func (adminRepo *AdminRepoGorm) GetByWhere(where repoInterface.AdminSearchWhere) model.Admin {
	admin := model.Admin{}
	_ = adminRepo.GetSearchWhereTx(where, nil).Limit(1).Find(&admin)
	return admin
}

// 该方法需要自己去完善 GetSearchWhereTx方法内部
func (adminRepo *AdminRepoGorm) GetIDByWhere(where repoInterface.AdminSearchWhere) []uint64 {
	var ID []uint64
	tx := adminRepo.GetSearchWhereTx(where, nil)
	tx.Select("id").Model(model.Admin{}).Scan(&ID)
	return ID
}

func (adminRepo *AdminRepoGorm) Create(admin *[]model.Admin) error {
	return adminRepo.Orm.Create(admin).Error
}

func (adminRepo *AdminRepoGorm) Save(admin *model.Admin, _select ...string) error {
	return repoComm.SaveModel(adminRepo.Orm, admin, _select...)
}

func (adminRepo *AdminRepoGorm) SaveOmit(admin *model.Admin, _omit ...string) error {
	return repoComm.SaveModelOmit(adminRepo.Orm, admin, _omit...)
}

// 这里因为gorm的缘故 传入的admin主键必须不为空
func (adminRepo *AdminRepoGorm) Delete(admin model.Admin) (rowsAffected int64, err error) {
	tx := adminRepo.Orm.Delete(admin)
	return tx.RowsAffected, tx.Error
}

// 为了避免更换源之后的一些麻烦 该方法不建议在仓库结构AdminRepoGorm以外使用
func (adminRepo *AdminRepoGorm) deleteByWhere(query string, args ...interface{}) (rowsAffected int64, err error) {
	tx := adminRepo.Orm.Where(query, args...).Delete(model.Admin{})
	return tx.RowsAffected, tx.Error
}

func (adminRepo *AdminRepoGorm) DeleteByID(ID ...uint64) (rowsAffected int64, err error) {
	if len(ID) == 1 {
		return adminRepo.deleteByWhere("id", ID[0])
	}
	return adminRepo.deleteByWhere("id in ?", ID)
}

func (adminRepo *AdminRepoGorm) UpdateByWhere(where repoInterface.AdminSearchWhere, data interface{}) (rowsAffected int64, err error) {
	tx := adminRepo.GetSearchWhereTx(where, nil)
	r := tx.Updates(data)
	return r.RowsAffected, r.Error
}

func (adminRepo *AdminRepoGorm) DeleteByWhere(where repoInterface.AdminSearchWhere) (rowsAffected int64, err error) {
	tx := adminRepo.GetSearchWhereTx(where, nil)
	r := tx.Delete(model.Admin{})
	return r.RowsAffected, r.Error
}

func (adminRepo *AdminRepoGorm) GetSearchWhereTx(where repoInterface.AdminSearchWhere, tx0 *gorm.DB) *gorm.DB {
	var tx *gorm.DB
	if tx0 == nil {
		tx = adminRepo.Orm.Model(model.Admin{})
	} else {
		tx = tx0.Model(model.Admin{})
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
	if len(where.IDIn) > 0 {
		tx.Where("id in ?", where.IDIn)
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
	if where.Username != "" {
		tx.Where("username", where.Username)
	}
	if where.UsernameNeq != "" {
		tx.Where("username <> ?", where.UsernameNeq)
	}
	if where.UsernameNull {
		tx.Where("username is null")
	}
	if where.UsernameNotNull {
		tx.Where("username is not null")
	}
	if where.UsernameLike != "" {
		tx.Where("username like ?", "%"+where.UsernameLike+"%")
	}
	//需要额外调整
	if where.Phone != "" {
		tx.Where("phone", where.Phone)
	}
	if where.PhoneNeq != "" {
		tx.Where("phone <> ?", where.PhoneNeq)
	}
	if where.PhoneNull {
		tx.Where("phone is null")
	}
	if where.PhoneNotNull {
		tx.Where("phone is not null")
	}
	if where.PhoneLike != "" {
		tx.Where("phone like ?", "%"+where.PhoneLike+"%")
	}
	//需要额外调整
	if where.QQ != "" {
		tx.Where("q_q", where.QQ)
	}
	if where.QQNeq != "" {
		tx.Where("q_q <> ?", where.QQNeq)
	}
	if where.QQNull {
		tx.Where("q_q is null")
	}
	if where.QQNotNull {
		tx.Where("q_q is not null")
	}
	if where.QQLike != "" {
		tx.Where("q_q like ?", "%"+where.QQLike+"%")
	}
	//需要额外调整
	if where.Status != "" {
		tx.Where("status", where.Status)
	}
	if where.StatusNeq != "" {
		tx.Where("status <> ?", where.StatusNeq)
	}
	if where.StatusNull {
		tx.Where("status is null")
	}
	if where.StatusNotNull {
		tx.Where("status is not null")
	}
	if where.StatusLt != "" {
		tx.Where("status < ?", where.StatusLt)
	}
	if where.StatusElt != "" {
		tx.Where("status <= ?", where.StatusElt)
	}
	if where.StatusGt != "" {
		tx.Where("status > ?", where.StatusGt)
	}
	if where.StatusEgt != "" {
		tx.Where("status >= ?", where.StatusEgt)
	}
	if len(where.StatusIn) > 0 {
		tx.Where("status in ?", where.StatusIn)
	}
	if len(where.StatusNotIn) > 0 {
		tx.Where("status not in ?", where.StatusNotIn)
	}
	if where.StatusSort != "" {
		if where.StatusSort == "asc" {
			tx.Order("status asc")
		} else {
			tx.Order("status desc")
		}
	}
	//需要额外调整
	if where.RealName != "" {
		tx.Where("real_name", where.RealName)
	}
	if where.RealNameNeq != "" {
		tx.Where("real_name <> ?", where.RealNameNeq)
	}
	if where.RealNameNull {
		tx.Where("real_name is null")
	}
	if where.RealNameNotNull {
		tx.Where("real_name is not null")
	}
	if where.RealNameLike != "" {
		tx.Where("real_name like ?", "%"+where.RealNameLike+"%")
	}
	//需要额外调整
	if where.Avatar != "" {
		tx.Where("avatar", where.Avatar)
	}
	if where.AvatarNeq != "" {
		tx.Where("avatar <> ?", where.AvatarNeq)
	}
	if where.AvatarNull {
		tx.Where("avatar is null")
	}
	if where.AvatarNotNull {
		tx.Where("avatar is not null")
	}
	if where.AvatarLike != "" {
		tx.Where("avatar like ?", "%"+where.AvatarLike+"%")
	}
	//需要额外调整
	if where.Sex != "" {
		tx.Where("sex", where.Sex)
	}
	if where.SexNeq != "" {
		tx.Where("sex <> ?", where.SexNeq)
	}
	if where.SexNull {
		tx.Where("sex is null")
	}
	if where.SexNotNull {
		tx.Where("sex is not null")
	}
	if where.SexLt != "" {
		tx.Where("sex < ?", where.SexLt)
	}
	if where.SexElt != "" {
		tx.Where("sex <= ?", where.SexElt)
	}
	if where.SexGt != "" {
		tx.Where("sex > ?", where.SexGt)
	}
	if where.SexEgt != "" {
		tx.Where("sex >= ?", where.SexEgt)
	}
	if len(where.SexIn) > 0 {
		tx.Where("sex in ?", where.SexIn)
	}
	if len(where.SexNotIn) > 0 {
		tx.Where("sex not in ?", where.SexNotIn)
	}
	if where.SexSort != "" {
		if where.SexSort == "asc" {
			tx.Order("sex asc")
		} else {
			tx.Order("sex desc")
		}
	}
	//需要额外调整
	if where.Password != "" {
		tx.Where("password", where.Password)
	}
	if where.PasswordNeq != "" {
		tx.Where("password <> ?", where.PasswordNeq)
	}
	if where.PasswordNull {
		tx.Where("password is null")
	}
	if where.PasswordNotNull {
		tx.Where("password is not null")
	}
	if where.PasswordLike != "" {
		tx.Where("password like ?", "%"+where.PasswordLike+"%")
	}
	//需要额外调整
	if where.Salt != "" {
		tx.Where("salt", where.Salt)
	}
	if where.SaltNeq != "" {
		tx.Where("salt <> ?", where.SaltNeq)
	}
	if where.SaltNull {
		tx.Where("salt is null")
	}
	if where.SaltNotNull {
		tx.Where("salt is not null")
	}
	if where.SaltLike != "" {
		tx.Where("salt like ?", "%"+where.SaltLike+"%")
	}
	//需要额外调整
	if where.Token != "" {
		tx.Where("token", where.Token)
	}
	if where.TokenNeq != "" {
		tx.Where("token <> ?", where.TokenNeq)
	}
	if where.TokenNull {
		tx.Where("token is null")
	}
	if where.TokenNotNull {
		tx.Where("token is not null")
	}
	if where.TokenLike != "" {
		tx.Where("token like ?", "%"+where.TokenLike+"%")
	}
	//需要额外调整
	if where.TokenStatus != "" {
		tx.Where("token_status", where.TokenStatus)
	}
	if where.TokenStatusNeq != "" {
		tx.Where("token_status <> ?", where.TokenStatusNeq)
	}
	if where.TokenStatusNull {
		tx.Where("token_status is null")
	}
	if where.TokenStatusNotNull {
		tx.Where("token_status is not null")
	}
	if where.TokenStatusLt != "" {
		tx.Where("token_status < ?", where.TokenStatusLt)
	}
	if where.TokenStatusElt != "" {
		tx.Where("token_status <= ?", where.TokenStatusElt)
	}
	if where.TokenStatusGt != "" {
		tx.Where("token_status > ?", where.TokenStatusGt)
	}
	if where.TokenStatusEgt != "" {
		tx.Where("token_status >= ?", where.TokenStatusEgt)
	}
	if len(where.TokenStatusIn) > 0 {
		tx.Where("token_status in ?", where.TokenStatusIn)
	}
	if len(where.TokenStatusNotIn) > 0 {
		tx.Where("token_status not in ?", where.TokenStatusNotIn)
	}
	if where.TokenStatusSort != "" {
		if where.TokenStatusSort == "asc" {
			tx.Order("token_status asc")
		} else {
			tx.Order("token_status desc")
		}
	}
	//需要额外调整
	if where.LastLoginTime != "" {
		tx.Where("last_login_time", where.LastLoginTime)
	}
	if where.LastLoginTimeNeq != "" {
		tx.Where("last_login_time <> ?", where.LastLoginTimeNeq)
	}
	if where.LastLoginTimeNull {
		tx.Where("last_login_time is null")
	}
	if where.LastLoginTimeNotNull {
		tx.Where("last_login_time is not null")
	}
	if where.LastLoginTimeLt != "" {
		tx.Where("last_login_time < ?", where.LastLoginTimeLt)
	}
	if where.LastLoginTimeElt != "" {
		tx.Where("last_login_time <= ?", where.LastLoginTimeElt)
	}
	if where.LastLoginTimeGt != "" {
		tx.Where("last_login_time > ?", where.LastLoginTimeGt)
	}
	if where.LastLoginTimeEgt != "" {
		tx.Where("last_login_time >= ?", where.LastLoginTimeEgt)
	}
	if len(where.LastLoginTimeIn) > 0 {
		tx.Where("last_login_time in ?", where.LastLoginTimeIn)
	}
	if len(where.LastLoginTimeNotIn) > 0 {
		tx.Where("last_login_time not in ?", where.LastLoginTimeNotIn)
	}
	if where.LastLoginTimeSort != "" {
		if where.LastLoginTimeSort == "asc" {
			tx.Order("last_login_time asc")
		} else {
			tx.Order("last_login_time desc")
		}
	}
	//需要额外调整
	if where.RolesID != "" {
		tx.Where("roles_id", where.RolesID)
	}
	if where.RolesIDNeq != "" {
		tx.Where("roles_id <> ?", where.RolesIDNeq)
	}
	if where.RolesIDNull {
		tx.Where("roles_id is null")
	}
	if where.RolesIDNotNull {
		tx.Where("roles_id is not null")
	}
	if where.RolesIDLike != "" {
		tx.Where("roles_id like ?", "%"+where.RolesIDLike+"%")
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
	if where.DescNotNull {
		tx.Where("desc is not null")
	}
	if where.DescLike != "" {
		tx.Where("desc like ?", "%"+where.DescLike+"%")
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
	if len(where.CreatedAtIn) > 0 {
		tx.Where("created_at in ?", where.CreatedAtIn)
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
	if len(where.UpdatedAtIn) > 0 {
		tx.Where("updated_at in ?", where.UpdatedAtIn)
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
	//需要额外调整
	if where.RolesIDSlices != "" {
		tx.Where("roles_id_slices", where.RolesIDSlices)
	}
	if where.RolesIDSlicesNeq != "" {
		tx.Where("roles_id_slices <> ?", where.RolesIDSlicesNeq)
	}
	if where.RolesIDSlicesNull {
		tx.Where("roles_id_slices is null")
	}
	if where.RolesIDSlicesNotNull {
		tx.Where("roles_id_slices is not null")
	}
	//需要额外调整
	if where.Permissions != "" {
		tx.Where("permissions", where.Permissions)
	}
	if where.PermissionsNeq != "" {
		tx.Where("permissions <> ?", where.PermissionsNeq)
	}
	if where.PermissionsNull {
		tx.Where("permissions is null")
	}
	if where.PermissionsNotNull {
		tx.Where("permissions is not null")
	}
	//需要额外调整
	if where.RolesName != "" {
		tx.Where("roles_name", where.RolesName)
	}
	if where.RolesNameNeq != "" {
		tx.Where("roles_name <> ?", where.RolesNameNeq)
	}
	if where.RolesNameNull {
		tx.Where("roles_name is null")
	}
	if where.RolesNameNotNull {
		tx.Where("roles_name is not null")
	}
	where.SelectParams.SetTxGorm(tx)
	return tx
}

// 返回数据总数
func (adminRepo *AdminRepoGorm) GetTotalCount(where repoInterface.AdminSearchWhere) int64 {
	tx := adminRepo.GetSearchWhereTx(where, nil)
	var r int64
	tx.Count(&r)
	return r
}

func (adminRepo *AdminRepoGorm) GetList(where repoInterface.AdminSearchWhere) []model.Admin {
	admin := make([]model.Admin, 0, where.SelectParams.RetSize)
	tx := adminRepo.GetSearchWhereTx(where, nil)
	tx.Find(&admin)
	return admin
}

func (adminRepo *AdminRepoGorm) GetByID(ID uint64, _select ...string) model.Admin {
	admin := model.Admin{}
	tx := adminRepo.Orm.Where("id", ID)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&admin)
	return admin
}
func (adminRepo *AdminRepoGorm) GetByPhone(phone string, _select ...string) model.Admin {
	admin := model.Admin{}
	tx := adminRepo.Orm.
		Where("phone", phone)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&admin)
	return admin
}

func (adminRepo *AdminRepoGorm) DeleteByPhone(phone string) (rowsAffected int64, err error) {
	tx := adminRepo.Orm.
		Where("phone", phone)
	r := tx.Delete(model.Admin{})
	return r.RowsAffected, r.Error
}
func (adminRepo *AdminRepoGorm) GetByToken(token string, _select ...string) model.Admin {
	admin := model.Admin{}
	tx := adminRepo.Orm.
		Where("token", token)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&admin)
	return admin
}

func (adminRepo *AdminRepoGorm) DeleteByToken(token string) (rowsAffected int64, err error) {
	tx := adminRepo.Orm.
		Where("token", token)
	r := tx.Delete(model.Admin{})
	return r.RowsAffected, r.Error
}
func (adminRepo *AdminRepoGorm) GetByUsername(username string, _select ...string) model.Admin {
	admin := model.Admin{}
	tx := adminRepo.Orm.
		Where("username", username)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&admin)
	return admin
}

func (adminRepo *AdminRepoGorm) DeleteByUsername(username string) (rowsAffected int64, err error) {
	tx := adminRepo.Orm.
		Where("username", username)
	r := tx.Delete(model.Admin{})
	return r.RowsAffected, r.Error
}
func (adminRepo *AdminRepoGorm) GetByRealName(realName string, _select ...string) []model.Admin {
	admin := make([]model.Admin, 0)
	tx := adminRepo.Orm.
		Where("real_name", realName)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&admin)
	return admin
}

func (adminRepo *AdminRepoGorm) DeleteByRealName(realName string) (rowsAffected int64, err error) {
	tx := adminRepo.Orm.
		Where("real_name", realName)
	r := tx.Delete(model.Admin{})
	return r.RowsAffected, r.Error
}
func (adminRepo *AdminRepoGorm) GetByStatus(status uint8, _select ...string) []model.Admin {
	admin := make([]model.Admin, 0)
	tx := adminRepo.Orm.
		Where("status", status)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&admin)
	return admin
}

func (adminRepo *AdminRepoGorm) DeleteByStatus(status uint8) (rowsAffected int64, err error) {
	tx := adminRepo.Orm.
		Where("status", status)
	r := tx.Delete(model.Admin{})
	return r.RowsAffected, r.Error
}

func (adminRepo *AdminRepoGorm) GetByIDLock(ID uint64, _select ...string) (model.Admin, repoComm.ReleaseLock) {
	if !orm.IsBeginTransaction(adminRepo.Orm) {
		panic("adminRepo.GetByIDLock is must beginTransaction")
	}
	admin := model.Admin{}
	tx := orm.LockForUpdate(adminRepo.Orm.Where("id", ID))
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&admin)

	//这里返回一个空的释放锁方法 因为gorm在事务提交或回滚后会自动释放
	return admin, func() {}
}

func (adminRepo *AdminRepoGorm) ScanByWhere(where repoInterface.AdminSearchWhere, dest any) error {
	return adminRepo.GetSearchWhereTx(where, nil).Find(dest).Error
}

func (adminRepo *AdminRepoGorm) ScanByOrWhere(dest any, where ...repoInterface.AdminSearchWhere) error {
	tx := adminRepo.Orm.Model(model.Admin{})
	for _, v := range where {
		tx.Or(adminRepo.GetSearchWhereTx(v, nil))
	}
	return tx.Find(dest).Error
}
