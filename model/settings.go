package model

import (
	"iris-init/model/mField"
	"time"
)

const (
	SettingsKeyWebsiteTitle = "websiteTitle"
)

type Settings struct {
	mField.FieldsPk
	Key   string `gorm:"size:200;index:idx_settings_key,unique"`
	Name  string `gorm:"size:30;not null"`
	Desc  string `gorm:"size:100;default:''"`
	Value string `gorm:"type:text;"`
	mField.FieldsTimeUnixModel
}

func (settings Settings) TableName() string {
	return "settings"
}

func (settings Settings) ShowMap() map[string]interface{} {
	return map[string]interface{}{
		"ID":        settings.ID,
		"Key":       settings.Key,
		"Name":      settings.Name,
		"Desc":      settings.Desc,
		"Value":     settings.Value,
		"CreatedAt": time.Unix(settings.CreatedAt, 0).Format(time.DateTime),
		"UpdatedAt": time.Unix(settings.UpdatedAt, 0).Format(time.DateTime),
	}
}
