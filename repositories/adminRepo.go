package repositories

import (
	"big_data_new/model"
	"big_data_new/orm"
	"big_data_new/repositories/repoComm"
	"big_data_new/repositories/repoInterface"
	"gorm.io/gorm"
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
	if where.Username != "" {
		tx.Where("username", where.Username)
	}
	if where.UsernameNeq != "" {
		tx.Where("username <> ?", where.UsernameNeq)
	}
	if where.UsernameNull {
		tx.Where("username is null")
	}
	if where.UsernameLike != "" {
		tx.Where("username like ?", "%"+where.UsernameLike+"%")
	}

	if len(where.UsernameIn) > 0 {
		tx.Where("username in ?", where.UsernameIn)
	}

	if where.UsernameNotNull {
		tx.Where("username is not null")
	}
	if where.UsernameLt != "" {
		tx.Where("username < ?", where.UsernameLt)
	}
	if where.UsernameElt != "" {
		tx.Where("username <= ?", where.UsernameElt)
	}
	if where.UsernameGt != "" {
		tx.Where("username > ?", where.UsernameGt)
	}
	if where.UsernameEgt != "" {
		tx.Where("username >= ?", where.UsernameEgt)
	}
	if len(where.UsernameNotIn) > 0 {
		tx.Where("username not in ?", where.UsernameNotIn)
	}
	if where.UsernameSort != "" {
		if where.UsernameSort == "asc" {
			tx.Order("username asc")
		} else {
			tx.Order("username desc")
		}
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
	if where.PhoneLike != "" {
		tx.Where("phone like ?", "%"+where.PhoneLike+"%")
	}

	if len(where.PhoneIn) > 0 {
		tx.Where("phone in ?", where.PhoneIn)
	}

	if where.PhoneNotNull {
		tx.Where("phone is not null")
	}
	if where.PhoneLt != "" {
		tx.Where("phone < ?", where.PhoneLt)
	}
	if where.PhoneElt != "" {
		tx.Where("phone <= ?", where.PhoneElt)
	}
	if where.PhoneGt != "" {
		tx.Where("phone > ?", where.PhoneGt)
	}
	if where.PhoneEgt != "" {
		tx.Where("phone >= ?", where.PhoneEgt)
	}
	if len(where.PhoneNotIn) > 0 {
		tx.Where("phone not in ?", where.PhoneNotIn)
	}
	if where.PhoneSort != "" {
		if where.PhoneSort == "asc" {
			tx.Order("phone asc")
		} else {
			tx.Order("phone desc")
		}
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
	if where.QQLike != "" {
		tx.Where("q_q like ?", "%"+where.QQLike+"%")
	}

	if len(where.QQIn) > 0 {
		tx.Where("q_q in ?", where.QQIn)
	}

	if where.QQNotNull {
		tx.Where("q_q is not null")
	}
	if where.QQLt != "" {
		tx.Where("q_q < ?", where.QQLt)
	}
	if where.QQElt != "" {
		tx.Where("q_q <= ?", where.QQElt)
	}
	if where.QQGt != "" {
		tx.Where("q_q > ?", where.QQGt)
	}
	if where.QQEgt != "" {
		tx.Where("q_q >= ?", where.QQEgt)
	}
	if len(where.QQNotIn) > 0 {
		tx.Where("q_q not in ?", where.QQNotIn)
	}
	if where.QQSort != "" {
		if where.QQSort == "asc" {
			tx.Order("q_q asc")
		} else {
			tx.Order("q_q desc")
		}
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

	if len(where.StatusIn) > 0 {
		tx.Where("status in ?", where.StatusIn)
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
	if where.RealNameLike != "" {
		tx.Where("real_name like ?", "%"+where.RealNameLike+"%")
	}

	if len(where.RealNameIn) > 0 {
		tx.Where("real_name in ?", where.RealNameIn)
	}

	if where.RealNameNotNull {
		tx.Where("real_name is not null")
	}
	if where.RealNameLt != "" {
		tx.Where("real_name < ?", where.RealNameLt)
	}
	if where.RealNameElt != "" {
		tx.Where("real_name <= ?", where.RealNameElt)
	}
	if where.RealNameGt != "" {
		tx.Where("real_name > ?", where.RealNameGt)
	}
	if where.RealNameEgt != "" {
		tx.Where("real_name >= ?", where.RealNameEgt)
	}
	if len(where.RealNameNotIn) > 0 {
		tx.Where("real_name not in ?", where.RealNameNotIn)
	}
	if where.RealNameSort != "" {
		if where.RealNameSort == "asc" {
			tx.Order("real_name asc")
		} else {
			tx.Order("real_name desc")
		}
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
	if where.AvatarLike != "" {
		tx.Where("avatar like ?", "%"+where.AvatarLike+"%")
	}

	if len(where.AvatarIn) > 0 {
		tx.Where("avatar in ?", where.AvatarIn)
	}

	if where.AvatarNotNull {
		tx.Where("avatar is not null")
	}
	if where.AvatarLt != "" {
		tx.Where("avatar < ?", where.AvatarLt)
	}
	if where.AvatarElt != "" {
		tx.Where("avatar <= ?", where.AvatarElt)
	}
	if where.AvatarGt != "" {
		tx.Where("avatar > ?", where.AvatarGt)
	}
	if where.AvatarEgt != "" {
		tx.Where("avatar >= ?", where.AvatarEgt)
	}
	if len(where.AvatarNotIn) > 0 {
		tx.Where("avatar not in ?", where.AvatarNotIn)
	}
	if where.AvatarSort != "" {
		if where.AvatarSort == "asc" {
			tx.Order("avatar asc")
		} else {
			tx.Order("avatar desc")
		}
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

	if len(where.SexIn) > 0 {
		tx.Where("sex in ?", where.SexIn)
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
	if where.PasswordLike != "" {
		tx.Where("password like ?", "%"+where.PasswordLike+"%")
	}

	if len(where.PasswordIn) > 0 {
		tx.Where("password in ?", where.PasswordIn)
	}

	if where.PasswordNotNull {
		tx.Where("password is not null")
	}
	if where.PasswordLt != "" {
		tx.Where("password < ?", where.PasswordLt)
	}
	if where.PasswordElt != "" {
		tx.Where("password <= ?", where.PasswordElt)
	}
	if where.PasswordGt != "" {
		tx.Where("password > ?", where.PasswordGt)
	}
	if where.PasswordEgt != "" {
		tx.Where("password >= ?", where.PasswordEgt)
	}
	if len(where.PasswordNotIn) > 0 {
		tx.Where("password not in ?", where.PasswordNotIn)
	}
	if where.PasswordSort != "" {
		if where.PasswordSort == "asc" {
			tx.Order("password asc")
		} else {
			tx.Order("password desc")
		}
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
	if where.SaltLike != "" {
		tx.Where("salt like ?", "%"+where.SaltLike+"%")
	}

	if len(where.SaltIn) > 0 {
		tx.Where("salt in ?", where.SaltIn)
	}

	if where.SaltNotNull {
		tx.Where("salt is not null")
	}
	if where.SaltLt != "" {
		tx.Where("salt < ?", where.SaltLt)
	}
	if where.SaltElt != "" {
		tx.Where("salt <= ?", where.SaltElt)
	}
	if where.SaltGt != "" {
		tx.Where("salt > ?", where.SaltGt)
	}
	if where.SaltEgt != "" {
		tx.Where("salt >= ?", where.SaltEgt)
	}
	if len(where.SaltNotIn) > 0 {
		tx.Where("salt not in ?", where.SaltNotIn)
	}
	if where.SaltSort != "" {
		if where.SaltSort == "asc" {
			tx.Order("salt asc")
		} else {
			tx.Order("salt desc")
		}
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
	if where.TokenLike != "" {
		tx.Where("token like ?", "%"+where.TokenLike+"%")
	}

	if len(where.TokenIn) > 0 {
		tx.Where("token in ?", where.TokenIn)
	}

	if where.TokenNotNull {
		tx.Where("token is not null")
	}
	if where.TokenLt != "" {
		tx.Where("token < ?", where.TokenLt)
	}
	if where.TokenElt != "" {
		tx.Where("token <= ?", where.TokenElt)
	}
	if where.TokenGt != "" {
		tx.Where("token > ?", where.TokenGt)
	}
	if where.TokenEgt != "" {
		tx.Where("token >= ?", where.TokenEgt)
	}
	if len(where.TokenNotIn) > 0 {
		tx.Where("token not in ?", where.TokenNotIn)
	}
	if where.TokenSort != "" {
		if where.TokenSort == "asc" {
			tx.Order("token asc")
		} else {
			tx.Order("token desc")
		}
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

	if len(where.TokenStatusIn) > 0 {
		tx.Where("token_status in ?", where.TokenStatusIn)
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

	if len(where.LastLoginTimeIn) > 0 {
		tx.Where("last_login_time in ?", where.LastLoginTimeIn)
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
	if where.RolesId != "" {
		tx.Where("roles_id", where.RolesId)
	}
	if where.RolesIdNeq != "" {
		tx.Where("roles_id <> ?", where.RolesIdNeq)
	}
	if where.RolesIdNull {
		tx.Where("roles_id is null")
	}
	if where.RolesIdLike != "" {
		tx.Where("roles_id like ?", "%"+where.RolesIdLike+"%")
	}

	if len(where.RolesIdIn) > 0 {
		tx.Where("roles_id in ?", where.RolesIdIn)
	}

	if where.RolesIdNotNull {
		tx.Where("roles_id is not null")
	}
	if where.RolesIdLt != "" {
		tx.Where("roles_id < ?", where.RolesIdLt)
	}
	if where.RolesIdElt != "" {
		tx.Where("roles_id <= ?", where.RolesIdElt)
	}
	if where.RolesIdGt != "" {
		tx.Where("roles_id > ?", where.RolesIdGt)
	}
	if where.RolesIdEgt != "" {
		tx.Where("roles_id >= ?", where.RolesIdEgt)
	}
	if len(where.RolesIdNotIn) > 0 {
		tx.Where("roles_id not in ?", where.RolesIdNotIn)
	}
	if where.RolesIdSort != "" {
		if where.RolesIdSort == "asc" {
			tx.Order("roles_id asc")
		} else {
			tx.Order("roles_id desc")
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
	if where.DescLt != "" {
		tx.Where("desc < ?", where.DescLt)
	}
	if where.DescElt != "" {
		tx.Where("desc <= ?", where.DescElt)
	}
	if where.DescGt != "" {
		tx.Where("desc > ?", where.DescGt)
	}
	if where.DescEgt != "" {
		tx.Where("desc >= ?", where.DescEgt)
	}
	if len(where.DescNotIn) > 0 {
		tx.Where("desc not in ?", where.DescNotIn)
	}
	if where.DescSort != "" {
		if where.DescSort == "asc" {
			tx.Order("desc asc")
		} else {
			tx.Order("desc desc")
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

	if len(where.PermissionsIn) > 0 {
		tx.Where("permissions in ?", where.PermissionsIn)
	}

	if where.PermissionsNotNull {
		tx.Where("permissions is not null")
	}
	if where.PermissionsLt != "" {
		tx.Where("permissions < ?", where.PermissionsLt)
	}
	if where.PermissionsElt != "" {
		tx.Where("permissions <= ?", where.PermissionsElt)
	}
	if where.PermissionsGt != "" {
		tx.Where("permissions > ?", where.PermissionsGt)
	}
	if where.PermissionsEgt != "" {
		tx.Where("permissions >= ?", where.PermissionsEgt)
	}
	if len(where.PermissionsNotIn) > 0 {
		tx.Where("permissions not in ?", where.PermissionsNotIn)
	}
	if where.PermissionsSort != "" {
		if where.PermissionsSort == "asc" {
			tx.Order("permissions asc")
		} else {
			tx.Order("permissions desc")
		}
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

	if len(where.RolesNameIn) > 0 {
		tx.Where("roles_name in ?", where.RolesNameIn)
	}

	if where.RolesNameNotNull {
		tx.Where("roles_name is not null")
	}
	if where.RolesNameLt != "" {
		tx.Where("roles_name < ?", where.RolesNameLt)
	}
	if where.RolesNameElt != "" {
		tx.Where("roles_name <= ?", where.RolesNameElt)
	}
	if where.RolesNameGt != "" {
		tx.Where("roles_name > ?", where.RolesNameGt)
	}
	if where.RolesNameEgt != "" {
		tx.Where("roles_name >= ?", where.RolesNameEgt)
	}
	if len(where.RolesNameNotIn) > 0 {
		tx.Where("roles_name not in ?", where.RolesNameNotIn)
	}
	if where.RolesNameSort != "" {
		if where.RolesNameSort == "asc" {
			tx.Order("roles_name asc")
		} else {
			tx.Order("roles_name desc")
		}
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
