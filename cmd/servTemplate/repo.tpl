package repositories

import (
	"gorm.io/gorm"
	"iris-init/model"
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
func ({{.Alias}}Repo {{.Model}}RepoGorm) GetByWhere(where repoInterface.{{.Model}}SearchWhere) model.{{.Model}} {
	{{.Alias}} := model.{{.Model}}{}
	_ = {{.Alias}}Repo.GetSearchWhereTx(where, nil).First(&{{.Alias}})
	return {{.Alias}}
}

//该方法需要自己去完善 GetSearchWhereTx方法内部
func ({{.Alias}}Repo {{.Model}}RepoGorm) GetIDByWhere(where repoInterface.{{.Model}}SearchWhere) []uint64 {
	var id []uint64
	tx := {{.Alias}}Repo.GetSearchWhereTx(where, nil)
	tx.Select("id").Model(model.{{.Model}}{}).Scan(&id)
	return id
}

func ({{.Alias}}Repo {{.Model}}RepoGorm) Save({{.Alias}} *model.{{.Model}}, _select ...string) error {
	return repoComm.SaveModel({{.Alias}}Repo.Orm, {{.Alias}}, _select...)
}

func ({{.Alias}}Repo {{.Model}}RepoGorm) SaveOmit({{.Alias}} *model.{{.Model}}, _omit ...string) error {
	return repoComm.SaveModelOmit({{.Alias}}Repo.Orm, {{.Alias}}, _omit...)
}

//这里因为gorm的缘故 传入的{{.Alias}}主键必须不为空
func ({{.Alias}}Repo {{.Model}}RepoGorm) Delete({{.Alias}} model.{{.Model}}) (rowsAffected int64, err error) {
    tx := {{.Alias}}Repo.Orm.Delete({{.Alias}})
	return tx.RowsAffected, tx.Error
}

//为了避免更换源之后的一些麻烦 该方法不建议在仓库结构{{.Model}}RepoGorm以外使用
func ({{.Alias}}Repo {{.Model}}RepoGorm) deleteByWhere(query string, args ...interface{}) (rowsAffected int64, err error) {
	tx := {{.Alias}}Repo.Orm.Where(query, args...).Delete(model.{{.Model}}{})
	return tx.RowsAffected, tx.Error
}

func ({{.Alias}}Repo {{.Model}}RepoGorm) DeleteByID(id ...uint64) (rowsAffected int64, err error) {
	if len(id) == 1 {
		return {{.Alias}}Repo.deleteByWhere("id", id[0])
	}
	return {{.Alias}}Repo.deleteByWhere("id in ?", id)
}


func ({{.Alias}}Repo {{.Model}}RepoGorm) GetSearchWhereTx(where repoInterface.{{.Model}}SearchWhere, tx0 *gorm.DB) *gorm.DB {
	var tx *gorm.DB
	if tx0 == nil {
		tx = {{.Alias}}Repo.Orm.Model(model.{{.Model}}{})
	} else {
		tx = tx0.Model(model.{{.Model}}{})
	}
   {{- range .ModelField}}
   //需要额外调整
   	if false {
        tx.Where("{{.Name}}", where.{{.Name}})
   	}
   {{- end}}
	return tx
}

//返回数据总数
func ({{.Alias}}Repo {{.Model}}RepoGorm) GetTotalCount(where repoInterface.{{.Model}}SearchWhere) int64 {
	tx := {{.Alias}}Repo.GetSearchWhereTx(where, nil)
	var r int64
	tx.Count(&r)
	return r
}

func ({{.Alias}}Repo {{.Model}}RepoGorm) GetList(where repoInterface.{{.Model}}SearchWhere) []model.{{.Model}} {
	{{.Alias}} := make([]model.{{.Model}}, 0, where.SelectParams.RetSize)
	tx := {{.Alias}}Repo.GetSearchWhereTx(where, nil)
	where.SelectParams.SetTxGorm(tx).Find(&{{.Alias}})
	return {{.Alias}}
}

func ({{.Alias}}Repo {{.Model}}RepoGorm) GetByID(id uint64, _select ...string) model.{{.Model}} {
	if id == 0 {
		return model.{{.Model}}{}
	}
	{{.Alias}} := model.{{.Model}}{}
	tx := {{.Alias}}Repo.Orm.Where("id", id)
	if len(_select) > 0 {
		tx = tx.Select(_select)
	}
	tx.First(&{{.Alias}})
	return {{.Alias}}
}