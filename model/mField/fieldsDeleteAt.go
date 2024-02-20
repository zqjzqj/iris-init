package mField

import "gorm.io/plugin/soft_delete"

type FieldsDeleteAt struct {
	DeletedAt soft_delete.DeletedAt `gorm:"index:delete_at;" soft_delete:"true"`
}
