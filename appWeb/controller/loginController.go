package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"jd-fxl/appWeb"
	"jd-fxl/global"
	"jd-fxl/model"
	"jd-fxl/services"
)

type LoginController struct {
	Adm model.Admin
}

func (login LoginController) GetLogin(ctx iris.Context) mvc.Result {
	if login.Adm.ID > 0 {
		ctx.Redirect("/")
		return mvc.Response{}
	}
	return appWeb.ResponseDataViewForm("site/login.html", appWeb.DataView{
		PageCss: []string{"/static/style/login.css"},
	}, ctx)
}

func (login LoginController) PostLogin(ctx iris.Context) appWeb.ResponseFormat {
	admServ := services.NewAdminService()
	adm, err := admServ.LoginByPwd(ctx.PostValueTrim("Username"), ctx.PostValueTrim("Password"))
	if err != nil {
		return appWeb.NewFailErrResponse(err, nil)
	}
	//设置session
	sessions.Get(ctx).Set(global.ReqTokenName, adm.Token.String)
	return appWeb.NewSuccessResponse("", map[string]interface{}{
		appWeb.AjaxLocationKey: "/",
	})
}
