package {{.Package}}

import (
	"github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/mvc"
    "iris-init/appWeb"
    "iris-init/services"
    {{- if .View}}
    "iris-init/global"
    {{- end}}
    "net/http"
)

type {{.Model}}Controller struct {}

func ({{.Alias}}Ctrl {{.Model}}Controller) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(http.MethodGet, "list", "GetList").SetName("{{.Model}}@{{.Model}}List")
    b.Handle(http.MethodGet, "item", "GetItem").SetName("{{.Model}}@{{.Model}}:详情")
    b.Handle(http.MethodPost, "edit", "PostEdit").SetName("{{.Model}}@{{.Model}}:编辑")
    b.Handle(http.MethodPost, "delete", "PostDelete").SetName("{{.Model}}@{{.Model}}:删除")
}

func ({{.Alias}}Ctrl {{.Model}}Controller) GetList(ctx iris.Context) any {
    {{.Alias}}Serv := services.New{{.Model}}Service()
    {{.Alias}}, pager := {{.Alias}}Serv.ListPage(ctx)
    {{- if .View}}
    if global.IsApiReq(ctx) {
        return appWeb.NewPagerResponse(map[string]interface{}{
            "List": {{.Alias}}Serv.ShowMapList({{.Alias}}),
        }, pager)
    }
    return appWeb.ResponseDataViewForm("{{.Alias}}/list.html", appWeb.DataView{
        Pager: pager,
        Data: map[string]interface{}{
            "List": {{.Alias}}Serv.ShowMapList({{.Alias}}),
        },
    }, ctx)
    {{- else}}
    return appWeb.NewPagerResponse({{.Alias}}Serv.ShowMapList({{.Alias}}), pager)
    {{- end}}
}

func ({{.Alias}}Ctrl {{.Model}}Controller) GetItem(ctx iris.Context) any {
	{{.Alias}}Serv := services.New{{.Model}}Service()
	{{.Alias}} := {{.Alias}}Serv.GetItem(ctx)
	{{- if .View}}
    if global.IsApiReq(ctx) {
        return appWeb.NewSuccessResponse("", map[string]interface{}{
            "Item": {{.Alias}}.ShowMap(),
        })
    }
    return appWeb.ResponseDataViewForm("{{.Alias}}/item.html", appWeb.DataView{
        Data: map[string]interface{}{
            "Item": {{.Alias}}.ShowMap(),
        },
    }, ctx)
    {{- else}}
    return appWeb.NewSuccessResponse("", map[string]interface{}{
        "Item": {{.Alias}}.ShowMap(),
    })
    {{- end}}
}

func ({{.Alias}}Ctrl {{.Model}}Controller) PostEdit(ctx iris.Context) appWeb.ResponseFormat {
	{{.Alias}}, err := services.New{{.Model}}Service().EditByCtx(ctx)
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	return appWeb.NewSuccessResponse("", {{.Alias}}.ShowMap())
}

func ({{.Alias}}Ctrl {{.Model}}Controller) PostDelete(ctx iris.Context) appWeb.ResponseFormat {
	err := services.New{{.Model}}Service().DeleteByCtx(ctx)
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	return appWeb.NewSuccessResponse("", nil)
}