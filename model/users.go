package model

import "iris-init/model/mField"

type Users struct {
	mField.FieldsPk
	Username string
	mField.FieldsTimeUnixModel
}

func (user Users) TableName() string {
	return "users"
}

func (user Users) ShowMap() map[string]interface{} {
	return map[string]interface{}{
		"ID":        user.ID,
		"Username":  user.Username,
		"CreatedAt": user.CreatedAt,
		"UpdatedAt": user.UpdatedAt,
	}
}
