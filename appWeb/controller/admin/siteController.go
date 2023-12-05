package admin

import (
	"big_data_new/appWeb"
	"big_data_new/model"
	"big_data_new/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"net/http"
)

type SiteController struct {
	Adm model.Admin
}

func (siteCtrl SiteController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(http.MethodGet, "perms-menus", "GetPermsMenus")
}

func (siteCtrl SiteController) Get(ctx iris.Context) appWeb.ResponseFormat {
	return appWeb.NewSuccessResponse("OK", nil)
}

// 获取当前用户的权限菜单
func (siteCtrl SiteController) GetPermsMenus(ctx iris.Context) appWeb.ResponseFormat {
	services.NewAdminService().RefreshPermissions(&siteCtrl.Adm, false, false)
	menus := services.NewPermissionsService().GetPremAsMenu(siteCtrl.Adm.Permissions)
	return appWeb.NewSuccessResponse("OK", menus)
}
