package model

import (
	"iris-init/model/mField"
	"time"
)

const (
	RoleAdmin     = "*"
	RoleAdminName = "超级管理员"
)

type Roles struct {
	mField.FieldsPk            `mapstructure:",squash"`
	Name                       string   `gorm:"size:100;comment:角色名称;not null;index:idx_name"`
	Remark                     string   `gorm:"size:241;comment:角色备注;default:''"`
	PermIdents                 []string `gorm:"-"`
	mField.FieldsTimeUnixModel `mapstructure:",squash"`
}

func (role Roles) TableName() string {
	return "roles"
}

func (role Roles) ShowMap() map[string]interface{} {
	return map[string]interface{}{
		"ID":         role.ID,
		"Name":       role.Name,
		"Remark":     role.Remark,
		"IsSysRole":  role.IsSysRole(),
		"CreatedAt":  time.Unix(role.CreatedAt, 0).Format(time.DateTime),
		"UpdatedAt":  time.Unix(role.UpdatedAt, 0).Format(time.DateTime),
		"PermIdents": role.PermIdents,
	}
}

func (role Roles) IsSysRole() bool {
	for _, v := range GetSysRoles() {
		if v.ID == role.ID {
			return true
		}
	}
	return false
}

// 返回系统内置角色
func GetSysRoles() []Roles {
	return nil
}
