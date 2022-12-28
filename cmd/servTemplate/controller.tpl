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

}

func ({{.Alias}}Ctrl {{.Model}}Controller) GetList(ctx iris.Context) appWeb.ResponseFormat {
    {{.Alias}}Serv := services.New{{.Model}}Service()
    {{.Alias}}, pager := {{.Alias}}Serv.ListPage(ctx)
    return appWeb.NewPagerResponse({{.Alias}}Serv.ShowMapList({{.Alias}}), pager)
}

func ({{.Alias}}Ctrl {{.Model}}Controller) GetItem(ctx iris.Context) appWeb.ResponseFormat {
	{{.Alias}}Serv := services.New{{.Model}}Service()
	{{.Alias}} := {{.Alias}}Serv.GetItem(ctx)
	return appWeb.NewSuccessResponse("", {{.Alias}}.ShowMap())
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
