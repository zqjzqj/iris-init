package model

import (
	"iris-init/model/mField"
)

type RolesAdmin struct {
	mField.FieldsPkUUidBinary
	RoleID  uint64 `gorm:"comment:角色id;index:idx_role_admin_id"`
	AdminID uint64 `gorm:"comment:管理员id;index:idx_role_admin_id;index:idx_admin_id"`
	Admin   Admin  `gorm:"foreignKey:ID;references:AdminID;"`
	Role    Roles  `gorm:"foreignKey:ID;references:RoleID;"`
	mField.FieldsTimeUnixModel
}

func (ra RolesAdmin) TableName() string {
	return "roles_admin"
}

func (ra RolesAdmin) ShowMap() map[string]interface{} {
	return nil
}
