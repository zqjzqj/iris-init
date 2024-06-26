package model

import (
	"database/sql"
	"iris-init/global"
	"iris-init/model/mField"
	"strings"
	"time"
)

const (
	AdminRootID = 1
)

var AdmStatusDescMap = map[uint8]string{
	global.IsYes: "正常",
	global.IsNo:  "禁用",
}

type Admin struct {
	mField.FieldsPk `mapstructure:",squash"`
	Username        string         `gorm:"size:50;not null;comment:用户名/手机号码;index:idx_username,unique" label:"用户名"`
	Phone           sql.NullString `gorm:"size:15;default:null;index:idx_phone,unique"`
	QQ              string         `gorm:"size:20;default:'';comment:qq号码"`
	Status          uint8          `gorm:"default:1;comment:0禁用 1正常;index:idx_status" `
	RealName        string         `gorm:"size:20;comment:姓名;default:'';index:idx_name" mapstructure:"real_name"`
	Avatar          string         `gorm:"type:text;comment:头像"`
	Sex             uint8          `gorm:"comment:性别;comment:1男性，2女性,0或其他值为未知;default:0"`
	Password        string         `gorm:"type:char(32);default:'';comment:密码md5"`
	Salt            string         `gorm:"type:varchar(32);default:'';comment:盐"`
	Token           sql.NullString `gorm:"size:32;index:idx_token,unique;comment:用户登陆token"`
	TokenStatus     uint8          `gorm:"default:0;comment:0禁用 1正常" mapstructure:"token_status"`
	LastLoginTime   int64          `gorm:"type:int(11) unsigned;comment:最近一次登陆时间;default:0" mapstructure:"last_login_time"`

	//为防止其他异常 使用SetRoleID进行修改
	RolesID                      string `gorm:"type:text;comment:角色id ','分割" mapstructure:"roles_id"`
	Desc                         string `gorm:"type:text;comment:描述简介"`
	mField.FieldsTimeUnixModel   `mapstructure:",squash"`
	mField.FieldsExtendsJsonType `mapstructure:",squash"`
	RolesAdmin                   []RolesAdmin `gorm:"foreignKey:AdminID;references:ID;"`

	RolesIDSlices []string `gorm:"-"`
	Permissions   []string `gorm:"-"`
	RolesName     []string `gorm:"-"`
}

func (adm Admin) TableName() string {
	return "admin"
}

func (adm Admin) IsStatusYes() bool {
	return adm.Status == global.IsYes
}

func (adm Admin) IsRootAccount() bool {
	return adm.ID == AdminRootID
}

func (adm *Admin) SetRoleID(RoleID []string) {
	adm.RolesIDSlices = global.RemoveDuplicateElement(RoleID)
	adm.RolesID = strings.Join(adm.RolesIDSlices, ",")
}

func (adm Admin) GetRoleIDSlices() []string {
	return adm.RolesIDSlices
}

func (adm *Admin) RefreshRoleIDSlices() []string {
	if adm.RolesID == "" {
		return nil
	}
	adm.RolesIDSlices = strings.Split(adm.RolesID, ",")
	return adm.RolesIDSlices
}

func (adm *Admin) IsRootRole() bool {
	if adm.IsRootAccount() { //系统内置的root账户
		return true
	}
	if adm.RolesIDSlices == nil {
		adm.RefreshRoleIDSlices()
	}
	return global.InSlice(RoleAdmin, adm.RolesIDSlices)
}

func (adm Admin) CheckPwd(pwd string) bool {
	if global.PwdPlaintext2CipherText(pwd, adm.Salt) == adm.Password {
		return true
	}
	return false
}

func (adm Admin) TokenValid() bool {
	return adm.TokenStatus == global.IsYes
}

func (adm *Admin) SetRolesNameByRolesAdm() {
	rolesAdm_len := len(adm.RolesAdmin)
	if rolesAdm_len == 0 && !adm.IsRootRole() {
		return
	}
	if adm.IsRootRole() {
		//去掉多余的权限ID
		adm.Permissions = []string{RoleAdmin}
		adm.RolesName = []string{RoleAdminName}
	} else {
		adm.RolesName = make([]string, 0, rolesAdm_len)
	}
	for _, v := range adm.RolesAdmin {
		adm.RolesName = append(adm.RolesName, v.Role.Name)
	}
}

func (adm Admin) ShowMap() map[string]interface{} {
	if adm.IsRootRole() {
		adm.Permissions = nil
	}
	adm.SetRolesNameByRolesAdm()
	r := map[string]interface{}{
		"ID":            adm.ID,
		"Phone":         adm.Phone.String,
		"Username":      adm.Username,
		"RealName":      adm.RealName,
		"Avatar":        adm.Avatar,
		"Sex":           adm.Sex,
		"SexDesc":       global.SexDescMap[adm.Sex],
		"LastLoginTime": "",
		"CreatedAt":     time.Unix(adm.CreatedAt, 0).Format(time.DateTime),
		"UpdatedAt":     time.Unix(adm.UpdatedAt, 0).Format(time.DateTime),
		"RolesID":       adm.RolesID,
		"Status":        adm.Status,
		"StatusDesc":    AdmStatusDescMap[adm.Status],
		"QQ":            adm.QQ,
		"Desc":          adm.Desc,
		"Permissions":   adm.Permissions,
		"RolesName":     adm.RolesName,
		"RolesNameStr":  strings.Join(adm.RolesName, ","),
	}
	if adm.LastLoginTime > 0 {
		r["LastLoginTime"] = time.Unix(adm.LastLoginTime, 0).Format(time.DateTime)
	}
	if adm.Avatar == "" {
		r["Avatar"] = global.DefaultAvatar
	}
	return r
}

func (adm Admin) ShowMapHasToken() map[string]interface{} {
	r := adm.ShowMap()
	r["Token"] = adm.Token.String
	r["TokenStatus"] = adm.TokenStatus
	return r
}
