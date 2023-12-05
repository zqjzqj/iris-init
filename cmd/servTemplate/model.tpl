package model

import (
	"iris-init/model/mField"
	"time"
)

type {{.Model}} struct {
	mField.FieldsPk
	mField.FieldsTimeUnixModel
}

func ({{.Alias}} {{.Model}}) TableName() string {
	return "{{.TableName}}"
}

func ({{.Alias}} {{.Model}}) ShowMap() map[string]interface{} {
	return map[string]interface{}{
		"ID":         {{.Alias}}.ID,
		"CreatedAt":  time.Unix({{.Alias}}.CreatedAt, 0).Format(time.DateTime),
		"UpdatedAt":  time.Unix({{.Alias}}.UpdatedAt, 0).Format(time.DateTime),
	}
}
