package migrates

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"iris-init/logs"
)

type MigrateInterface interface {
	GetId() string
	Migrate() *gormigrate.Gormigrate
}

type ModelTableNameInterface interface {
	TableName() string
}

var _migrates = []MigrateInterface{
	MigrateInit{},
}

func Migrate() {
	_migrates = append(_migrates, MM...)
	for _, m := range _migrates {
		gm := m.Migrate()
		if err := gm.Migrate(); err != nil {
			_ = gm.RollbackLast()
			logs.Fatal("迁移"+m.GetId()+"执行失败", err)
		}
	}
	logs.PrintlnSuccess("migrate success !")
}

func Rollback(id string) {
	if id == "" {
		return
	}
	_migrates = append(_migrates, MM...)
	for _, m := range _migrates {
		if id == m.GetId() {
			gm := m.Migrate()
			if err := gm.RollbackLast(); err != nil {
				logs.Fatal("迁移回滚"+m.GetId()+"执行失败", err)
			}
			break
		}
	}
	logs.PrintlnSuccess("migrate rollback success !")
}
