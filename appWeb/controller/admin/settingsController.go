package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris-init/appWeb"
	"iris-init/global"
	"iris-init/model"
	"iris-init/services"
	"net/http"
)

type SettingsController struct {
	Admin model.Admin
}

func (settingsCtrl SettingsController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(http.MethodGet, "list", "GetList").SetName("基础设置@系统设置")
	b.Handle(http.MethodPost, "edit", "PostEdit").SetName("基础设置@系统设置:编辑")
}

func (settingsCtrl SettingsController) GetList(ctx iris.Context) any {
	settingsServ := services.NewSettingsService()
	settings := settingsServ.ListAvailable()
	if global.IsApiReq(ctx) {
		return appWeb.NewSuccessResponse("", map[string]interface{}{
			"List": settingsServ.ShowMapList(settings),
		})
	}
	return appWeb.ResponseDataViewForm("settings/list.html", appWeb.DataView{
		Data: map[string]interface{}{
			"List": settingsServ.ShowMapList(settings),
		},
	}, ctx)
}

func (settingsCtrl SettingsController) PostEdit(ctx iris.Context) appWeb.ResponseFormat {
	settingsServ := services.NewSettingsService()
	err := settingsServ.SaveByCtx(ctx)
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	return appWeb.NewSuccessResponse("", nil)
}
