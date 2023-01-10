package {{.Package}}

import (
	"github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/mvc"
    "iris-init/appWeb"
    "iris-init/services"
    "net/http"
)

type {{.Model}}Controller struct {}

func ({{.Alias}}Ctrl {{.Model}}Controller) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(http.MethodGet, "list", "GetList")
	b.Handle(http.MethodGet, "item", "GetItem")
	b.Handle(http.MethodPost, "edit", "PostEdit")
	b.Handle(http.MethodPost, "delete", "PostDelete")

	{{- if .View}}
	b.Handle(http.MethodGet, "list-view.html", "GetListView")
	b.Handle(http.MethodGet, "item-view.html", "GetItemView")
	{{- end }}
}

func ({{.Alias}}Ctrl {{.Model}}Controller) GetList(ctx iris.Context) appWeb.ResponseFormat {
    {{.Alias}}Serv := services.New{{.Model}}Service()
    {{.Alias}}, pager := {{.Alias}}Serv.ListPage(ctx)
    return appWeb.NewPagerResponse({{.Alias}}Serv.ShowMapList({{.Alias}}), pager)
}
{{- if .View}}

func ({{.Alias}}Ctrl {{.Model}}Controller) GetListView(ctx iris.Context) mvc.Result {
	{{.Alias}}Serv := services.New{{.Model}}Service()
	{{.Alias}}, pager := {{.Alias}}Serv.ListPage(ctx)
	return appWeb.ResponseDataViewForm("{{.Alias}}/list.html", appWeb.DataView{
		Pager: pager,
		Data: map[string]interface{}{
			"List": {{.Alias}}Serv.ShowMapList({{.Alias}}),
		},
	}, ctx)
}

{{- end }}
func ({{.Alias}}Ctrl {{.Model}}Controller) GetItem(ctx iris.Context) appWeb.ResponseFormat {
	{{.Alias}}Serv := services.New{{.Model}}Service()
	{{.Alias}} := {{.Alias}}Serv.GetItem(ctx)
	return appWeb.NewSuccessResponse("", {{.Alias}}.ShowMap())
}
{{- if .View}}

func ({{.Alias}}Ctrl {{.Model}}Controller) GetItemView(ctx iris.Context) mvc.Result {
	return appWeb.ResponseDataViewForm("{{.Alias}}/item.html", appWeb.DataView{
		Data: map[string]interface{}{
			"item": services.New{{.Model}}Service().GetItem(ctx).ShowMap(),
		},
	}, ctx)
}

{{- end }}

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
