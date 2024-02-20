package services

import (
    "github.com/kataras/iris/v12"
    "iris-init/global"
    "iris-init/model"
    "iris-init/repositories"
    "iris-init/repositories/repoComm"
    "iris-init/repositories/repoInterface"
    "iris-init/sErr"
    {{- $stop := false }}
    {{- range .UniqueField}}
     {{- if not $stop}}
     "reflect"
     {{- $stop := true }}
     {{- end}}
    {{- end}}
)

func New{{.Model}}Service() {{.Model}}Service {
	return {{.Model}}Service{repo: repositories.New{{.Model}}Repo()}
}

func New{{.Model}}ServiceByOrm(orm any) {{.Model}}Service {
	r := {{.Model}}Service{repo: repositories.New{{.Model}}Repo()}
	r.repo.SetOrm(orm)
	return r
}

func New{{.Model}}ServiceByRepo(repo repoInterface.{{.Model}}Repo) {{.Model}}Service {
	return {{.Model}}Service{repo: repo}
}

type {{.Model}}Service struct {
	repo repoInterface.{{.Model}}Repo
}

func ({{.Alias}}Serv {{.Model}}Service) ListPage(ctx iris.Context) ([]model.{{.Model}}, *global.Pager) {
	where := repoInterface.{{.Model}}SearchWhere{}
	_ = ctx.ReadQuery(&where)
	where.SelectParams = repoComm.SelectFrom{
        OrderBy: []repoComm.OrderByParams{
            {
                Column: "ID",
                Desc:   true,
            },
        },
    }
	pager := global.NewPager(ctx)
    if pager.Size < 0 {
        return {{.Alias}}Serv.repo.GetList(where), nil
    }
	pager.SetTotal({{.Alias}}Serv.repo.GetTotalCount(where))
	if pager.Total == 0 {
		return []model.{{.Model}}{}, pager
	}
	where.SelectParams.Offset = pager.Offset
    where.SelectParams.Limit = pager.Size
    where.SelectParams.RetSize = pager.Size
	return {{.Alias}}Serv.repo.GetList(where), pager
}

func ({{.Alias}}Serv {{.Model}}Service) ListAvailable(_select ...string) []model.{{.Model}} {
	if len(_select) == 0 {
		_select = nil
	}
	return {{.Alias}}Serv.repo.GetList(repoInterface.{{.Model}}SearchWhere{
		SelectParams: repoComm.SelectFrom{
			Select: _select,
		},
	})
}


func ({{.Alias}}Serv {{.Model}}Service) ListByWhere(where repoInterface.{{.Model}}SearchWhere) []model.{{.Model}} {
	return {{.Alias}}Serv.repo.GetList(where)
}

func ({{.Alias}}Serv {{.Model}}Service) TotalCount(where repoInterface.{{.Model}}SearchWhere) int64 {
	return {{.Alias}}Serv.repo.GetTotalCount(where)
}

// 获取一条数据根据ctx
// 这里条件为ID 传入ctx是方便后续修改参数条件
func ({{.Alias}}Serv {{.Model}}Service) GetItem(ctx iris.Context, _select ...string) model.{{.Model}} {
	return {{.Alias}}Serv.repo.GetBy{{.Pk.Name}}(ctx.URLParamUint64("{{.Pk.Name}}"), _select...)
}

func ({{.Alias}}Serv {{.Model}}Service) GetBy{{.Pk.Name}}({{.Pk.Name}} {{.Pk.Type}}, _select ...string) model.{{.Model}} {
	return {{.Alias}}Serv.repo.GetBy{{.Pk.Name}}({{.Pk.Name}}, _select...)
}

func ({{.Alias}}Serv {{.Model}}Service) GetByWhere(where repoInterface.{{.Model}}SearchWhere) model.{{.Model}} {
	return {{.Alias}}Serv.repo.GetByWhere(where)
}

func ({{.Alias}}Serv {{.Model}}Service) ScanByWhere(where repoInterface.{{.Model}}SearchWhere, dest any) error {
    return {{.Alias}}Serv.repo.ScanByWhere(where, dest)
}

func ({{.Alias}}Serv {{.Model}}Service) ScanByOrWhere(dest any, where ...repoInterface.{{.Model}}SearchWhere) error {
    return {{.Alias}}Serv.repo.ScanByOrWhere(dest, where...)
}

