package repoComm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SaveModel(_orm, _model any, _select ...string) error {
	switch tx := _orm.(type) {
	case *gorm.DB:
		if len(_select) > 0 {
			return tx.Select(_select).Omit(clause.Associations).Save(_model).Error
		}
		return tx.Omit(clause.Associations).Save(_model).Error
	default:
		panic("不支持该orm类型..")
	}
	return nil
}

type OrderByParams struct {
	Column  string
	Desc    bool
	Reorder bool
}

//用于关联预加载
type PreloadParams struct {
	Query string
	Args  []interface{}
}

type WhereParams struct {
	Query string
	Args  []interface{}
	Func  func(orm any) any
}

type SelectFrom struct {
	RetSize int // 用来控制切片初始化容量
	Offset  int
	Limit   int
	Select  []string
	OrderBy []OrderByParams
	Preload []PreloadParams
	Where   []WhereParams
}

func (selectFrom SelectFrom) SetTxGorm(tx *gorm.DB) *gorm.DB {
	if selectFrom.Offset > -1 {
		tx.Offset(selectFrom.Offset)
	}
	if selectFrom.Limit > 0 {
		tx.Limit(selectFrom.Limit)
	}
	if len(selectFrom.OrderBy) > 0 {
		for _, v := range selectFrom.OrderBy {
			tx.Order(clause.OrderByColumn{Column: clause.Column{Name: v.Column}, Desc: v.Desc, Reorder: v.Reorder})
		}
	}
	if len(selectFrom.Select) > 0 {
		tx.Select(selectFrom.Select)
	}
	if len(selectFrom.Preload) > 0 {
		for _, v := range selectFrom.Preload {
			tx.Preload(v.Query, v.Args...)
		}
	}
	if len(selectFrom.Where) > 0 {
		for _, v := range selectFrom.Where {
			if v.Func != nil {
				_func := v.Func
				_func(tx)
			} else {
				tx.Where(v.Query, v.Args...)
			}
		}
	}
	return tx
}
