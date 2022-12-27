package model

import "iris-init/model/mField"

type Area struct {
	mField.FieldsPk
	Pid        uint   `gorm:"type:int(11) unsigned;default:0;index:idx_pid"`
	ShortName  string `gorm:"size:100;default:'';comment:简称"`
	Name       string `gorm:"size:100;default:'';comment:名称"`
	MergerName string `gorm:"size:241;default:'';comment:全称"`
	Level      uint8  `gorm:"comment:层级 1 2 3 省市区县;index:idx_level"`
	PinYin     string `gorm:"size:100;comment:拼音;default:'';"`
	Code       string `gorm:"size:100;comment:长途区号;default:''"`
	ZipCode    string `gorm:"size:100;default:'';comment:邮编"`
	First      string `gorm:"size:50;default:'';首字母;index:idx_first"`
	Lng        string `gorm:"size:100;default:'';comment:经度"`
	Lat        string `gorm:"size:100;default:'';comment:纬度"`
	Children   []Area `gorm:"foreignKey:ID;references:Pid;association_autoupdate:false;association_autocreate:false"`
}

func (area Area) TableName() string {
	return "area"
}