func ({{.Alias}}Serv {{.Model}}Service) UpdateByWhere(where repoInterface.{{.Model}}SearchWhere, data interface{}) (rowsAffected int64, err error) {
	return {{.Alias}}Serv.repo.UpdateByWhere(where, data)
}

{{- range  $key, $item := .UniqueField}}
func ({{$.Alias}}Serv {{$.Model}}Service) GetBy{{$key}}({{- range $item}}{{.NameFirstLower}} {{.Type}}, {{- end}} _select ...string) model.{{$.Model}} {
    var v reflect.Value
    {{- range $item}}
    v = reflect.ValueOf({{.NameFirstLower}})
    if !v.IsValid() { // 值不存在
         return model.{{$.Model}}{};
    }
    {{- end}}
    return {{$.Alias}}Serv.repo.GetBy{{$key}}({{- range $item}}{{.NameFirstLower}}, {{- end}} _select...)
}

func ({{$.Alias}}Serv {{$.Model}}Service) Check{{$key}}Valid({{$.Alias}} model.{{$.Model}}) error {
    var v reflect.Value
    {{- range $item}}
    v = reflect.ValueOf({{$.Alias}}.{{.Name}})
    if !v.IsValid() { // 值不存在
         return sErr.New("无效的{{.Label}}")
    }
    {{- end}}
    _{{$.Alias}} := {{$.Alias}}Serv.GetBy{{$key}}({{- range $item}}{{- if eq .TypeOrigin "sql.NullString" }}{{$.Alias}}.{{.Name}}.String{{- else}}{{$.Alias}}.{{.Name}}{{- end}}, {{- end}} "id")
    if _{{$.Alias}}.ID > 0 && {{$.Alias}}.ID != _{{$.Alias}}.ID {
        return sErr.NewFmt("{{- range $item}}{{.Label}}.{{- end}}已存在: {{- range $item}}%s.{{- end}}", {{- range $item}}{{$.Alias}}.{{.Name}},{{- end}})
    }
    return nil
}

func ({{$.Alias}}Serv {{$.Model}}Service) DeleteBy{{$key}}({{- range $item}}{{.NameFirstLower}} {{.Type}}, {{- end}}) error {
    _, err := {{$.Alias}}Serv.repo.DeleteBy{{$key}}({{- range $item}}{{.NameFirstLower}}, {{- end}})
    return err
}
{{- end}}

{{- range  $key, $item := .IndexField}}
func ({{$.Alias}}Serv {{$.Model}}Service) GetBy{{$key}}({{- range $item}}{{.NameFirstLower}} {{.Type}}, {{- end}} _select ...string) []model.{{$.Model}} {
    return {{$.Alias}}Serv.repo.GetBy{{$key}}({{- range $item}}{{.NameFirstLower}}, {{- end}} _select...)
}

func ({{$.Alias}}Serv {{$.Model}}Service) DeleteBy{{$key}}({{- range $item}}{{.NameFirstLower}} {{.Type}}, {{- end}}) error {
    _, err := {{$.Alias}}Serv.repo.DeleteBy{{$key}}({{- range $item}}{{.NameFirstLower}}, {{- end}})
    return err
}
{{- end}}

func ({{.Alias}}Serv {{.Model}}Service) GetBy{{.Pk.Name}}Lock({{.Pk.Name}} {{.Pk.Type}}, _select ...string) (model.{{.Model}}, repoComm.ReleaseLock) {
	return {{.Alias}}Serv.repo.GetBy{{.Pk.Name}}Lock({{.Pk.Name}}, _select...)
}

{{range  $key, $item := .ReferencesField}}
func ({{$.Alias}}Serv {{$.Model}}Service) Reload{{$item.Name}}({{$.Alias}} *model.{{$.Model}}) {
    {{$.Alias}}.{{$item.Name}} = New{{$item.ReferencesModel}}ServiceByOrm({{$.Alias}}Serv.repo.GetOrm()).
                    GetBy{{$item.ForeignKey}}({{$.Alias}}.{{$item.References}})
}
{{- end}}

// 通过请求ctx编辑/新增一条数据
func ({{.Alias}}Serv {{.Model}}Service) SaveByCtx(ctx iris.Context) (model.{{.Model}}, error) {
	{{.Alias}}Validator := {{.Model}}Validator{}
	err := global.ScanValidatorByRequestPost(ctx, &{{.Alias}}Validator)
	if err != nil {
		return model.{{.Model}}{}, err
	}
	return {{.Alias}}Serv.SaveByValidator({{.Alias}}Validator)
}

