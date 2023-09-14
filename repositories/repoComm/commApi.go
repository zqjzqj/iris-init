package repoComm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ReleaseLock func()

type RepoInterface interface {
	SetOrm(orm any) //该方法主要是在事务中 修改当前仓库的session
	GetOrm() any
	ResetOrm()     //该方法用户还原被修改的orm
	ResetLastOrm() // 还原上一次事务的orm
	//第一个参数为执行事务的代码func
	//第二个参数为关联的数据仓库
	//第三个参数为关联的其他数据仓库 与第二个参与一样 分开是因为必须要求有一个传递一个关联仓库就算是nil值以防开发时忘记
	Transaction(f func() error, _repos ...RepoInterface) error
}

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

func SaveModelOmit(_orm, _model any, _omit ...string) error {
	switch tx := _orm.(type) {
	case *gorm.DB:
		if len(_omit) > 0 {
			_omit = append(_omit, clause.Associations)
			return tx.Omit(_omit...).Save(_model).Error
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
	Raw     bool
}

// 用于关联预加载
type PreloadParams struct {
	Query string
	Args  func() SelectFrom
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
			tx.Order(clause.OrderByColumn{Column: clause.Column{Name: v.Column, Raw: v.Raw}, Desc: v.Desc, Reorder: v.Reorder})
		}
	}
	if len(selectFrom.Select) > 0 {
		tx.Select(selectFrom.Select)
	}
	if len(selectFrom.Preload) > 0 {
		for _, v := range selectFrom.Preload {
			if v.Args == nil {
				tx.Preload(v.Query)
			} else {
				_args := v.Args()
				tx.Preload(v.Query, func(_tx *gorm.DB) *gorm.DB {
					//这里where为空主要是需要去给_tx一个AddClause 不然SetTxGorm返回的tx无法生效
					return _args.SetTxGorm(_tx.Where(""))
				})
			}
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
