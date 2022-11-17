package mField

type FieldsTimeUnixModel struct {
	CreatedAt int64 `gorm:"autoCreateTime;type:int(11) unsigned;" mapstructure:"created_at"`
	UpdatedAt int64 `gorm:"autoUpdateTime;type:int(11) unsigned;" mapstructure:"updated_at"`
}
