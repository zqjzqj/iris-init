package mField

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FieldsPk struct {
	ID uint64 `gorm:"primarykey" label:"ID"`
}

type FieldsPkUUidBinary struct {
	ID []byte `gorm:"primarykey;type:binary(16)" label:"ID"  OnlyRead:"true"`
}

func (pk *FieldsPkUUidBinary) BeforeCreate(tx *gorm.DB) (err error) {
	pk.ID, _ = uuid.New().MarshalBinary()
	return
}

func (pk FieldsPkUUidBinary) GetIDString() string {
	_uuid := uuid.New()
	_ = _uuid.UnmarshalBinary(pk.ID)
	return _uuid.String()
}
