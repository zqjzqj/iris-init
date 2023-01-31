package services

import (
    "github.com/kataras/iris/v12"
    "iris-init/global"
    "iris-init/model"
    "iris-init/repositories"
    "iris-init/repositories/repoComm"
    "iris-init/repositories/repoInterface"
    "iris-init/sErr"
)

func New{{.Model}}Service() {{.Model}}Service {
	return {{.Model}}Service{repo: repositories.New{{.Model}}Repo()}
}

type {{.Model}}Service struct {
	repo repoInterface.{{.Model}}Repo
}

func ({{.Alias}}Serv {{.Model}}Service) ListPage(ctx iris.Context) ([]model.{{.Model}}, *global.Pager) {
	where := repoInterface.{{.Model}}SearchWhere{}
	_ = ctx.ReadQuery(&where)
	pager := global.NewPager(ctx)
	pager.SetTotal({{.Alias}}Serv.repo.GetTotalCount(where))
	if pager.Total == 0 {
		return []model.{{.Model}}{}, pager
	}
	where.SelectParams = repoComm.SelectFrom{
		Offset:  pager.Offset,
		Limit:   pager.Size,
		RetSize: pager.Size,
		OrderBy: []repoComm.OrderByParams{
            {
                Column: "ID",
                Desc:   true,
            },
		},
	}
	return {{.Alias}}Serv.repo.GetList(where), pager
}

// 获取一条数据根据ctx
// 这里条件为ID 传入ctx是方便后续修改参数条件
func ({{.Alias}}Serv {{.Model}}Service) GetItem(ctx iris.Context, _select ...string) model.{{.Model}} {
	return {{.Alias}}Serv.repo.GetByID(ctx.URLParamUint64("ID"), _select...)
}

func ({{.Alias}}Serv {{.Model}}Service) GetByID(id uint64, _select ...string) model.{{.Model}} {
    if id == 0 {
		return model.{{.Model}}{}
	}
	return {{.Alias}}Serv.repo.GetByID(id, _select...)
}

// 通过请求ctx编辑/新增一条数据
func ({{.Alias}}Serv {{.Model}}Service) EditByCtx(ctx iris.Context) (model.{{.Model}}, error) {
	{{.Alias}}Validator := {{.Model}}Validator{}
	err := ctx.ReadBody(&{{.Alias}}Validator)
	if err != nil {
		return model.{{.Model}}{}, err
	}
	return {{.Alias}}Serv.EditByValidator({{.Alias}}Validator)
}

func ({{.Alias}}Serv {{.Model}}Service) EditByValidator({{.Alias}}Validator {{.Model}}Validator) (model.{{.Model}}, error) {
	{{.Alias}}, err := {{.Alias}}Serv.Get{{.Model}}ByValidate({{.Alias}}Validator)
	if err != nil {
		return {{.Alias}}, err
	}
	err = {{.Alias}}Serv.repo.Save(&{{.Alias}})
	return {{.Alias}}, err
}

func ({{.Alias}}Serv {{.Model}}Service) DeleteByCtx(ctx iris.Context) error {
	return {{.Alias}}Serv.Delete({{.Alias}}Serv.repo.GetByID(uint64(ctx.PostValueInt64Default("ID", 0))))
}

func ({{.Alias}}Serv {{.Model}}Service) Delete({{.Alias}} model.{{.Model}}) error {
	if {{.Alias}}.ID == 0 {
		return nil
	}
	_, err := {{.Alias}}Serv.repo.DeleteByID({{.Alias}}.ID)
    return err
}

func ({{.Alias}}Serv {{.Model}}Service) ShowMapList({{.Alias}} []model.{{.Model}}) []map[string]interface{} {
	_{{.Alias}} := []map[string]interface{}{}
	for _, v := range {{.Alias}} {
		_{{.Alias}} = append(_{{.Alias}}, v.ShowMap())
	}
	return _{{.Alias}}
}

// 验证参数 并返回到一个新的adm model
func ({{.Alias}}Serv {{.Model}}Service) Get{{.Model}}ByValidate({{.Alias}}Validator {{.Model}}Validator) (model.{{.Model}}, error) {
	err := {{.Alias}}Validator.Validate()
	if err != nil {
		return model.{{.Model}}{}, err
	}
	var {{.Alias}} model.{{.Model}}
	if {{.Alias}}Validator.ID > 0 {
		{{.Alias}} = {{.Alias}}Serv.repo.GetByID({{.Alias}}Validator.ID)
		if {{.Alias}}.ID == 0 {
			return {{.Alias}}, sErr.New("无效的ID")
		}
	} else {
		{{.Alias}} = model.{{.Model}}{}
	}
	//完成其他的赋值逻辑处理...
    {{- range .ModelField}}
    {{$.Alias}}.{{.Name}} = {{$.Alias}}Validator.{{.Name}}
    {{- end}}
	return {{.Alias}}, nil
}

type {{.Model}}Validator struct {
	{{- range .ModelField}}
	{{.Name}}   {{.Type}} `label:"{{.Label}}"`
	{{- end}}
}

func ({{.Alias}}Validator {{.Model}}Validator) Validate() error {
	err := global.ValidateV9Struct({{.Alias}}Validator)
	if err != nil {
		return err
	}
	return nil
}
