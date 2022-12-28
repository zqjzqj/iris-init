package {{.Package}}

import (
	"github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/mvc"
    "iris-init/appWeb"
    "iris-init/services"
    "net/http"
)

type {{.Model}}Controller struct {}

func ({{.Alias}}Ctr {{.Model}}Controller) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(http.MethodGet, "list", "GetList")
	b.Handle(http.MethodGet, "item", "GetItem")
	b.Handle(http.MethodPost, "edit", "PostEdit")
	b.Handle(http.MethodPost, "delete", "PostDelete")

}

func ({{.Alias}}Ctr {{.Model}}Controller) GetList(ctx iris.Context) appWeb.ResponseFormat {
    return appWeb.NewPagerResponse(services.New{{.Model}}Service().ListPage(ctx))
}

func ({{.Alias}}Ctr {{.Model}}Controller) GetItem(ctx iris.Context) appWeb.ResponseFormat {
	{{.Alias}}Serv := services.New{{.Model}}Service()
	{{.Alias}} := {{.Alias}}Serv.GetItem(ctx)
	return appWeb.NewSuccessResponse("", {{.Alias}}.ShowMap())
}

func ({{.Alias}}Ctr {{.Model}}Controller) PostEdit(ctx iris.Context) appWeb.ResponseFormat {
	{{.Alias}}, err := services.New{{.Model}}Service().EditByCtx(ctx)
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	return appWeb.NewSuccessResponse("", {{.Alias}}.ShowMap())
}

func ({{.Alias}}Ctr {{.Model}}Controller) PostDelete(ctx iris.Context) appWeb.ResponseFormat {
	err := services.New{{.Model}}Service().DeleteByCtx(ctx)
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	return appWeb.NewSuccessResponse("", nil)
}
