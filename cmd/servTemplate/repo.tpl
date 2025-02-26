package repositories

import (
	"gorm.io/gorm"
	"iris-init/model"
	"iris-init/orm"
	"iris-init/repositories/repoComm"
	"iris-init/repositories/repoInterface"
)

type {{.Model}}RepoGorm struct {
	repoComm.RepoGorm
}

func New{{.Model}}Repo() repoInterface.{{.Model}}Repo {
	return &{{.Model}}RepoGorm{repoComm.NewRepoGorm()}
}

//该方法需要自己去完善 GetSearchWhereTx方法内部
func ({{.Alias}}Repo *{{.Model}}RepoGorm) GetByWhere(where repoInterface.{{.Model}}SearchWhere) model.{{.Model}} {
	{{.Alias}} := model.{{.Model}}{}
	_ = {{.Alias}}Repo.GetSearchWhereTx(where, nil).Limit(1).Find(&{{.Alias}})
	return {{.Alias}}
}

//该方法需要自己去完善 GetSearchWhereTx方法内部
func ({{.Alias}}Repo *{{.Model}}RepoGorm) GetIDByWhere(where repoInterface.{{.Model}}SearchWhere) []{{.Pk.Type}} {
	var {{.Pk.Name}} []{{.Pk.Type}}
	tx := {{.Alias}}Repo.GetSearchWhereTx(where, nil)
	tx.Select("{{.Pk.NameSnake}}").Model(model.{{.Model}}{}).Scan(&{{.Pk.Name}})
	return {{.Pk.Name}}
}

func ({{.Alias}}Repo *{{.Model}}RepoGorm) Create({{.Alias}} *[]model.{{.Model}}) error {
	return {{.Alias}}Repo.Orm.Create({{.Alias}}).Error
}

func ({{.Alias}}Repo *{{.Model}}RepoGorm) Save({{.Alias}} *model.{{.Model}}, _select ...string) error {
	return repoComm.SaveModel({{.Alias}}Repo.Orm, {{.Alias}}, _select...)
}

func ({{.Alias}}Repo *{{.Model}}RepoGorm) SaveOmit({{.Alias}} *model.{{.Model}}, _omit ...string) error {
	return repoComm.SaveModelOmit({{.Alias}}Repo.Orm, {{.Alias}}, _omit...)
}

//这里因为gorm的缘故 传入的{{.Alias}}主键必须不为空
func ({{.Alias}}Repo *{{.Model}}RepoGorm) Delete({{.Alias}} model.{{.Model}}) (rowsAffected int64, err error) {
    tx := {{.Alias}}Repo.Orm.Delete({{.Alias}})
	return tx.RowsAffected, tx.Error
}

//为了避免更换源之后的一些麻烦 该方法不建议在仓库结构{{.Model}}RepoGorm以外使用
func ({{.Alias}}Repo *{{.Model}}RepoGorm) deleteByWhere(query string, args ...interface{}) (rowsAffected int64, err error) {
	tx := {{.Alias}}Repo.Orm.Where(query, args...).Delete(&model.{{.Model}}{})
	return tx.RowsAffected, tx.Error
}

func ({{.Alias}}Repo *{{.Model}}RepoGorm) DeleteByID({{.Pk.Name}} ...{{.Pk.Type}}) (rowsAffected int64, err error) {
	if len({{.Pk.Name}}) == 1 {
		return {{.Alias}}Repo.deleteByWhere("{{.Pk.NameSnake}}", {{.Pk.Name}}[0])
	}
	return {{.Alias}}Repo.deleteByWhere("{{.Pk.NameSnake}} in ?", {{.Pk.Name}})
}

func ({{$.Alias}}Repo *{{$.Model}}RepoGorm) UpdateByWhere(where repoInterface.{{.Model}}SearchWhere, data interface{}) (rowsAffected int64, err error) {
	tx := {{$.Alias}}Repo.GetSearchWhereTx(where, nil)
	r := tx.Updates(data)
	return r.RowsAffected, r.Error
}

func ({{$.Alias}}Repo *{{$.Model}}RepoGorm) DeleteByWhere(where repoInterface.{{.Model}}SearchWhere) (rowsAffected int64, err error) {
	tx := {{$.Alias}}Repo.GetSearchWhereTx(where, nil)
	r := tx.Delete(&model.{{$.Model}}{})
	return r.RowsAffected, r.Error
}

