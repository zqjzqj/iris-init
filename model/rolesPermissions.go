package model

import (
	"iris-init/model/mField"
)

type RolesPermissions struct {
	mField.FieldsPkUUidBinary
	RoleID          uint64 `gorm:"comment:角色id;index:idx_role_id"`
	PermissionIdent string `gorm:"size:215;not null;comment:权限标识---不是权限表的主键 是 method@path"`
	mField.FieldsTimeUnixModel
}

func (rp RolesPermissions) TableName() string {
	return "roles_permissions"
}
