package migrates

import (
    "github.com/go-gormigrate/gormigrate/v2"
    "gorm.io/gorm"
    "iris-init/model"
    "iris-init/orm"
)

type Migrate_{{.MigrateName}} struct{}

func (m Migrate_{{.MigrateName}}) GetId() string {
	return "Migrate_{{.MigrateName}}"
}

func (m Migrate_{{.MigrateName}}) Migrate() *gormigrate.Gormigrate {
	mArr := []ModelTableNameInterface{
	    {{- range .Models}}
	    model.{{.}}{},
	    {{- end}}
	}

	return gormigrate.New(orm.GetDb(), gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: m.GetId(),
			Migrate: func(tx *gorm.DB) error {
				for _, ma := range mArr {
					err := tx.Migrator().AutoMigrate(ma)
					if err != nil {
						return err
					}
				}
                return nil
			},
			Rollback: func(tx *gorm.DB) error {
				for _, ma := range mArr {
					err := tx.Migrator().DropTable(ma.TableName())
					if err != nil {
						return err
					}
				}
				return nil
			},
		},
	})
}
