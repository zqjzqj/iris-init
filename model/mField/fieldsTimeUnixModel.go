package mField

type FieldsTimeUnixModel struct {
	CreatedAt int64 `gorm:"autoCreateTime;type:int(11) unsigned;" label:"创建时间"`
	UpdatedAt int64 `gorm:"autoUpdateTime;type:int(11) unsigned;" label:"更新时间"`
}
