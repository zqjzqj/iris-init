package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"jd-fxl/appWeb"
	"jd-fxl/model"
	"jd-fxl/services"
)

type SiteController struct {
	Adm model.Admin
}

func (site SiteController) Get(ctx iris.Context) mvc.Result {
	services.NewAdminService().RefreshPermissions(&site.Adm, false)
	menus := services.NewPermissionService().GetPremAsMenu(site.Adm.Permissions)
	return appWeb.ResponseDataViewForm("site/index.html", appWeb.DataView{
		Data: map[string]interface{}{
			"Menus": menus,
		},
	}, ctx)
}

func (site SiteController) GetErr(ctx iris.Context) mvc.Result {
	return appWeb.ResponseErrView(ctx.URLParamTrim("Msg"), ctx)
}