func ({{.Alias}}Serv {{.Model}}Service) SaveByValidator({{.Alias}}Validator {{.Model}}Validator) (model.{{.Model}}, error) {
	{{.Alias}}, err := {{.Alias}}Serv.Get{{.Model}}ByValidate({{.Alias}}Validator)
	if err != nil {
		return {{.Alias}}, err
	}
	err = {{.Alias}}Serv.Save(&{{.Alias}})
	return {{.Alias}}, err
}

func ({{.Alias}}Serv {{.Model}}Service) Save({{.Alias}} *model.{{.Model}}) error {
	return {{.Alias}}Serv.repo.Save({{.Alias}})
}

func ({{.Alias}}Serv {{.Model}}Service) Create({{.Alias}} *[]model.{{.Model}}) error {
	return {{.Alias}}Serv.repo.Create({{.Alias}})
}

func ({{.Alias}}Serv {{.Model}}Service) DeleteByCtx(ctx iris.Context) error {
	return {{.Alias}}Serv.DeleteByID(uint64(ctx.PostValueInt64Default("{{.Pk.Name}}", 0)))
}

//这个方法目前是与DeleteByID功能一致 主要是用来扩展的 根据model的多条件作删除 需要开发者自己完成业务逻辑
func ({{.Alias}}Serv {{.Model}}Service) Delete({{.Alias}} model.{{.Model}}) error {
	_, err := {{.Alias}}Serv.repo.DeleteByID({{.Alias}}.{{.Pk.Name}})
    return err
}

func ({{.Alias}}Serv {{.Model}}Service) DeleteByID({{.Pk.Name}} ...{{.Pk.Type}}) error {
	_, err := {{.Alias}}Serv.repo.DeleteByID({{.Pk.Name}}...)
	return err
}

func ({{.Alias}}Serv {{.Model}}Service) ShowMapList({{.Alias}} []model.{{.Model}}) []map[string]interface{} {
	_{{.Alias}} := []map[string]interface{}{}
	for _, v := range {{.Alias}} {
		_{{.Alias}} = append(_{{.Alias}}, v.ShowMap())
	}
	return _{{.Alias}}
}

// 验证参数 并返回到一个新的{{.Model}} model
func ({{.Alias}}Serv {{.Model}}Service) Get{{.Model}}ByValidate({{.Alias}}Validator {{.Model}}Validator) (model.{{.Model}}, error) {
	err := {{.Alias}}Validator.Validate()
	if err != nil {
		return model.{{.Model}}{}, err
	}
	var {{.Alias}} model.{{.Model}}
	if {{.Alias}}Validator.{{.Pk.Name}} > 0 {
		{{.Alias}} = {{.Alias}}Serv.repo.GetBy{{.Pk.Name}}({{.Alias}}Validator.{{.Pk.Name}})
		if {{.Alias}}.{{.Pk.Name}} == 0 {
			return {{.Alias}}, sErr.New("无效的{{.Pk.Name}}")
		}
	} else {
		{{.Alias}} = model.{{.Model}}{}
	}
	//完成其他的赋值逻辑处理...
    {{- range .ModelField}}
        {{- if not .OnlyRead}}
        {{$.Alias}}.{{.Name}} = {{$.Alias}}Validator.{{.Name}}
        {{- end}}
    {{- end}}

    {{- range $key, $item := .UniqueField}}
    if err = {{$.Alias}}Serv.Check{{$key}}Valid({{$.Alias}}); err != nil {
        return {{$.Alias}}, err
    }
    {{- end}}
	return {{.Alias}}, nil
}

type {{.Model}}Validator struct {
	{{- range .ModelField}}
        {{- if not .OnlyRead}}
            {{- if ne .ValidateLabel ""}}
            {{.Name}}   {{.Type}} `{{.ValidateLabel}} label:"{{.Label}}"`
            {{- else}}
            {{.Name}}   {{.Type}} `label:"{{.Label}}"`
            {{- end}}
        {{- end}}
	{{- end}}
}

func ({{.Alias}}Validator *{{.Model}}Validator) Validate() error {
	err := global.ValidateV9Struct({{.Alias}}Validator)
	if err != nil {
		return err
	}
	return nil
}