func ({{.Alias}}Repo *{{.Model}}RepoGorm) GetSearchWhereTx(where repoInterface.{{.Model}}SearchWhere, tx0 *gorm.DB) *gorm.DB {
	var tx *gorm.DB
	if tx0 == nil {
		tx = {{.Alias}}Repo.Orm.Model(model.{{.Model}}{})
	} else {
		tx = tx0.Model(model.{{.Model}}{})
	}
   {{- range .ModelField}}
    {{- if eq .IsStruct false}}
       //需要额外调整
        if where.{{.Name}}Any != nil {
           tx.Where(where.{{.Name}}Any)
       }
        if where.{{.Name}} != "" {
            tx.Where("{{.NameSnake}}", where.{{.Name}})
        }
        if where.{{.Name}}Neq != "" {
            tx.Where("{{.NameSnake}} <> ?", where.{{.Name}}Neq)
        }
        if where.{{.Name}}Null {
            tx.Where("{{.NameSnake}} is null")
        }

       {{- if eq .Type "string"}}
        if where.{{.Name}}Like != "" {
           tx.Where("{{.NameSnake}} like ?", "%"+where.{{.Name}}Like+"%")
        }
       {{- end}}

        if len(where.{{.Name}}In) > 0 {
          tx.Where("{{.NameSnake}} in ?", where.{{.Name}}In)
        }

        if where.{{.Name}}NotNull {
            tx.Where("{{.NameSnake}} is not null")
        }
        if where.{{.Name}}Lt != "" {
          tx.Where("{{.NameSnake}} < ?", where.{{.Name}}Lt)
        }
        if where.{{.Name}}Elt != "" {
          tx.Where("{{.NameSnake}} <= ?", where.{{.Name}}Elt)
        }
        if where.{{.Name}}Gt != "" {
          tx.Where("{{.NameSnake}} > ?", where.{{.Name}}Gt)
        }
        if where.{{.Name}}Egt != "" {
          tx.Where("{{.NameSnake}} >= ?", where.{{.Name}}Egt)
        }
         if len(where.{{.Name}}NotIn) > 0 {
          tx.Where("{{.NameSnake}} not in ?", where.{{.Name}}NotIn)
        }
        if where.{{.Name}}Sort != "" {
            if where.{{.Name}}Sort == "asc" {
                tx.Order("{{.NameSnake}} asc")
            } else {
                tx.Order("{{.NameSnake}} desc")
            }
        }
   {{- end}}
   {{- if eq .IsSoftDelete true}}
        if where.Unscoped_Del {
            tx.Unscoped()
        }
   {{- end}}
   {{- end}}
    where.SelectParams.SetTxGorm(tx)
    return tx
}

//返回数据总数
func ({{.Alias}}Repo *{{.Model}}RepoGorm) GetTotalCount(where repoInterface.{{.Model}}SearchWhere) int64 {
	tx := {{.Alias}}Repo.GetSearchWhereTx(where, nil)
	var r int64
	tx.Count(&r)
	return r
}

func ({{.Alias}}Repo *{{.Model}}RepoGorm) GetList(where repoInterface.{{.Model}}SearchWhere) []model.{{.Model}} {
	{{.Alias}} := make([]model.{{.Model}}, 0, where.SelectParams.RetSize)
	tx := {{.Alias}}Repo.GetSearchWhereTx(where, nil)
	tx.Find(&{{.Alias}})
	return {{.Alias}}
}

func ({{.Alias}}Repo *{{.Model}}RepoGorm) GetBy{{.Pk.Name}}({{.Pk.Name}} {{.Pk.Type}}, _select ...string) model.{{.Model}} {
	{{.Alias}} := model.{{.Model}}{}
	tx := {{.Alias}}Repo.Orm.Where("{{.Pk.NameSnake}}", {{.Pk.Name}})
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&{{.Alias}})
	return {{.Alias}}
}

{{- range $key, $item := .UniqueField}}
func ({{$.Alias}}Repo *{{$.Model}}RepoGorm) GetBy{{$key}}({{- range $item}}{{.NameFirstLower}} {{.Type}}, {{- end}} _select ...string) model.{{$.Model}} {
    {{$.Alias}} := model.{{$.Model}}{}
	tx := {{$.Alias}}Repo.Orm.
	{{- range $index, $val := $item}}
	{{- if eq $index (sub (len $item) 1)}}
	Where("{{$val.NameSnake}}", {{$val.NameFirstLower}})
	{{- else}}
	Where("{{$val.NameSnake}}", {{$val.NameFirstLower}}).
	{{- end}}
	{{- end}}
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&{{$.Alias}})
	return {{$.Alias}}
}

