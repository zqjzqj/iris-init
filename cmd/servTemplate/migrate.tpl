package migrates

import (
    "github.com/go-gormigrate/gormigrate/v2"
    "gorm.io/gorm"
    "big_data_new/model"
    "big_data_new/orm"
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
