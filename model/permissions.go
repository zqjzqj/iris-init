package model

import (
	"fmt"
	"jd-fxl/global"
	"jd-fxl/model/mField"
)

const (
	//按顺序递增
	PermissionsLevelDir  = 0
	PermissionsLevelMenu = 1
	PermissionsLevelBtn  = 2
)

type Permissions struct {
	mField.FieldsPk
	Pid      uint64        `gorm:"comment:父级id;index:idx_pid,unique"`
	Level    uint8         `gorm:"comment:权限的类型级别，0目录级别,1菜单级别，2按钮级别;index:idx_level"`
	Name     string        `gorm:"size:100;comment:权限名称;not null;index:idx_pid,unique"`
	Method   string        `gorm:"size:15;default:'';comment:请求方法 GET POST DELETE等;"`
	Path     string        `gorm:"size:200;default:'';comment:权限路径;"`
	Sort     uint          `gorm:"type:int(11) unsigned;default:100;comment:排序 asc"`
	Ident    string        `gorm:"size:215;not null;comment:权限唯一标识, method@path组成【目录级别无路由则生成一个唯一的字符串填充】, 这个字段主要是便于查询;index:idx_ident,unique;"`
	Children []Permissions `gorm:"foreignKey:Pid;references:ID"`
	mField.FieldsTimeUnixModel
}

func (p Permissions) TableName() string {
	return "permissions"
}

func (p *Permissions) GenerateIdent() {
	if p.Method == "" || p.Path == "" {
		p.Ident = global.Md5(fmt.Sprintf("%d-%s", p.Level, p.Name))
		return
	}
	p.Ident = fmt.Sprintf("%s@%s", p.Method, p.Path)
}