func ({{$.Alias}}Repo *{{$.Model}}RepoGorm) GetBy{{$key}}Lock({{- range $item}}{{.NameFirstLower}} {{.Type}}, {{- end}} _select ...string) (model.{{$.Model}}, repoComm.ReleaseLock) {
    if !orm.IsBeginTransaction({{$.Alias}}Repo.Orm) {
		panic("{{$.Alias}}Repo.GetBy{{$key}}Lock is must beginTransaction")
	}
    {{$.Alias}} := model.{{$.Model}}{}
	tx := {{$.Alias}}Repo.Orm.
	{{- range $index, $val := $item}}
	{{- if eq $index (sub (len $item) 1)}}
	Where("{{$val.NameSnake}}", {{$val.NameFirstLower}})
	{{- else}}
	Where("{{$val.NameSnake}}", {{$val.NameFirstLower}}).
	{{- end}}
	{{- end}}
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	orm.LockForUpdate(tx).Find(&{{$.Alias}})
	//这里返回一个空的释放锁方法 因为gorm在事务提交或回滚后会自动释放
    return {{$.Alias}}, func() {}
}


func ({{$.Alias}}Repo *{{$.Model}}RepoGorm) DeleteBy{{$key}}({{- range $item}}{{.NameFirstLower}} {{.Type}}, {{- end}}) (rowsAffected int64, err error) {
	tx := {{$.Alias}}Repo.Orm.
	{{- range $index, $val := $item}}
	{{- if eq $index (sub (len $item) 1)}}
	Where("{{$val.NameSnake}}", {{$val.NameFirstLower}})
	{{- else}}
	Where("{{$val.NameSnake}}", {{$val.NameFirstLower}}).
	{{- end}}
	{{- end}}
	r := tx.Delete(&model.{{$.Model}}{})
    return r.RowsAffected, r.Error
}
{{- end}}

{{- range $key, $item := .IndexField}}
func ({{$.Alias}}Repo *{{$.Model}}RepoGorm) GetBy{{$key}}({{- range $item}}{{.NameFirstLower}} {{.Type}}, {{- end}} _select ...string) []model.{{$.Model}} {
    {{$.Alias}} := make([]model.{{$.Model}}, 0)
	tx := {{$.Alias}}Repo.Orm.
	{{- range $index, $val := $item}}
	{{- if eq $index (sub (len $item) 1)}}
	Where("{{$val.NameSnake}}", {{$val.NameFirstLower}})
	{{- else}}
	Where("{{$val.NameSnake}}", {{$val.NameFirstLower}}).
	{{- end}}
	{{- end}}
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&{{$.Alias}})
	return {{$.Alias}}
}

func ({{$.Alias}}Repo *{{$.Model}}RepoGorm) DeleteBy{{$key}}({{- range $item}}{{.NameFirstLower}} {{.Type}}, {{- end}}) (rowsAffected int64, err error) {
	tx := {{$.Alias}}Repo.Orm.
	{{- range $index, $val := $item}}
	{{- if eq $index (sub (len $item) 1)}}
	Where("{{$val.NameSnake}}", {{$val.NameFirstLower}})
	{{- else}}
	Where("{{$val.NameSnake}}", {{$val.NameFirstLower}}).
	{{- end}}
	{{- end}}
	r := tx.Delete(&model.{{$.Model}}{})
    return r.RowsAffected, r.Error
}
{{- end}}

func ({{.Alias}}Repo *{{.Model}}RepoGorm) GetBy{{.Pk.Name}}Lock({{.Pk.Name}} {{.Pk.Type}}, _select ...string) (model.{{.Model}}, repoComm.ReleaseLock) {
	if !orm.IsBeginTransaction({{.Alias}}Repo.Orm) {
		panic("{{.Alias}}Repo.GetBy{{.Pk.Name}}Lock is must beginTransaction")
	}
	{{.Alias}} := model.{{.Model}}{}
	tx := orm.LockForUpdate({{.Alias}}Repo.Orm.Where("{{.Pk.NameSnake}}", {{.Pk.Name}}))
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.Find(&{{.Alias}})

	//这里返回一个空的释放锁方法 因为gorm在事务提交或回滚后会自动释放
	return {{.Alias}}, func() {}
}

func ({{.Alias}}Repo *{{.Model}}RepoGorm) ScanByWhere(where repoInterface.{{.Model}}SearchWhere, dest any) error {
	return {{.Alias}}Repo.GetSearchWhereTx(where, nil).Find(dest).Error
}

func ({{.Alias}}Repo *{{.Model}}RepoGorm) ScanByOrWhere(dest any, where ...repoInterface.{{.Model}}SearchWhere) error {
	tx := {{.Alias}}Repo.Orm.Model(model.{{.Model}}{})
	for _, v := range where {
		tx.Or({{.Alias}}Repo.GetSearchWhereTx(v, nil))
		v.SelectParams.SetTxGorm(tx)
	}
	return tx.Find(dest).Error
}
