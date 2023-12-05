package admin

import (
	"big_data_new/appWeb"
	"big_data_new/model"
	"big_data_new/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"net/http"
)

type SettingsController struct {
	Admin model.Admin
}

func (settingsCtrl SettingsController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(http.MethodGet, "list", "GetList").SetName("基础设置@系统设置")
	b.Handle(http.MethodPost, "edit", "PostEdit").SetName("基础设置@系统设置:编辑")
}

func (settingsCtrl SettingsController) GetList(ctx iris.Context) appWeb.ResponseFormat {
	settingsServ := services.NewSettingsService()
	settings := settingsServ.ListAvailable()
	return appWeb.NewSuccessResponse("", map[string]interface{}{
		"List": settingsServ.ShowMapList(settings),
	})
}

func (settingsCtrl SettingsController) PostEdit(ctx iris.Context) appWeb.ResponseFormat {
	settingsServ := services.NewSettingsService()
	err := settingsServ.SaveByCtx(ctx)
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	return appWeb.NewSuccessResponse("", nil)
}
