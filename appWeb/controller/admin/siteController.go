package admin

import (
	"9xbet_risk/appWeb"
	"9xbet_risk/model"
	"9xbet_risk/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type SiteController struct {
	Adm model.Admin
}

func (siteCtrl SiteController) Get(ctx iris.Context) mvc.Result {
	services.NewAdminService().RefreshPermissions(&siteCtrl.Adm, false, false)
	menus := services.NewPermissionsService().GetPremAsMenu(siteCtrl.Adm.Permissions)
	return appWeb.ResponseDataViewForm("site/index.html", appWeb.DataView{
		Data: map[string]interface{}{
			"Menus": menus,
		},
	}, ctx)
}

func (siteCtrl SiteController) GetErr(ctx iris.Context) mvc.Result {
	return appWeb.ResponseErrView(ctx.URLParamTrim("Msg"), ctx)
}
