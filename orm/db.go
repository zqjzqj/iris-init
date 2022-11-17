package orm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func LockForUpdate(tx *gorm.DB) *gorm.DB {
	return tx.Clauses(clause.Locking{Strength: "UPDATE"})
}

//这里是默认的db 后续如果跟换数据库驱动的话 可以在这里写代码
func GetDb() *gorm.DB {
	return GetMysqlDef().DB
}

func IsBeginTransaction(tx *gorm.DB) bool {
	if tx == nil {
		return false
	}
	if committer, ok := tx.Statement.ConnPool.(gorm.TxCommitter); ok && committer != nil {
		return true
	}
	return false
}

func IsExists(tx *gorm.DB, query string, args ...any) bool {
	var i int64
	if tx.Where(query, args...).Count(&i); i > 0 {
		return true
	}
	return false
}

//禁用所有关联
func OmitAssoc(tx *gorm.DB) *gorm.DB {
	return tx.Omit(clause.Associations)
}

func SaveOmitAssoc(tx *gorm.DB, value interface{}) *gorm.DB {
	return OmitAssoc(tx).Save(value)
}

func CreateOmitAssoc(tx *gorm.DB, value interface{}) *gorm.DB {
	return OmitAssoc(tx).Create(value)
}
