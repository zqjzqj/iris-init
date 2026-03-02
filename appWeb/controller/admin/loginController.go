package admin

import (
	"9xbet_risk/appWeb"
	"9xbet_risk/global"
	"9xbet_risk/model"
	"9xbet_risk/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type LoginController struct {
	Adm model.Admin
}

func (loginCtrl LoginController) GetLogin(ctx iris.Context) mvc.Result {
	if loginCtrl.Adm.ID > 0 {
		ctx.Redirect("/")
		return mvc.Response{}
	}
	return appWeb.ResponseDataViewForm("site/login.html", appWeb.DataView{
		PageCss: []string{"/static/style/login.css"},
	}, ctx)
}

func (loginCtrl LoginController) PostLogin(ctx iris.Context) appWeb.ResponseFormat {
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
